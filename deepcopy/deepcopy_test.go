package deepcopy

import (
	"testing"
)

type Value struct {
	Value string
}

type Src struct {
	String   string
	Pstring  *string
	Int      int
	Pint     *int
	Float64  float64
	Pfload64 *float64
	Struct   Value
	Pstruct  *Value
}

type Dst struct {
	String   string
	Pstring  *string
	Int      int
	Pint     *int
	Float64  float64
	Pfload64 *float64
	Struct   Value
	Pstruct  *Value
}

func TestCopy(t *testing.T) {
	ps := "pcopy"
	pi := -99
	pf := -3.14
	pv := Value{Value: "p-struct-value"}

	src := Src{
		String:   "copy",
		Pstring:  &ps,
		Int:      99,
		Pint:     &pi,
		Float64:  3.14,
		Pfload64: &pf,
		Struct:   Value{Value: "struct-value"},
		Pstruct:  &pv,
	}

	dst := Dst{}

	err := Copy(&src, &dst)
	if err != nil {
		t.Fatal(err)
	}

	if src.String != dst.String || *src.Pstring != *dst.Pstring ||
		src.Int != dst.Int || *src.Pint != *dst.Pint ||
		src.Float64 != dst.Float64 || *src.Pfload64 != *dst.Pfload64 ||
		src.Struct != dst.Struct || *src.Pstruct != *dst.Pstruct {
		t.Fatal("copy failed")
	}
}
