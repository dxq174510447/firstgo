package astree

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAst1(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "astree.go", nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		//var s string
		switch x := n.(type) {
		case *ast.InterfaceType:
			fmt.Println("interface", x.Pos(), x.Interface, x.End())
			for _, m := range x.Methods.List {
				if m.Doc != nil {
					for _, c := range m.Doc.List {
						fmt.Println(c.Text)
					}
				}
				fu := m.Type.(*ast.FuncType)

				for _, f := range fu.Params.List {
					fmt.Println(f.Names[0].String())
				}
				fmt.Println(fu.Results.NumFields())
				for _, f := range fu.Results.List {
					id := f.Type.(*ast.Ident)
					fmt.Println(id.Name, id.Obj)
				}
				for _, f := range m.Names {
					fmt.Println(f.String(), f.Obj.Type)
				}
			}
		case *ast.StructType:
			fmt.Println("struct", x.Pos(), x.Struct, x.End())
		case *ast.BasicLit:
			//s = x.Value
		case *ast.Ident:
			//s = x.Name
		}
		//if s != "" {
		//	fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		//}
		return true
	})

}
