package hide

import (
	"encoding/json"
	"math/big"
)

type Uint32 uint32

func (i *Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(Uint32Obfuscate(uint32(*i), nil))
}

func (i *Uint32) UnmarshalJSON(data []byte) error {
	var obf uint32
	if err := json.Unmarshal(data, &obf); err != nil {
		*i = Uint32(obf)
		return err
	}
	*i = Uint32(Uint32Deobfuscate(obf, nil))
	return nil
}

func Uint32Obfuscate(val uint32, prime *big.Int) uint32 {
	if prime == nil {
		prime = Default.uint32prime
	}
	bg := new(big.Int).SetUint64(uint64(val))
	modularMultiplicativeInverse(bg, prime, uint32Max)

	return uint32(bg.Uint64())
}

func Uint32Deobfuscate(val uint32, coprime *big.Int) uint32 {
	if coprime == nil {
		coprime = Default.uint32coprime
	}
	bg := new(big.Int).SetUint64(uint64(val))
	modularMultiplicativeInverse(bg, coprime, uint32Max)

	return uint32(bg.Uint64())
}
