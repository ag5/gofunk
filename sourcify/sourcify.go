package sourcify

import (
	"github.com/ag5/gofunk/funk"
	"go/ast"
	"go/token"
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
	if !val.IsValid() {
		return nil
	}
	typ := val.Type()
	switch typ.Kind() {
	case reflect.Bool:
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return &ast.BasicLit{
			Kind:  token.INT,
			Value: val.String(),
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
			Type: &ast.ArrayType{
				Len: &ast.BasicLit{Kind: token.INT,
					Value: strconv.Itoa(val.Len())},
				Elt: &ast.Ident{
					Name: typ.Elem().Name(),
				},
			},
			Elts: funk.MapRange(0, val.Len(), func(i int) ast.Expr {
				return AnyToExpr(val.Index(i).Interface())
			}),
		}

	case reflect.Chan:
		panic("NIY")
	case reflect.Func:
		panic("NIY")
	case reflect.Interface:
		if !val.Elem().IsNil() {
			AnyToExpr(val.Elem().Interface())
		}
	case reflect.Map:
		panic("NIY")
	case reflect.Pointer:
		if !val.IsNil() {
			return &ast.StarExpr{
				X: AnyToExpr(val.Elem().Interface()),
			}
		}
		return nil
	case reflect.Slice:
		return &ast.CompositeLit{
			Type: ,
			Elts: funk.MapRange(0, val.Len(), func(i int) ast.Expr {
				return AnyToExpr(val.Index(i).Interface())
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
					Value: AnyToExpr(val.Field(i).Interface()),
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
	case reflect.Array:
		return &ast.ArrayType{
			Len: &ast.BasicLit{Kind: token.INT,
				Value: strconv.Itoa(typ.Len())},
			Elt: TypeAsExpr(typ.Elem()),
		}
	case reflect.Slice:
		return &ast.ArrayType{
,
			Elt: TypeAsExpr(typ.Elem())
		}
	case reflect.Struct:
		return &ast.Ident{
			Name: typ.Elem().Name(),
		}
	}
	return nil
}
