package hide

import (
	"encoding/json"
	"math/big"
)

type Int32 int32

func (i *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(Int32Obfuscate(int32(*i), nil))
}

func (i *Int32) UnmarshalJSON(data []byte) error {
	var obf int32
	if err := json.Unmarshal(data, &obf); err != nil {
		*i = Int32(obf)
		return err
	}
	*i = Int32(Int32Deobfuscate(obf, nil))
	return nil
}

func Int32Obfuscate(val int32, prime *big.Int) int32 {
	if prime == nil {
		prime = Default.int32prime
	}
	bg := new(big.Int).SetInt64(int64(val))
	modularMultiplicativeInverse(bg, prime, int32Max)
	return int32(bg.Int64())
}

func Int32Deobfuscate(val int32, coprime *big.Int) int32 {
	if coprime == nil {
		coprime = Default.int32coprime
	}
	bg := new(big.Int).SetInt64(int64(val))
	modularMultiplicativeInverse(bg, coprime, int32Max)

	return int32(bg.Int64())
}
