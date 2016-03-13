package hide

import (
	"encoding/json"
	"math/big"
)

type Uint64 uint64

func (i *Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(Uint64Obfuscate(uint64(*i), nil))
}

func (i *Uint64) UnmarshalJSON(data []byte) error {
	var obf uint64
	if err := json.Unmarshal(data, &obf); err != nil {
		*i = Uint64(obf)
		return err
	}
	*i = Uint64(Uint64Deobfuscate(obf, nil))
	return nil
}

func Uint64Obfuscate(val uint64, prime *big.Int) uint64 {
	if prime == nil {
		prime = Default.uint64prime
	}
	bg := new(big.Int).SetUint64(val)
	modularMultiplicativeInverse(bg, prime, uint64Max)

	return bg.Uint64()
}

func Uint64Deobfuscate(val uint64, coprime *big.Int) uint64 {
	if coprime == nil {
		coprime = Default.uint64coprime
	}
	bg := new(big.Int).SetUint64(val)
	modularMultiplicativeInverse(bg, coprime, uint64Max)

	return bg.Uint64()
}
