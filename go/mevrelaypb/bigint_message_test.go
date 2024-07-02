package mevrelaypb

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBigInt(t *testing.T) {
	bigInt := big.NewInt(1000)
	randInt, err := rand.Int(rand.Reader, bigInt)
	require.NoError(t, err)

	randIntProto := NewBigInt(randInt)
	require.Equal(t, randInt, randIntProto.Int())
}
