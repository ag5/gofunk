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

func ToSource(obj any) string {
	w := &strings.Builder{}
	AnyToExpr(obj)

	return w.String()
}

func crIndent(w *strings.Builder, ind int) {
	w.WriteString("\n")
	for i := 0; i < ind; i++ {
		w.WriteString("    ")
	}
}

func AnyToExpr(obj any) ast.Expr {
	val := reflect.ValueOf(obj)
	return ValueToExpr(val)
}

func ValueToExpr(val reflect.Value) ast.Expr {
	if !val.IsValid() {
		return nil
	}
	typ := val.Type()
	switch typ.Kind() {
	case reflect.Bool:
		return &ast.Ident{
			Name: strconv.FormatBool(val.Bool()),
		}
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
				return ValueToExpr(val.Index(i))
			}),
		}

	case reflect.Chan:
		panic("NIY")
	case reflect.Func:
		panic("NIY")
	case reflect.Interface:
		if !val.Elem().IsNil() {
			ValueToExpr(val.Elem())
		}
	case reflect.Map:
		panic("NIY")
	case reflect.Pointer:
		if !val.IsNil() {
			return &ast.UnaryExpr{
				Op: token.AND,
				X:  ValueToExpr(val.Elem()),
			}
		}
		return nil
	case reflect.Slice:
		return &ast.CompositeLit{
			Type: TypeAsExpr(typ),
			Elts: funk.MapRange(0, val.Len(), func(i int) ast.Expr {
				return ValueToExpr(val.Index(i))
			}),
		}
	case reflect.String:
		str := val.Convert(reflect.TypeOf("string")).Interface().(string)
		return &ast.BasicLit{Kind: token.STRING,
			Value: strconv.Quote(str)}
	case reflect.Struct:
		return &ast.CompositeLit{
			Type: TypeAsExpr(typ),

			Elts: funk.MapRange(0, typ.NumField(), func(i int) ast.Expr {
				f := typ.Field(i)
				return &ast.KeyValueExpr{
					Key:   &ast.Ident{Name: f.Name},
					Value: ValueToExpr(val.Field(i)),
				}
			}),
		}

	case reflect.UnsafePointer:
		panic("NIY")
	default:
		panic("Unknown kind")
	}
	return nil
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
		panic("niy")
	case reflect.Map:
		panic("niy")
	case reflect.Pointer:
		panic("niy")
	case reflect.Slice:
		return &ast.ArrayType{
			Elt: TypeAsExpr(typ.Elem()),
		}
	case reflect.String:
		panic("niy")
	case reflect.Struct:
		return &ast.Ident{
			Name: typ.Name(),
		}
	case reflect.UnsafePointer:
		panic("niy")

	default:
		panic("Unknown kind")
	}
	return nil
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
