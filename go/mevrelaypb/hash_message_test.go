package mevrelaypb

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestNewHash(t *testing.T) {
	hash := common.HexToHash("0x1234")
	require.Equal(t, hash, NewHash(hash).ToCommon())
}
