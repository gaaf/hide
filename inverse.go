package hide

import (
	"errors"
	"math/big"
)

// Hide stores primes and coprimes used to obfuscate/deobfuscate different
// integer types
type Hide struct {
	int32prime    *big.Int
	int32coprime  *big.Int
	int64prime    *big.Int
	int64coprime  *big.Int
	uint32prime   *big.Int
	uint32coprime *big.Int
	uint64prime   *big.Int
	uint64coprime *big.Int
}

var (
	// Default stores default, fallback primes and coprimes
	Default = Hide{
		int32prime:    new(big.Int).SetInt64(1500450271),
		int32coprime:  new(big.Int).SetInt64(1482223135),
		int64prime:    new(big.Int).SetInt64(8230452606740808761),
		int64coprime:  new(big.Int).SetInt64(5754553063220537865),
		uint32prime:   new(big.Int).SetUint64(3877529093),
		uint32coprime: new(big.Int).SetUint64(3267000013),
		uint64prime:   new(big.Int).SetUint64(12764787846358441471),
		uint64coprime: new(big.Int).SetUint64(1510277086161461759),
	}

	bigOne = big.NewInt(1)

	// 1 more than maximum value for each type
	int32Max  = big.NewInt(2147483647)                       // 2^31-1
	int64Max  = big.NewInt(9223372036854775807)              // 2^63 -1
	uint32Max = new(big.Int).SetUint64(4294967295)           // 2^32 -1
	uint64Max = new(big.Int).SetUint64(18446744073709551615) // 2^64 -1

	ErrNil        = errors.New("prime is nil")
	ErrOutOfRange = errors.New("prime is greater than max value for the type")
	ErrNotAPrime  = errors.New("it is not a prime number")
)

func modularMultiplicativeInverse(val, prime, max *big.Int) {
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
	coprime := *prime
	h.int32coprime = coprime.ModInverse(&coprime, int32Max)

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
	coprime := *prime
	h.int64coprime = coprime.ModInverse(&coprime, int64Max)

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
	coprime := *prime
	h.uint32coprime = coprime.ModInverse(&coprime, uint32Max)

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
	coprime := *prime
	h.uint64coprime = coprime.ModInverse(&coprime, uint64Max)

	return nil
}
