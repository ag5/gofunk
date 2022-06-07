package sourcify

import (
	"bytes"
	"github.com/ag5/gofunk/funk"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func crIndent(w *strings.Builder, ind int) {
	w.WriteString("\n")
	for i := 0; i < ind; i++ {
		w.WriteString("    ")
	}
}

func ToSource(obj any) string {
	return PrintExpr(AnyToExpr(obj))
}

func AnyToExpr(obj any) ast.Expr {
	return ValueToExpr(reflect.ValueOf(obj), DefaultConfig())
}
func DefaultConfig() ValueToExprConfig {
	return ValueToExprConfig{
		SuppressZeroFields: true,
	}
}

type ValueToExprConfig struct {
	SuppressZeroFields bool
}

func ValueToExpr(val reflect.Value, config ValueToExprConfig) ast.Expr {
	//if !val.IsValid() {
	//	return nil
	//}
	typ := val.Type()
	switch typ.Kind() {
	case reflect.Bool:
		return &ast.Ident{Name: strconv.FormatBool(val.Bool())}
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return &ast.BasicLit{
			Kind:  token.INT,
			Value: strconv.FormatInt(val.Int(), 10),
		}
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return &ast.BasicLit{
			Kind:  token.INT,
			Value: strconv.FormatUint(val.Uint(), 10),
		}

	case reflect.Uintptr:
		panic("NIY")
	case reflect.Float32,
		reflect.Float64:
		panic("NIY")

	case reflect.Complex64,
		reflect.Complex128:
		panic("NIY")
	case reflect.Array:
		return &ast.CompositeLit{
			Type: TypeAsExpr(typ),
			Elts: funk.MapRange(0, val.Len(), func(i int) ast.Expr {
				return ValueToExpr(val.Index(i), config)
			}),
		}

	case reflect.Chan:
		panic("NIY")
	case reflect.Func:
		panic("NIY")
	case reflect.Interface:
		if !val.IsZero() && !val.Elem().IsZero() {
			return ValueToExpr(val.Elem(), config)
		}
		if config.SuppressZeroFields {
			return nil
		}
		return &ast.Ident{Name: "nil"}
	case reflect.Map:
		panic("NIY")
	case reflect.Pointer:
		if !val.IsNil() {
			return &ast.UnaryExpr{
				Op: token.AND,
				X:  ValueToExpr(val.Elem(), config),
			}
		}
		if config.SuppressZeroFields {
			return nil
		}
		return &ast.Ident{Name: "nil"}
	case reflect.Slice:
		return &ast.CompositeLit{
			Type: TypeAsExpr(typ),
			Elts: funk.MapRange(0, val.Len(), func(i int) ast.Expr {
				return ValueToExpr(val.Index(i), config)
			}),
		}
	case reflect.String:
		str := val.Convert(reflect.TypeOf("string")).Interface().(string)
		return &ast.BasicLit{Kind: token.STRING,
			Value: strconv.Quote(str)}
	case reflect.Struct:
		return structToExpr(val, typ, config)

	case reflect.UnsafePointer:
		panic("NIY")
	default:
		panic("Unknown kind")
	}
	panic("Should not occur")
}

func structToExpr(val reflect.Value, typ reflect.Type, config ValueToExprConfig) ast.Expr {
	var start = 0
	var stop = typ.NumField()
	var elts []ast.Expr
	for i := start; i < stop; i++ {
		f := typ.Field(i)
		fval := val.Field(i)
		if f.IsExported() {
			if !config.SuppressZeroFields || !fval.IsZero() {

				fvalExpr := ValueToExpr(fval, config)
				if fvalExpr != nil {
					elts = append(elts, &ast.KeyValueExpr{
						Key:   &ast.Ident{Name: f.Name},
						Colon: 0,
						Value: fvalExpr,
					})
				}
			}
		}
	}
	return &ast.CompositeLit{
		Type: TypeAsExpr(typ),
		Elts: elts,
	}
}

func TypeAsExpr(typ reflect.Type) ast.Expr {
	switch typ.Kind() {

	case reflect.Bool:
		return &ast.Ident{Name: "bool"}
	case reflect.Int:
		return &ast.Ident{Name: "int"}
	case reflect.Int8:
		return &ast.Ident{Name: "int8"}
	case reflect.Int16:
		return &ast.Ident{Name: "int16"}
	case reflect.Int32:
		return &ast.Ident{Name: "int32"}
	case reflect.Int64:
		return &ast.Ident{Name: "int64"}
	case reflect.Uint:
		return &ast.Ident{Name: "uint"}
	case reflect.Uint8:
		return &ast.Ident{Name: "uint8"}
	case reflect.Uint16:
		return &ast.Ident{Name: "uint16"}
	case reflect.Uint32:
		return &ast.Ident{Name: "uint32"}
	case reflect.Uint64:
		return &ast.Ident{Name: "uint64"}
	case reflect.Uintptr:
		panic("niy")
	case reflect.Float32:
		panic("niy")
	case reflect.Float64:
		panic("niy")
	case reflect.Complex64:
		panic("niy")
	case reflect.Complex128:
		panic("niy")
	case reflect.Array:
		return &ast.ArrayType{
			Len: &ast.BasicLit{Kind: token.INT,
				Value: strconv.Itoa(typ.Len())},
			Elt: TypeAsExpr(typ.Elem()),
		}
	case reflect.Chan:
		panic("niy")
	case reflect.Func:
		panic("niy")
	case reflect.Interface:
		return typeAsQualifiedName(typ)
	case reflect.Map:
		panic("niy")
	case reflect.Pointer:
		return &ast.StarExpr{X: TypeAsExpr(typ.Elem())}
	case reflect.Slice:
		return &ast.ArrayType{
			Elt: TypeAsExpr(typ.Elem()),
		}
	case reflect.String:
		return &ast.Ident{Name: "string"}
	case reflect.Struct:
		return typeAsQualifiedName(typ)
	case reflect.UnsafePointer:
		panic("niy")

	default:
		panic("Unknown kind")
	}
	return nil
}

func typeAsQualifiedName(typ reflect.Type) *ast.SelectorExpr {
	path := strings.Split(typ.PkgPath(), "/")
	pkgName := path[len(path)-1]
	return &ast.SelectorExpr{X: &ast.Ident{Name: pkgName}, Sel: &ast.Ident{Name: typ.Name()}}
}

func PrintExpr(res ast.Expr) string {
	buf := new(bytes.Buffer)
	fset := token.NewFileSet()
	err := format.Node(buf, fset, res)
	if err != nil {
		log.Fatal(err)
	}
	string := buf.String()
	return string
}
