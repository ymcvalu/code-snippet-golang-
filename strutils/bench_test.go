package strutils

import (
	"testing"
)

func BenchmarkS2B(b *testing.B) {
	s := "..................................................................."
	b.Run("s2b hack", s2b(s))
	b.Run("s2b common", s2bCommon(s))
}

func s2b(s string) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			S2B(s)
		}
	}
}

func s2bCommon(s string) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_s2bCommon(s)
		}
	}
}

func _s2bCommon(s string) []byte {
	return []byte(s)
}

func BenchmarkB2S(b *testing.B) {
	byts := []byte("....................................................................")
	b.Run("b2s hack", b2s(byts))
	b.Run("b2s common", b2sCommon(byts))
}

func b2s(byts []byte) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B2S(byts)
		}
	}
}

func b2sCommon(byts []byte) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_b2sCommon(byts)
		}
	}
}

func _b2sCommon(byts []byte) string {
	return string(byts)
}
