package numeric

import (
	"fmt"
	"math"
)

//ToFixedDecimal
/**
 * 四舍五入
 * @param n 要保留的小数点位数
 */
func ToFixedDecimal(num float64, n int) string {
	return fmt.Sprintf("%."+fmt.Sprint(n)+"f", num)
}

//FloorToFixedDecimal
//eg. FloorToFixedDecimal(0.125,2) return "0.12"
func FloorToFixedDecimal(num float64, n int) string {

	var factor float64 = 1
	for i := n; i > 0; i-- {
		factor = factor * 10
	}

	return ToFixedDecimal(math.Floor(num*factor)/factor, n)
}

// CeilToFixedDecimal
// eg. CeilToFixedDecimal(0.125 ,1) return "0.2"
func CeilToFixedDecimal(num float64, n int) string {
	var factor float64 = 1
	for i := n; i > 0; i-- {
		factor = factor * 10
	}
	return ToFixedDecimal(math.Ceil(num*factor)/factor, n)
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}

	switch v.(type) {
	case float64, float32:
		return fmt.Sprintf("%.6f", v)
	default:
		return fmt.Sprint(v)
	}
}
