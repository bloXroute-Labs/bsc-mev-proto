package mevrelaypb

import (
	"math/big"
)

// NewBigInt creates a new BigInt from a big.Int
func NewBigInt(i *big.Int) *BigInt { return &BigInt{Value: i.Bytes()} }

// Int returns the big.Int representation of the BigInt
func (b *BigInt) Int() *big.Int { return new(big.Int).SetBytes(b.Value) }
