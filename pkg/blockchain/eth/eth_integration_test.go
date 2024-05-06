package eth

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var eth *Eth

func TestMain(m *testing.M) {
	eth = NewEther("")
	os.Exit(m.Run())
}

func TestGetHeight(t *testing.T) {
	result, err := eth.GetHeight(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, result, int64(0))
}

func TestGetBlockTransactions(t *testing.T) {
	result, err := eth.GetTransactionsByBlock(context.Background(), int64(19735503))
	assert.NoError(t, err)
	assert.Greater(t, len(result.Transactions), 0)
}
