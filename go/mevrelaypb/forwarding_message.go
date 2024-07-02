package mevrelaypb

import (
	"math/big"
)

// GetPayBidTxGasUsedBigInt returns the PayBidTxGasUsed as a big.Int
func (x *ForwardRequest) GetPayBidTxGasUsedBigInt() *big.Int {
	return new(big.Int).SetUint64(x.PayBidTxGasUsed)
}
