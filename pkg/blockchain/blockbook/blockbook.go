package blockbook

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kainbr3/go-eth-sub/internal/types"
	"github.com/kainbr3/go-eth-sub/pkg/utils/converter"
	"github.com/kainbr3/go-eth-sub/pkg/utils/requests"
)

const (
	ETH_MAIN_TOKEN      = "ETH"
	ETH_MAIN_TOKEN_TYPE = "ETH"
	ETH_CHAIN_PRECISION = 18
)

type Blockbook struct {
	explorerApiUrl string
}

func NewBlockbook(blockbookUrl string) *Blockbook {
	blockbook := &Blockbook{explorerApiUrl: "https://eth1.trezor.io/api"}

	if blockbookUrl != "" {
		blockbook.explorerApiUrl = blockbookUrl
	}

	return blockbook
}

func (b *Blockbook) GetHeight(ctx context.Context) (int64, error) {
	endpoint := fmt.Sprintf("%s/v2", b.explorerApiUrl)
	response := &BlockchainStatusResponse{}

	err := requests.New(ctx, "POST", endpoint, response, nil)
	if err != nil {
		return 0, err
	}

	return int64(response.Blockbook.BestHeight), nil
}

func (b *Blockbook) GetTransactionsByBlock(ctx context.Context, block int64) ([]*types.Transaction, error) {
	endpoint := fmt.Sprintf("%s/v2/block/%d", b.explorerApiUrl, block)
	response := &TransactionsByBlockResponse{}

	err := requests.New(ctx, "POST", endpoint, response, nil)
	if err != nil {
		return nil, err
	}

	if len(response.Txs) == 0 {
		return nil, fmt.Errorf("any transaction for this block: %d", block)
	}

	result := parseTransactions(response.Txs)

	return result, nil
}

func (b *Blockbook) GetTransactionByHash(ctx context.Context, hash string) (*types.Transaction, error) {
	endpoint := fmt.Sprintf("%s/v2/tx/%s", b.explorerApiUrl, hash)
	response := Transaction{}

	err := requests.New(ctx, "POST", endpoint, &response, nil)
	if err != nil {
		return nil, err
	}

	result := parseTransaction(response)

	if result == nil {
		return nil, errors.New("transaction not found")
	}

	return result, nil
}

func (b *Blockbook) GetTransactionsByAddress(ctx context.Context, address string) ([]*types.Transaction, error) {
	endpoint := fmt.Sprintf("%s/v2/address/%s?details=txs", b.explorerApiUrl, address)
	response := TransactionsByAddressResponse{}

	err := requests.New(ctx, "POST", endpoint, &response, nil)
	if err != nil {
		return nil, err
	}

	if len(response.Transactions) == 0 {
		return nil, fmt.Errorf("any transaction for this address: %s", address)
	}

	result := parseTransactions(response.Transactions)

	if len(result) == 0 {
		return nil, fmt.Errorf("any transaction for this address: %s", address)
	}

	return result, nil
}

