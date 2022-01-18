// Package tool
// @Time  : 2022/1/17 上午9:40
// @Author: Jtyoui@qq.com
// @note  : 扫码文件工具类
package tool

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var DefaultName = "__local__"

type Method struct {
	Comments   []string
	MethodName string
	StructName string
	Params     []string
}

type GoFileScanner struct {
	methods map[string][]Method
}

func NewGoFileScanner() *GoFileScanner {
	return &GoFileScanner{
		methods: make(map[string][]Method),
	}
}

func (s *GoFileScanner) ParseFile(filename string) error {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, d := range astFile.Decls {
		switch specDecl := d.(type) {
		case *ast.FuncDecl:
			method := Method{}
			method.StructName = DefaultName
			if specDecl.Recv != nil {
				exp, ok := specDecl.Recv.List[0].Type.(*ast.StarExpr)
				if ok {
					method.StructName = fmt.Sprint(exp.X)
				}
			}
			method.MethodName = specDecl.Name.Name
			method.Comments = s.parserComments(specDecl)
			method.Params = s.parserParams(specDecl)
			s.methods[method.StructName] = append(s.methods[method.StructName], method)
		}
	}

	return nil
}

func (s *GoFileScanner) GetMethods(name string) []Method {
	return s.methods[name]
}

func (s *GoFileScanner) parserComments(f *ast.FuncDecl) []string {
	if f.Doc == nil {
		return nil
	}
	var comments []string
	for _, l := range f.Doc.List {
		comments = append(comments, l.Text)
	}
	return comments
}
func (s *GoFileScanner) parserParams(f *ast.FuncDecl) []string {
	var params []string
	for _, l := range f.Type.Params.List {
		params = append(params, l.Names[0].Name)
	}
	return params
}
