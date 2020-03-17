package crypto

import (
	"fmt"
	"testing"
)

func BenchmarkGetMd5String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("md5%d", i)
		GetMd5String(str)
	}
}

func BenchmarkGetSM3String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf("sm3%d", i)
		GetSM3String(str)
	}
}
