package hide_test

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	. "github.com/c2h5oh/hide"
)

func init() {
	if err := Default.SetInt32(new(big.Int).SetInt64(1580030173)); err != nil {
		panic(err)
	}
	if err := Default.SetInt64(new(big.Int).SetInt64(8230452606740808761)); err != nil {
		panic(err)
	}

	if err := Default.SetUint32(new(big.Int).SetUint64(1500450271)); err != nil {
		panic(err)
	}
	if err := Default.SetUint64(new(big.Int).SetUint64(12764787846358441471)); err != nil {
		panic(err)
	}

	if err := Default.SetXor(new(big.Int).SetUint64(3469983624777167712)); err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())
}

func TestInt32Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := rand.Int31()
		r := Default.Int32Deobfuscate(Default.Int32Obfuscate(v))
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}

func TestUint32Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := uint32(rand.Int31() * (1 + rand.Int31n(1)))
		r := Default.Uint32Deobfuscate(Default.Uint32Obfuscate(v))
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}

func TestInt64Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := rand.Int63()
		r := Default.Int64Deobfuscate(Default.Int64Obfuscate(v))
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}

func TestUint64Reversible(t *testing.T) {
	for i := 0; i < 100000; i++ {
		v := uint64(rand.Int63() * (1 + rand.Int63n(1)))
		r := Default.Uint64Deobfuscate(Default.Uint64Obfuscate(v))
		if v != r {
			t.Logf("Expected %d, actual %d", v, r)
			t.Fail()
		}
	}
}
