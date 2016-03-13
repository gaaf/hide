package hide

import (
	"math/rand"
	"testing"
)

func TestInt32Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := rand.Int31()
		r := Int32Deobfuscate(Int32Obfuscate(v, nil), nil)
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}

func TestUint32Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := uint32(rand.Int31() * (1 + rand.Int31n(1)))
		r := Uint32Deobfuscate(Uint32Obfuscate(v, nil), nil)
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}

func TestInt64Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := rand.Int63()
		r := Int64Deobfuscate(Int64Obfuscate(v, nil), nil)
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}

func TestUint64Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := uint64(rand.Int63() * (1 + rand.Int63n(1)))
		r := Uint64Deobfuscate(Uint64Obfuscate(v, nil), nil)
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}
