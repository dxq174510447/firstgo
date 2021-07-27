package main

import (
	"flag"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/util"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var proxyDic string
var proxyFile string
var outDic string

var workspace string

type ProxyInterface struct {
	// 目录包
	TargetPackage string
	// 目标类名
	TargetClazz string
	// 目标文件名
	TargetFile string
	// 目标文件绝对路径
	TargetUri string

	// 引入的包
	TargetImport []string

	// 目标对象注释
	TargetAnno []string

	// 方法
	Method []*ProxyMethod
}

type ProxyMethod struct {
	MethodName  string
	TargetAnno  []string
	ParamField  []*ProxyField
	ReturnField []*ProxyField
}

type ProxyField struct {
	FieldName   string
	FieldType   string
	TypePackage string
	TargetAnno  []string
	// 是否是指针 0否1是
	FieldAddress int
	// 是否是slice 0否1是
	FieldArray int
	// 是否是map 0否1是
	FieldMap   int
	KeyField   *ProxyField
	ValueField *ProxyField
}

func main() {
	flag.StringVar(&proxyDic, "d", "1", "-d proxy dictionary")
	flag.StringVar(&proxyFile, "f", "iuserservice.go", "-f proxy file")
	flag.StringVar(&outDic, "o", "2", "-o dictionary to out put new file")
	flag.Parse()

	workspace = getCurrentAbPath()

	if proxyFile != "" {
		if !filepath.IsAbs(proxyFile) {
			proxyFile = filepath.Join(workspace, proxyFile)
		}
		GenerateByFile()
	} else if proxyDic != "" {
		if !filepath.IsAbs(proxyDic) {
			proxyDic = filepath.Join(workspace, proxyDic)
		}
		GenerateByDic()
	}

}

func GenerateByFile() {
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
						pf := getProxyField(f)
						pm.ReturnField = append(pm.ReturnField, pf)
					}
				}

				if fu.Params != nil {
					for _, f := range fu.Params.List {
						pf := getProxyField(f)
						pm.ParamField = append(pm.ParamField, pf)
					}
				}
			}
		}
		return true
	})
	fmt.Println(util.JsonUtil.To2String(result))
}
func GenerateByDic() {

}

// https://tehub.com/a/44BceBfRK0
// getCurrentAbPath 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func isFirstCharUpperCase(content string) bool {
	firstChar := content[0]
	if firstChar < 65 || firstChar > 90 {
		return false
	} else {
		return true
	}
}

func getProxyField(f *ast.Field) *ProxyField {
	var fieldName string
	var fieldType string
	var typePackage string
	var fieldAddress int = 0 //0否1是
	var fieldArray int = 0
	var fieldMap int = 0

	if len(f.Names) > 0 {
		fieldName = f.Names[0].Name
	}

	switch field := f.Type.(type) {
	case *ast.MapType:
		fieldMap = 1

		keyField := &ProxyField{}
		switch f1 := field.Key.(type) {
		case *ast.SelectorExpr:
			fieldType = f1.Sel.Name
			if f1.X != nil {
				tp := f1.X.(*ast.Ident)
				typePackage = tp.Name
			}
		case *ast.Ident:
			fieldType = f1.Name
		case *ast.StarExpr:
			fieldAddress = 1
			if f1.X != nil {
				tp := f1.X.(*ast.SelectorExpr)
				fieldType = tp.Sel.Name

				tp1 := tp.X.(*ast.Ident)
				typePackage = tp1.Name
			}
		}

	case *ast.ArrayType:
		fieldArray = 1
		switch f1 := field.Elt.(type) {
		case *ast.SelectorExpr:
			fieldType = f1.Sel.Name
			if f1.X != nil {
				tp := f1.X.(*ast.Ident)
				typePackage = tp.Name
			}
		case *ast.Ident:
			fieldType = f1.Name
		case *ast.StarExpr:
			fieldAddress = 1
			if f1.X != nil {
				tp := f1.X.(*ast.SelectorExpr)
				fieldType = tp.Sel.Name

				tp1 := tp.X.(*ast.Ident)
				typePackage = tp1.Name
			}
		}
	case *ast.SelectorExpr:
		fieldType = field.Sel.Name
		if field.X != nil {
			tp := field.X.(*ast.Ident)
			typePackage = tp.Name
		}
	case *ast.Ident:
		fieldType = field.Name
	case *ast.StarExpr:
		fieldAddress = 1
		if field.X != nil {

			switch f1 := field.X.(type) {
			case *ast.SelectorExpr:
				fieldType = f1.Sel.Name
				if f1.X != nil {
					tp := f1.X.(*ast.Ident)
					typePackage = tp.Name
				}
			case *ast.Ident:
				fieldType = f1.Name
			}
		}
	}
	pf := &ProxyField{}
	pf.FieldName = fieldName
	pf.FieldType = fieldType
	pf.TypePackage = typePackage
	pf.FieldAddress = fieldAddress
	pf.FieldArray = fieldArray
	pf.FieldMap = fieldMap
	return pf
}
