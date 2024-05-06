package blockbook

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var blockbook *Blockbook

func TestMain(m *testing.M) {
	blockbook = NewBlockbook("")
	os.Exit(m.Run())
}

func TestGetHeight(t *testing.T) {
	result, err := blockbook.GetHeight(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, result, int64(0))
}

func TestGetTransactionsByBlock(t *testing.T) {
	result, err := blockbook.GetTransactionsByBlock(context.Background(), 19773030)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, len(result), 0)
}

func TestGetTransactionByHash(t *testing.T) {
	txHash := "0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad"
	result, err := blockbook.GetTransactionByHash(context.Background(), txHash)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Hash, txHash)
}

func TestGetTransactionsByAddress(t *testing.T) {
	result, err := blockbook.GetTransactionsByAddress(context.Background(), "0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, len(result), 0)
}
