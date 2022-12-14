package go_tools

import (
	"math/rand"
	"time"
)

var char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ"

//其他经常用的零碎方法集合

//RandStr
//生成一个随机字符串
func RandStr(l int) (rndStr string) {
	var (
		passBytes = make([]byte, l)
	)

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		rd.Read(passBytes)
		if len(passBytes) == l {
			break
		}
	}

	for _, i := range passBytes {
		c := char[i%61]
		rndStr += string(c)
	}

	return
}
