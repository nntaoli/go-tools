package numeric

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"strconv"
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
	case float64:
		fNumber := v.(float64)
		fNumber = math.Round(fNumber*100000000.0) / 100000000.0
		return strconv.FormatFloat(fNumber, 'f', -1, 64) //StripTrailingZeros
	case float32:
		fNumber := v.(float32)
		fNumber2 := math.Round(float64(fNumber*100000000.0)) / 100000000.0
		return strconv.FormatFloat(fNumber2, 'f', -1, 64) //StripTrailingZeros
	default:
		return fmt.Sprint(v)
	}

}

//see:	github.com/spf13/cast

//func ToFloat64(v interface{}) (float64, error) {
//	return strconv.ParseFloat(fmt.Sprint(v), 64)
//}
//
//func ToInt64(v interface{}) (int64, error) {
//	return strconv.ParseInt(fmt.Sprint(v), 10, 64)
//}
//
//func ToUint64(v interface{}) (uint64, error) {
//	return strconv.ParseUint(fmt.Sprint(v), 10, 64)
//}

// BigMul
//两个大数相乘，防止溢出
func BigMul(a, b int64) string {
	var r = big.NewInt(0)
	return r.Mul(big.NewInt(a), big.NewInt(b)).String()
}

// RandN
// Generate a random number between min and max
func RandN(min, max uint64) uint64 {
	var (
		rndBytes []byte
		n        int
		err      error
	)

	rndBytes = make([]byte, 8)

re:
	n, err = rand.Read(rndBytes)
	if err != nil || n < 8 {
		goto re
	}

	return binary.BigEndian.Uint64(rndBytes)%(max-min) + min
}
