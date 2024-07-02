package mevrelaypb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetParentHashCommon returns the ParentHash as a common.Hash
func (x *RawBid) GetParentHashCommon() common.Hash {
	return common.HexToHash(x.ParentHash)
}

// GetUnRevertibleCommon returns the UnRevertible as a slice of common.Hash
func (x *RawBid) GetUnRevertibleCommon() []common.Hash {
	hashes := make([]common.Hash, 0, len(x.UnRevertible))
	for _, hash := range x.UnRevertible {
		hashes = append(hashes, common.BytesToHash(hash))
	}
	return hashes
}

// GetTxsHexBytes returns the Txs as a slice of hexutil.Bytes
func (x *RawBid) GetTxsHexBytes() []hexutil.Bytes {
	txs := make([]hexutil.Bytes, 0, len(x.Txs))
	for _, tx := range x.Txs {
		txs = append(txs, tx)
	}
	return txs
}