func parseTransaction(tx Transaction) *types.Transaction {
	txType := tx.EthereumSpecific.ParsedData.Name
	if !strings.EqualFold(txType, "Transfer") {
		return nil
	}

	from := ""
	to := ""

	if len(tx.Vin) > 0 {
		if len(tx.Vin[0].Addresses) > 0 {
			from = tx.Vin[0].Addresses[0]
		}
	}

	if len(tx.Vout) > 0 {
		if len(tx.Vout[0].Addresses) > 0 {
			to = tx.Vout[0].Addresses[0]
		}
	}

	gasPrice := tx.EthereumSpecific.GasPrice
	gasPriceEthFloat := converter.StringValueIntoFloatWithPrecision(gasPrice, ETH_CHAIN_PRECISION)
	gasPriceGweiFloat := gasPriceEthFloat * 1000000000

	parsedGasPriceEth := strconv.FormatFloat(gasPriceEthFloat, 'f', -1, 64)
	parsedGasPriceGwei := strconv.FormatFloat(gasPriceGweiFloat, 'f', -1, 64)

	timestampUnix := int64(tx.BlockTime)
	timestampTime := time.Unix(timestampUnix, 0)
	timestap := timestampTime.Format("02-01-2006 15:04:05")

	result := &types.Transaction{
		Hash:          tx.Txid,
		Block:         int64(tx.BlockHeight),
		BlockHash:     tx.BlockHash,
		From:          from,
		To:            to,
		Asset:         ETH_MAIN_TOKEN,
		AssetType:     ETH_MAIN_TOKEN_TYPE,
		Contract:      "",
		Amount:        converter.StringValueIntoRawWithPrecision(tx.Value, ETH_CHAIN_PRECISION),
		Fee:           converter.StringValueIntoRawWithPrecision(tx.Fees, ETH_CHAIN_PRECISION),
		Confirmations: tx.Confirmations,
		Timestamp:     timestap,
		Nonce:         tx.EthereumSpecific.Nonce,
		Status:        parseStatus(tx.EthereumSpecific.Status),
		Type:          txType,
		GasLimit:      int64(tx.EthereumSpecific.GasLimit),
		GasUsed:       int64(tx.EthereumSpecific.GasUsed),
		GasPriceEth:   parsedGasPriceEth,
		GasPriceGwei:  parsedGasPriceGwei,
	}

	return result
}

func parseTokenTransferTransaction(tx Transaction) []*types.Transaction {
	result := []*types.Transaction{}

	for _, token := range tx.TokenTransfers {
		if !strings.EqualFold(tx.EthereumSpecific.ParsedData.Name, "Transfer") {
			continue
		}

		gasPrice := tx.EthereumSpecific.GasPrice
		gasPriceEthFloat := converter.StringValueIntoFloatWithPrecision(gasPrice, ETH_CHAIN_PRECISION)
		gasPriceGweiFloat := gasPriceEthFloat * 1000000000

		parsedGasPriceEth := strconv.FormatFloat(gasPriceEthFloat, 'f', -1, 64)
		parsedGasPriceGwei := strconv.FormatFloat(gasPriceGweiFloat, 'f', -1, 64)

		timestampUnix := int64(tx.BlockTime)
		timestampTime := time.Unix(timestampUnix, 0)
		timestap := timestampTime.Format("02-01-2006 15:04:05")

		parsedTx := &types.Transaction{
			Hash:          tx.Txid,
			Block:         int64(tx.BlockHeight),
			BlockHash:     tx.BlockHash,
			From:          token.From,
			To:            token.To,
			Asset:         token.Symbol,
			AssetType:     token.Type,
			Contract:      token.Contract,
			Amount:        converter.StringValueIntoRawWithPrecision(token.Value, token.Decimals),
			Fee:           converter.StringValueIntoRawWithPrecision(tx.Fees, ETH_CHAIN_PRECISION),
			Confirmations: tx.Confirmations,
			Timestamp:     timestap,
			Nonce:         tx.EthereumSpecific.Nonce,
			Status:        parseStatus(tx.EthereumSpecific.Status),
			Type:          "Token Transfer",
			GasLimit:      int64(tx.EthereumSpecific.GasLimit),
			GasUsed:       int64(tx.EthereumSpecific.GasUsed),
			GasPriceEth:   parsedGasPriceEth,
			GasPriceGwei:  parsedGasPriceGwei,
		}

		result = append(result, parsedTx)
	}

	return result
}

func parseTransactions(txs []Transaction) []*types.Transaction {
	result := []*types.Transaction{}

	for _, tx := range txs {
		if len(tx.TokenTransfers) > 0 {
			result = append(result, parseTokenTransferTransaction(tx)...)
		} else {
			parsedTx := parseTransaction(tx)

			if parsedTx != nil {
				result = append(result, parseTransaction(tx))
			}
		}
	}

	return result
}

func parseStatus(statusCode int) string {
	status := ""

	switch statusCode {
	case 0:
		status = "failed"
	case 1:
		status = "confirmed"
	case -1:
		status = "pending"
	}

	return status
}
