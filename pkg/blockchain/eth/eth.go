package eth

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kainbr3/go-eth-sub/pkg/utils/requests"
)

type Eth struct {
	nodeApiUrl string
}

func NewEther(nodeUrl string) *Eth {
	eth := &Eth{nodeApiUrl: "https://cloudflare-eth.com"}

	if nodeUrl != "" {
		eth.nodeApiUrl = nodeUrl
	}

	return eth
}

func (e *Eth) GetHeight(ctx context.Context) (int64, error) {
	payload := &RPCRequest{
		Jsonrpc: "2.0",
		ID:      31,
		Method:  "eth_blockNumber",
	}

	response := &RPCResponse{}
	err := requests.New(ctx, "POST", e.nodeApiUrl, response, payload)
	if err != nil {
		return 0, err
	}

	lastBlockHex := response.Result.(string)
	result, err := strconv.ParseInt(lastBlockHex, 0, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (e *Eth) GetTransactionsByBlock(ctx context.Context, block int64) (*BlockTransactionsResult, error) {
	blockDecimal := strconv.FormatInt(block, 16)
	blockHex := fmt.Sprintf("0x%s", blockDecimal)

	payload := &RPCRequest{
		Jsonrpc: "2.0",
		ID:      31,
		Method:  "eth_getBlockByNumber",
		Params: []any{
			blockHex,
			true,
		},
	}

	response := &RPCTransactionsResponse{}
	err := requests.New(ctx, "POST", e.nodeApiUrl, response, payload)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}
