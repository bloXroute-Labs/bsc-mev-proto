package mevrelaypb

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetTxsHexBytes returns the Txs as a slice of hexutil.Bytes
func (x *RawBid) GetTxsHexBytes() []hexutil.Bytes {
	txs := make([]hexutil.Bytes, 0, len(x.Txs))
	for _, tx := range x.Txs {
		txs = append(txs, tx)
	}
	return txs
}
