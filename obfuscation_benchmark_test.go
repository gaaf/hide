package hide_test

import (
	"testing"

	. "github.com/c2h5oh/hide"
)

func BenchmarkInt32Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Default.Int32Obfuscate(int32(i))
	}
}

func BenchmarkInt64Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Default.Int64Obfuscate(int64(i))
	}
}

func BenchmarkUint32Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Default.Uint32Obfuscate(uint32(i))
	}
}

func BenchmarkUint64Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Default.Uint64Obfuscate(uint64(i))
	}
}
