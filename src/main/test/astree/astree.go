package main

import (
	"flag"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/util"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var proxyDic string
var outDic string
var workspace string

func main() {
	var proxyFile string

	flag.StringVar(&proxyDic, "d", "", "-d proxy dictionary")
	flag.StringVar(&proxyFile, "f", "iuserservice.go", "-f proxy file")
	flag.StringVar(&outDic, "o", "", "-o dictionary to out put new file")
	flag.Parse()

	workspace = getCurrentAbPath()
	if outDic == "" {
		outDic = workspace
	}

	if proxyFile != "" {
		if !filepath.IsAbs(proxyFile) {
			proxyFile = filepath.Join(workspace, proxyFile)
		}
		GenerateByFile(proxyFile)
	} else if proxyDic != "" {
		if !filepath.IsAbs(proxyDic) {
			proxyDic = filepath.Join(workspace, proxyDic)
		}
		GenerateByDic()
	}

}

func GenerateByFile(proxyFile string) {
	var file fs.FileInfo
	var err error
	if file, err = os.Stat(proxyFile); err == nil {
	} else {
		panic(fmt.Errorf("%s can not open", proxyFile))
	}

	var result *ProxyInterface = &ProxyInterface{}
	result.TargetFile = file.Name()
	result.TargetUri = proxyFile
	//result.TargetClazz = ""
	//result.TargetPackage = ""
	result.TargetAnno = make([]string, 0)
	result.TargetImport = make([]string, 0)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, proxyFile, nil, parser.ParseComments)

	// 把节点和comment 映射起来
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		//var s string
		switch x := n.(type) {
		case *ast.GenDecl:
			if len(x.Specs) > 0 {
				l1 := x.Specs[0]
				if ts, ok := l1.(*ast.TypeSpec); ok {
					if ts.Type != nil {
						if _, ok1 := ts.Type.(*ast.InterfaceType); ok1 {
							//fmt.Println(ts.Name.Name)
							result.TargetClazz = ts.Name.Name
							if comments, ok1 := cmap[n]; ok1 {
								if len(comments) > 0 && len(comments[0].List) > 0 {
									for _, comment := range comments[0].List {
										result.TargetAnno = append(result.TargetAnno, comment.Text)
									}
								}
							}
						}
					}
				}
			}
		case *ast.File:
			result.TargetPackage = x.Name.Name
		case *ast.ImportSpec:
			if x.Path.Value != "" {
				result.TargetImport = append(result.TargetImport, strings.ReplaceAll(x.Path.Value, "\"", ""))
			}
		case *ast.InterfaceType:
			for _, m := range x.Methods.List {
				methodName := m.Names[0].Name
				if !isFirstCharUpperCase(methodName) {
					continue
				}
				pm := &ProxyMethod{}
				pm.MethodName = methodName
				pm.TargetAnno = make([]string, 0)
				pm.ParamField = make([]*ProxyField, 0)
				pm.ReturnField = make([]*ProxyField, 0)
				result.Method = append(result.Method, pm)

				if m.Doc != nil {
					for _, c := range m.Doc.List {
						pm.TargetAnno = append(pm.TargetAnno, c.Text)
					}
				}
				fu := m.Type.(*ast.FuncType)

				if fu.Results != nil {
					for _, f := range fu.Results.List {
						pf := getProxyField(f, cmap)
						pm.ReturnField = append(pm.ReturnField, pf)
					}
				}

				if fu.Params != nil {
					for _, f := range fu.Params.List {
						pf := getProxyField(f, cmap)
						pm.ParamField = append(pm.ParamField, pf)
					}
				}
			}
		}
		return true
	})
	fmt.Println(util.JsonUtil.To2String(result))
	targetFile := strings.ReplaceAll(result.TargetFile, ".go", "_proxy.go")
	targetUri := filepath.Join(outDic, targetFile)
	_, proxyPackage := filepath.Split(outDic)
	fmt.Println(outDic, targetUri, proxyPackage)
	result.ProxyPackage = proxyPackage
	result.ProxyClazz = fmt.Sprintf("%sProxy", result.TargetClazz)
	result.ProxyInstance = firstLower(result.ProxyClazz)

	err = generateProxyFile(result, targetUri)
	if err != nil {
		panic(err)
	}
}
func GenerateByDic() {

}
