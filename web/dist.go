// Package web
// @Time  : 2022/2/13 下午12:31
// @Author: Jtyoui@qq.com
// @note  : 打包前端文件dist中间件
package web

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/ginRoute/tool"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type MatchModel uint8

const (
	FuzzyMatching    MatchModel = iota // 模糊匹配
	AccurateMatching                   // 精准匹配
	RegexpMatching                     // 正则匹配

	DefaultRouterHtml string = "index.html"
)

// Rule 规则表
type Rule struct {
	Router   string     // 路由地址
	Resource string     // dist里面的静态资源地址
	Model    MatchModel // 替换模式
}

type autoReplace struct {
	rules []*Rule
	f     func(url string) string
}

// AddRules 增加规则表用来替换资源
func (a *autoReplace) AddRules(rule ...*Rule) {
	a.rules = append(a.rules, rule...)
}

// CustomRuleFunc 增加自定义替换函数
// 优先走自定义替换函数，如果自定义函数返回空字符串，那么就会走规则表（如果规则表被定义的话）
func (a *autoReplace) CustomRuleFunc(f func(url string) (resource string)) {
	a.f = f
}

// NewDist 初始化规则表结构体
//
// LoadFs 挂载go fs里面的内容，请求url首先会去拦截，在去寻找fs里面的资源
// AddRules 增加一些自定义规则
func NewDist() *autoReplace {
	return &autoReplace{}
}

func LoadDistFs(efs embed.FS) gin.HandlerFunc {
	return NewDist().LoadFs(efs)
}

// LoadFs 默认index.html静态文件
func (a *autoReplace) LoadFs(efs embed.FS) gin.HandlerFunc {
	if a.f == nil && a.rules == nil {
		a.AddRules(defaultRule())
	}
	return a.static(efs)
}

// DefaultRule 默认规则，路由从根目录开始
func defaultRule() *Rule {
	r := &Rule{
		Router:   "/",
		Resource: DefaultRouterHtml,
		Model:    AccurateMatching,
	}
	return r
}

// head 根据资源不同的类型响应不同的Content-Type
func head(c *gin.Context, suffix string) {
	typ := ""
	switch strings.ToLower(suffix) {
	case ".html", ".htm", ".css":
		typ = "text/" + suffix[1:]
	case ".js":
		typ = "application/javascript"
	case ".ico":
		typ = "image/x-icon"
	case ".png", ".jpg", ".jpeg":
		typ = "image/" + suffix[1:]
	case ".woff", ".woff2":
		typ = "font/" + suffix[1:]
	default:
		c.Header("Content-Type", "application/json; charset=UTF-8")
		return
	}
	c.Header("Content-Type", typ)
}

// 根据替换模式进行不同的规则
func (m MatchModel) match(url, router string) (ok bool) {
	switch m {
	case FuzzyMatching:
		ok = strings.Contains(url, router)
	case AccurateMatching:
		ok = url == router
	case RegexpMatching:
		ok = regexp.MustCompile(router).MatchString(url)
	default:
		ok = false
	}
	return
}

/*
	muxPath 根据不同的参数来获取不同的静态网页地址
	paths { "/" : "index.html" , "/home" : "home.html" }
	比如： / --> index.html  /home --> home.html
	如果路由获取不到自定义资源名称，返回路由地址
*/
func (a *autoReplace) muxPath(url string) string {
	if a.f != nil {
		resource := a.f(url)
		if resource != "" {
			return resource
		}
	}
	for _, auto := range a.rules {
		if auto.Model.match(url, auto.Router) {
			return auto.Resource
		}
	}
	return url
}

// realValidPath 获取真实
func realValidPath(dir string, address string) string {
	// 拼接地址
	root := filepath.Join(dir, address)

	// 全部的路径符号需要将\转为/
	root = tool.ReplaceSepByFS(root)

	return root
}

// fileResource 搜索资源路径，如果搜索到资源返回真实路径地址，否则返回空字符串
// 搜索资源的算法：先按照路由的地址去搜索，搜索不到在按到资源的名称去搜索
// 比如 路由地址：/css/index.css 先去css文件夹搜索，css文件夹搜索不到在去其它文件夹搜索
func fileResource(efs *embed.FS, html string, dir string) string {
	if dirs, err := efs.ReadDir(dir); err == nil {
		// 获取真实有效的资源地址
		root := realValidPath(dir, html)

		// 判断资源地址是否存在
		if _, err := existFile(efs, root); err == nil {
			return root
		}

		// 不存在在去递归寻找
		for _, address := range dirs {
			if address.IsDir() {
				// 这一步相当重要，如果不用拼接校验有效的路由，会出现意想不到的错误
				newDir := realValidPath(dir, address.Name())
				return fileResource(efs, html, newDir)
			} else {
				// 获取资源的名称
				name := path.Base(html)
				if address.Name() == name {
					return dir + "/" + name
				}
			}
		}
	}
	return ""
}

// existFile 判断是否存在改文件以及是否具有权限
func existFile(efs *embed.FS, staticFile string) (file fs.File, err error) {
	file, err = efs.Open(staticFile)
	return
}

// static 多路径静态文件
func (a *autoReplace) static(efs embed.FS) gin.HandlerFunc {
	agent := func(c *gin.Context) {
		// 获取前段网页的路由地址
		url := strings.TrimSpace(c.Request.URL.Path)

		// 根据paths映射表，从路由地址获取静态网页名称
		// 如果不存在直接返回路由
		html := a.muxPath(url)

		// 判断静态资源是否存在,存在就返回真实路径
		// 不存在返回空字符串
		staticFile := fileResource(&efs, html, ".")
		if staticFile == "" {
			return
		}

		// 再次读取资源，如果读取不到，立即结束
		file, err := existFile(&efs, staticFile)
		if err != nil {
			fmt.Println("寻找到静态资源，但是无法获取资源信息：" + err.Error())
			return
		}

		// 根据静态资源或者路由名称来获取后缀名称
		// 比如index.html -> html , index.js -> js
		suffix := path.Ext(html)

		// 根据后缀名称来封住不同的响应Content-Type
		head(c, suffix)

		// 获取文件的状态，判断是否是可读的
		stat, err := file.Stat()
		if err != nil {
			panic("获取静态资源状态错误：" + err.Error())
			return
		}

		// 判断是否是文件
		if !stat.IsDir() {
			// 读取文件内容
			text, err := efs.ReadFile(staticFile)
			if err != nil {
				panic("读取静态资源内容失败：" + err.Error())
				return
			} else {
				// 将文件内容写回去
				c.Status(http.StatusOK)
				_, err = c.Writer.Write(text)
				if err != nil {
					fmt.Println("静态资源发送失败：" + err.Error())
					return
				}
				c.Abort()
			}
		}
	}
	return agent
}
