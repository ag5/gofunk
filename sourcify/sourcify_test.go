package sourcify

import (
	"github.com/ag5/gofunk/assertor"
	"testing"
)

func TestConstant(t *testing.T) {
	assertor.AssertEquals(t, "1", PrintExpr(AnyToExpr(1)))
	assertor.AssertEquals(t, "0", PrintExpr(AnyToExpr(0)))
	assertor.AssertEquals(t, "true", PrintExpr(AnyToExpr(true)))
	assertor.AssertEquals(t, "false", PrintExpr(AnyToExpr(false)))
}

func TestLiteralArray(t *testing.T) {
	arr := []int{1, 2, 3}
	assertor.AssertEquals(t, "[]int{1, 2, 3}", PrintExpr(AnyToExpr(arr)))
}

type Struct1 struct {
	Foo string
	Bar int
}

func TestStruct1(t *testing.T) {
	str := Struct1{Foo: "Tomaat", Bar: 7}
	assertor.AssertEquals(t, "Struct1{Foo: \"Tomaat\", Bar: 7}", PrintExpr(AnyToExpr(str)))
	strp := &Struct1{Foo: "Tomaat", Bar: 7}
	assertor.AssertEquals(t, "&Struct1{Foo: \"Tomaat\", Bar: 7}", PrintExpr(AnyToExpr(strp)))

	// Here 'Bar' has a default value; we expect the field to be skipped:
	str = Struct1{Foo: "Tomaat"}
	assertor.AssertEquals(t, "Struct1{Foo: \"Tomaat\"}", PrintExpr(AnyToExpr(str)))
}
