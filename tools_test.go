package go_tools

import "testing"

//BenchmarkRandStr-6   	  129830	      9590 ns/op
func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(12)
	}
}

func TestRandStr(t *testing.T) {
	t.Log(RandStr(6))
	t.Log(RandStr(8))
	t.Log(RandStr(9))
	t.Log(RandStr(10))
}
