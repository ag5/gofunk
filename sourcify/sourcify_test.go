package sourcify

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConstant(t *testing.T) {
	var expected any = "1"
	var got any = PrintExpr(AnyToExpr(1))
	require.Equal(t, expected, got)
	var expected2 any = "0"
	var got2 any = PrintExpr(AnyToExpr(0))
	require.Equal(t, expected2, got2)
	var expected3 any = "true"
	var got3 any = PrintExpr(AnyToExpr(true))
	require.Equal(t, expected3, got3)
	var expected4 any = "false"
	var got4 any = PrintExpr(AnyToExpr(false))
	require.Equal(t, expected4, got4)
}

func TestLiteralArray(t *testing.T) {
	arr := []int{1, 2, 3}
	var expected any = "[]int{1, 2, 3}"
	var got any = PrintExpr(AnyToExpr(arr))
	require.Equal(t, expected, got)
}

type Struct1 struct {
	Foo string
	Bar int
}

func TestStruct1(t *testing.T) {
	str := Struct1{Foo: "Tomaat", Bar: 7}
	var expected any = "sourcify.Struct1{Foo: \"Tomaat\", Bar: 7}"
	var got any = PrintExpr(AnyToExpr(str))
	require.Equal(t, expected, got)
	strp := &Struct1{Foo: "Tomaat", Bar: 7}
	var expected2 any = "&sourcify.Struct1{Foo: \"Tomaat\", Bar: 7}"
	var got2 any = PrintExpr(AnyToExpr(strp))
	require.Equal(t, expected2, got2)

	// Here 'Bar' has a default value; we expect the field to be skipped:
	str = Struct1{Foo: "Tomaat"}
	var expected3 any = "sourcify.Struct1{Foo: \"Tomaat\"}"
	var got3 any = PrintExpr(AnyToExpr(str))
	require.Equal(t, expected3, got3)
}
