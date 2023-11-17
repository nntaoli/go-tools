package numeric

import (
	"math"
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
	var a float32
	a = 1.02219999999998
	t.Log(ToString(a))
	t.Log(ToString(1000))
	t.Log(ToString("a123"))
	t.Log(ToString(nil))
}

func TestBigMul(t *testing.T) {
	t.Log(BigMul(math.MaxInt32, math.MaxInt64))
}

func TestRandNum(t *testing.T) {
	t.Log(RandN(10000, 99999))
}

func BenchmarkRandNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandN(10000, 99999)
	}
}
