package hide

import (
	"errors"
	"math/big"
)

// Hide stores primes and inverses used to obfuscate/deobfuscate different
// integer types
type Hide struct {
	int32prime    *big.Int
	int32inverse  *big.Int
	int64prime    *big.Int
	int64inverse  *big.Int
	uint32prime   *big.Int
	uint32inverse *big.Int
	uint64prime   *big.Int
	uint64inverse *big.Int
}

var (
	Default Hide

	bigOne = big.NewInt(1)

	// maximum value for each type
	int32Max  = big.NewInt(2147483647)                       // 2^31-1
	int64Max  = big.NewInt(9223372036854775807)              // 2^63 -1
	uint32Max = new(big.Int).SetUint64(4294967295)           // 2^32 -1
	uint64Max = new(big.Int).SetUint64(18446744073709551615) // 2^64 -1

	ErrNil        = errors.New("prime is nil")
	ErrOutOfRange = errors.New("prime is greater than max value for the type")
	ErrNotAPrime  = errors.New("it is not a prime number")
)

func modularMultiplicativeInverse(val, prime, max *big.Int) {
	if prime == nil {
		panic(ErrNil)
	}

	val.Mul(val, prime)
	val.And(val, max)
}

func (h *Hide) SetInt32(prime *big.Int) error {
	if prime == nil {
		return ErrNil
	}

	if prime.Cmp(int32Max) > 0 {
		return ErrOutOfRange
	}

	if !prime.ProbablyPrime(100) {
		return ErrNotAPrime
	}

	h.int32prime = prime
	inverse := *prime

	var plusOne big.Int
	h.int32inverse = inverse.ModInverse(&inverse, plusOne.Add(int32Max, bigOne))

	return nil
}

func (h *Hide) SetInt64(prime *big.Int) error {
	if prime == nil {
		return ErrNil
	}

	if prime.Cmp(int64Max) > 0 {
		return ErrOutOfRange
	}

	if !prime.ProbablyPrime(100) {
		return ErrNotAPrime
	}

	h.int64prime = prime
	inverse := *prime

	var plusOne big.Int
	h.int64inverse = inverse.ModInverse(&inverse, plusOne.Add(int64Max, bigOne))

	return nil
}

func (h *Hide) SetUint32(prime *big.Int) error {
	if prime == nil {
		return ErrNil
	}

	if prime.Cmp(uint32Max) > 0 {
		return ErrOutOfRange
	}

	if !prime.ProbablyPrime(100) {
		return ErrNotAPrime
	}

	h.uint32prime = prime
	inverse := *prime

	var plusOne big.Int
	h.uint32inverse = inverse.ModInverse(&inverse, plusOne.Add(uint32Max, bigOne))

	return nil
}

func (h *Hide) SetUint64(prime *big.Int) error {
	if prime == nil {
		return ErrNil
	}

	if prime.Cmp(uint64Max) > 0 {
		return ErrOutOfRange
	}

	if !prime.ProbablyPrime(100) {
		return ErrNotAPrime
	}

	h.uint64prime = prime
	inverse := *prime

	var plusOne big.Int
	h.uint64inverse = inverse.ModInverse(&inverse, plusOne.Add(uint64Max, bigOne))

	return nil
}

func (h *Hide) Int32Obfuscate(i int32) int32 {
	return Int32Obfuscate(i, h.int32prime)
}

func (h *Hide) Int32Deobfuscate(i int32) int32 {
	return Int32Deobfuscate(i, h.int32inverse)
}

func (h *Hide) Int64Obfuscate(i int64) int64 {
	return Int64Obfuscate(i, h.int64prime)
}

func (h *Hide) Int64Deobfuscate(i int64) int64 {
	return Int64Deobfuscate(i, h.int64inverse)
}

func (h *Hide) Uint32Obfuscate(i uint32) uint32 {
	return Uint32Obfuscate(i, h.uint32prime)
}

func (h *Hide) Uint32Deobfuscate(i uint32) uint32 {
	return Uint32Deobfuscate(i, h.uint32inverse)
}

func (h *Hide) Uint64Obfuscate(i uint64) uint64 {
	return Uint64Obfuscate(i, h.uint64prime)
}

func (h *Hide) Uint64Deobfuscate(i uint64) uint64 {
	return Uint64Deobfuscate(i, h.uint64inverse)
}
