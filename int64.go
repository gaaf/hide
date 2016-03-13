package hide

import (
	"encoding/json"
	"math/big"
)

type Int64 int64

func (i *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(Int64Obfuscate(int64(*i), nil))
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	var obf int64
	if err := json.Unmarshal(data, &obf); err != nil {
		*i = Int64(obf)
		return err
	}
	*i = Int64(Int64Deobfuscate(obf, nil))
	return nil
}

func Int64Obfuscate(val int64, prime *big.Int) int64 {
	if prime == nil {
		prime = Default.int64prime
	}
	bg := new(big.Int).SetInt64(val)
	modularMultiplicativeInverse(bg, prime, int64Max)

	return bg.Int64()
}

func Int64Deobfuscate(val int64, coprime *big.Int) int64 {
	if coprime == nil {
		coprime = Default.int64coprime
	}
	bg := new(big.Int).SetInt64(val)
	modularMultiplicativeInverse(bg, coprime, int64Max)

	return bg.Int64()
}
