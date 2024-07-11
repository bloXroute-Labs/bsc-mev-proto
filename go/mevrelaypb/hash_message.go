package mevrelaypb

import (
	"github.com/ethereum/go-ethereum/common"
)

// NewHash creates a new Hash from a common.Hash
func NewHash(i common.Hash) *Hash { return &Hash{Value: i.Bytes()} }

// ToCommon returns the common.Hash representation of the Hash
func (b *Hash) ToCommon() common.Hash { return common.BytesToHash(b.Value) }

// NewHashes creates a new Hashes from a []common.Hash
func NewHashes(i []common.Hash) *Hashes {
	hashes := new(Hashes)
	for _, hash := range i {
		hashes.Value = append(hashes.Value, NewHash(hash))
	}

	return hashes
}

// ToCommon returns the []common.Hash representation of the Hashes
func (b *Hashes) ToCommon() []common.Hash {
	hashes := make([]common.Hash, 0, len(b.Value))
	for _, hash := range b.Value {
		hashes = append(hashes, hash.ToCommon())
	}
	return hashes
}
