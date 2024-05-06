package types

type AddressList struct {
	Address []string `json:"addresses"`
}

type Transaction struct {
	Hash          string `json:"hash"`
	Block         int64  `json:"block"`
	BlockHash     string `json:"block_hash"`
	From          string `json:"from"`
	To            string `json:"to"`
	Asset         string `json:"asset"`
	AssetType     string `json:"asset_type"`
	Contract      string `json:"contract"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	Confirmations int    `json:"confirmations"`
	Timestamp     string `json:"timestamp"`
	Nonce         int    `json:"nonce"`
	Status        string `json:"status"`
	Type          string `json:"Type"`
	GasLimit      int64  `json:"gas_limit"`
	GasUsed       int64  `json:"gas_used"`
	GasPriceGwei  string `json:"gas_price_gwei"`
	GasPriceEth   string `json:"gas_price_eth"`
}

type SubscribeStatus struct {
	Subscribe bool `json:"subscribe"`
}

type CurrentHeight struct {
	Block int `json:"block"`
}
