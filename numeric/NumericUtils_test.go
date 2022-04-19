package numeric

import (
	"testing"
)

func TestToFixedDecimal(t *testing.T) {
	num := 1.5346
	t.Log(ToFixedDecimal(num, 3)) // print 1.525
	t.Log(ToFixedDecimal(num, 0)) //print 2
}

func TestFloorToFixedDecimal(t *testing.T) {
	num := 1.5346
	t.Log(FloorToFixedDecimal(num, 3)) //print 1.534
	t.Log(FloorToFixedDecimal(num, 0)) //print 1
}

func TestCeilToFixedDecimal(t *testing.T) {
	num := 1.0341
	t.Log(CeilToFixedDecimal(num, 3)) //print 1.035
	t.Log(CeilToFixedDecimal(num, 0)) //print 2
}

func TestToString(t *testing.T) {
	t.Log(ToString(12.2121219))
	t.Log(ToString("21212"))
	t.Log(ToString(12))
}

func TestToFloat64(t *testing.T) {
	t.Log(ToFloat64("1.2301"))
	t.Log(ToFloat64(1.23012))
	t.Log(ToFloat64(123))
	t.Log(ToFloat64("ab")) //strconv.ParseFloat: parsing "ab": invalid syntax
	t.Log(ToFloat64(""))
}

func TestToInt64(t *testing.T) {
	t.Log(ToInt64("123"))
	t.Log(ToInt64(1234))
	t.Log(ToInt64("1.23")) //strconv.ParseInt: parsing "1.23": invalid syntax
	t.Log(ToInt64(1.23))   //strconv.ParseInt: parsing "1.23": invalid syntax
}

func TestToUint64(t *testing.T) {
	t.Log(ToUint64(1233))
	t.Log(ToUint64("1234"))
	t.Log(ToUint64(1.23))  //strconv.ParseUint: parsing "1.23": invalid syntax
	t.Log(ToUint64(-1.23)) //strconv.ParseUint: parsing "-1.23": invalid syntax
}
