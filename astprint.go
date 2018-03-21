package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"go/format"
)

func addInitiate(file *ast.File) {
	file.Decls = append([]ast.Decl{file.Decls[0],
		&ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						{
							Name: "count",
						},
					},
					Values: []ast.Expr {
						&ast.BasicLit{
							Kind: token.INT,
							Value: "0",
						},
					},
				},
			},
		}, }, (file.Decls[1:])...,
	)
}

func addPrintf(file *ast.File, f *ast.FuncDecl) {
	f.Body.List = append(f.Body.List, ast.Stmt(
		&ast.ExprStmt{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					ast.NewIdent("fmt"),
					ast.NewIdent("Printf"),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						ValuePos: file.Pos(),
						Kind: token.STRING,
						Value: "\"Count of Assignments = \" + fmt.Sprint(count)" ,
					},
				},
			},
		}, ),
	)
}

func addIncs(f *ast.FuncDecl, count int) {
	for i :=0; i < count; i++ {
		f.Body.List = append([]ast.Stmt{ast.Stmt(
			&ast.IncDecStmt{
				X: &ast.Ident{
					Name: "count",
				},
				Tok: token.INC,
			},
		)}, f.Body.List...)
	}
}

func addInc(curStmt *ast.BlockStmt) {
	curStmt.List = append([]ast.Stmt{ast.Stmt(
		&ast.IncDecStmt{
			X: &ast.Ident{
				Name:    "count",
			},
			Tok:    token.INC,
		},
	)}, curStmt.List...)
}

func countAssignments(file *ast.File) {
	addInitiate(file)
	var curStmt *ast.BlockStmt
	var count = 0
	var flag = true
	ast.Inspect(file, func(node ast.Node) bool {
		if f, ok := node.(*ast.FuncDecl); ok && f.Name.Name == "main" {
			addPrintf(file, f)
			addIncs(f, count-1)

		}
		if f, ok := node.(*ast.BlockStmt); ok {
			curStmt = f
		}

		if f, ok := node.(*ast.ValueSpec); ok {
			if len(f.Values) != 0 && flag {
				count++;
			}
		}

		if _, ok := node.(*ast.FuncDecl); ok {
			flag = false
		}

		_, ok := node.(*ast.AssignStmt)
		_, ok2 := node.(*ast.DeclStmt)

		if ok || ok2 {
			addInc(curStmt)
		}

		return true
	})

}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: astprint <filename.go>\n")
		return
	}

	fset := token.NewFileSet()

	file, err := parser.ParseFile(
		fset, os.Args[1], nil, parser.ParseComments,
	); if err == nil {
		ast.Fprint(os.Stdout, fset, file, nil)
	} else {
		fmt.Printf("Error: %v", err)
	}
	countAssignments(file)
	if dst, err := os.Create("result/res.go"); err == nil {
		format.Node(dst, fset, file)
	} else {
		fmt.Printf("Error: %v", err)
	}
}