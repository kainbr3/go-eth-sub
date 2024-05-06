package storage

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	me "github.com/kainbr3/go-eth-sub/internal/mappedErrors"
	"github.com/kainbr3/go-eth-sub/internal/types"
	"github.com/stretchr/testify/assert"
)

var memoryStorage *MemoryStorage

func TestMain(m *testing.M) {
	memoryStorage, _ = NewStorage()

	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	key := "get-test/tx/19773030/0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad/0x9C74B1b6792289690c3fE302b2eA4392c52a1E29/0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC"
	data := `{"hash":"0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad","block":19773030,"block_hash":"0x22f5d145956b2a8780b93eaf4152ec93e1533f6d774de243103b9068f23199b7","from":"0x9C74B1b6792289690c3fE302b2eA4392c52a1E29","to":"0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC","asset":"ETH","asset_type":"ETH","contract":"","amount":"0.012162","fee":"0.000169","confirmations":1241,"timestamp":"01-05-2024 01:55:47","nonce":1,"status":"confirmed","Type":"Transfer","gas_limit":21000,"gas_used":21000,"gas_price_gwei":"8.028027916000001","gas_price_eth":"0.000000008028027916"}`
	err := memoryStorage.Set(context.Background(), key, data, time.Hour*2)
	assert.NoError(t, err)

	result, err := memoryStorage.Get(context.Background(), key)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	transaction := &types.Transaction{}
	err = json.Unmarshal([]byte(result.(string)), transaction)
	assert.NoError(t, err)
	assert.Equal(t, transaction.Hash, "0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad")
}

func TestSet(t *testing.T) {
	key := "set-test/tx/19773030/0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad/0x9C74B1b6792289690c3fE302b2eA4392c52a1E29/0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC"
	data := `{"hash":"0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad","block":19773030,"block_hash":"0x22f5d145956b2a8780b93eaf4152ec93e1533f6d774de243103b9068f23199b7","from":"0x9C74B1b6792289690c3fE302b2eA4392c52a1E29","to":"0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC","asset":"ETH","asset_type":"ETH","contract":"","amount":"0.012162","fee":"0.000169","confirmations":1241,"timestamp":"01-05-2024 01:55:47","nonce":1,"status":"confirmed","Type":"Transfer","gas_limit":21000,"gas_used":21000,"gas_price_gwei":"8.028027916000001","gas_price_eth":"0.000000008028027916"}`
	err := memoryStorage.Set(context.Background(), key, data, time.Hour*2)
	assert.NoError(t, err)

	err = memoryStorage.Set(context.Background(), key, data, time.Hour*2)
	assert.NotNil(t, err)
	assert.Equal(t, err, me.ErrStorageSetDuplicate)

	err = memoryStorage.Delete(context.Background(), key)
	assert.NoError(t, err)
}

func TestDel(t *testing.T) {
	key := "del-test/tx/19773030/0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad/0x9C74B1b6792289690c3fE302b2eA4392c52a1E29/0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC"
	data := `{"hash":"0xda43eb02844964a2269948da8716b3f36cd85cccd6bdca68c65497dbfda29aad","block":19773030,"block_hash":"0x22f5d145956b2a8780b93eaf4152ec93e1533f6d774de243103b9068f23199b7","from":"0x9C74B1b6792289690c3fE302b2eA4392c52a1E29","to":"0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC","asset":"ETH","asset_type":"ETH","contract":"","amount":"0.012162","fee":"0.000169","confirmations":1241,"timestamp":"01-05-2024 01:55:47","nonce":1,"status":"confirmed","Type":"Transfer","gas_limit":21000,"gas_used":21000,"gas_price_gwei":"8.028027916000001","gas_price_eth":"0.000000008028027916"}`
	err := memoryStorage.Set(context.Background(), key, data, time.Hour*2)
	assert.NoError(t, err)

	err = memoryStorage.Delete(context.Background(), key)
	assert.NoError(t, err)
}
