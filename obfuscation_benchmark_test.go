package hide

import "testing"

func BenchmarkInt32Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Int32Obfuscate(int32(i), nil)
	}
}

func BenchmarkInt64Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Int64Obfuscate(int64(i), nil)
	}
}

func BenchmarkUint32Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Uint32Obfuscate(uint32(i), nil)
	}
}

func BenchmarkUint64Obfuscate(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Uint64Obfuscate(uint64(i), nil)
	}
}
