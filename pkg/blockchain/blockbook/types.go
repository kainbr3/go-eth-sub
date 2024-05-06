package blockbook

import "time"

type StatusBlockbook struct {
	Coin                         string    `json:"coin"`
	Host                         string    `json:"host"`
	Version                      string    `json:"version"`
	GitCommit                    string    `json:"gitCommit"`
	BuildTime                    time.Time `json:"buildTime"`
	SyncMode                     bool      `json:"syncMode"`
	InitialSync                  bool      `json:"initialSync"`
	InSync                       bool      `json:"inSync"`
	BestHeight                   int       `json:"bestHeight"`
	LastBlockTime                time.Time `json:"lastBlockTime"`
	InSyncMempool                bool      `json:"inSyncMempool"`
	LastMempoolTime              time.Time `json:"lastMempoolTime"`
	MempoolSize                  int       `json:"mempoolSize"`
	Decimals                     int       `json:"decimals"`
	DbSize                       int64     `json:"dbSize"`
	HasFiatRates                 bool      `json:"hasFiatRates"`
	HasTokenFiatRates            bool      `json:"hasTokenFiatRates"`
	CurrentFiatRatesTime         time.Time `json:"currentFiatRatesTime"`
	HistoricalFiatRatesTime      time.Time `json:"historicalFiatRatesTime"`
	HistoricalTokenFiatRatesTime time.Time `json:"historicalTokenFiatRatesTime"`
	SupportedStakingPools        []string  `json:"supportedStakingPools"`
	About                        string    `json:"about"`
}

type StatusBackend struct {
	Chain            string `json:"chain"`
	Blocks           int    `json:"blocks"`
	BestBlockHash    string `json:"bestBlockHash"`
	Difficulty       string `json:"difficulty"`
	Version          string `json:"version"`
	ConsensusVersion string `json:"consensus_version"`
}

type BlockchainStatusResponse struct {
	Blockbook StatusBlockbook `json:"blockbook"`
	Backend   StatusBackend   `json:"backend"`
}

type Vin struct {
	N         int      `json:"n"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

type Vout struct {
	Value     string   `json:"value"`
	N         int      `json:"n"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

type Param struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type ParsedData struct {
	MethodID string  `json:"methodId"`
	Name     string  `json:"name"`
	Function string  `json:"function"`
	Params   []Param `json:"params"`
}

type InternalTransfers struct {
	Type  int    `json:"type"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type TokenTransfers struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Value    string `json:"value"`
}

type Specific struct {
	Status            int                 `json:"status"`
	Nonce             int                 `json:"nonce"`
	GasLimit          int                 `json:"gasLimit"`
	GasUsed           int                 `json:"gasUsed"`
	GasPrice          string              `json:"gasPrice"`
	Data              string              `json:"data"`
	ParsedData        ParsedData          `json:"parsedData"`
	InternalTransfers []InternalTransfers `json:"internalTransfers"`
}

type Transaction struct {
	Txid             string           `json:"txid"`
	Vin              []Vin            `json:"vin"`
	Vout             []Vout           `json:"vout"`
	BlockHash        string           `json:"blockHash"`
	BlockHeight      int              `json:"blockHeight"`
	Confirmations    int              `json:"confirmations"`
	BlockTime        int              `json:"blockTime"`
	Value            string           `json:"value"`
	Fees             string           `json:"fees"`
	TokenTransfers   []TokenTransfers `json:"tokenTransfers"`
	EthereumSpecific Specific         `json:"ethereumSpecific"`
}

type AddressAliases map[string]map[string]string

type TransactionsByBlockResponse struct {
	Page              int            `json:"page"`
	TotalPages        int            `json:"totalPages"`
	ItemsOnPage       int            `json:"itemsOnPage"`
	Hash              string         `json:"hash"`
	PreviousBlockHash string         `json:"previousBlockHash"`
	NextBlockHash     string         `json:"nextBlockHash"`
	Height            int            `json:"height"`
	Confirmations     int            `json:"confirmations"`
	Size              int            `json:"size"`
	Time              int            `json:"time"`
	Version           int            `json:"version"`
	MerkleRoot        string         `json:"merkleRoot"`
	Nonce             string         `json:"nonce"`
	Bits              string         `json:"bits"`
	Difficulty        string         `json:"difficulty"`
	TxCount           int            `json:"txCount"`
	Txs               []Transaction  `json:"txs"`
	AddressAliases    AddressAliases `json:"addressAliases"`
}

type MultiTokenValues struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Tokens struct {
	Type             string             `json:"type"`
	Name             string             `json:"name"`
	Contract         string             `json:"contract"`
	Transfers        int                `json:"transfers"`
	Symbol           string             `json:"symbol"`
	Decimals         int                `json:"decimals,omitempty"`
	Balance          string             `json:"balance,omitempty"`
	MultiTokenValues []MultiTokenValues `json:"multiTokenValues,omitempty"`
}

type TransactionsByAddressResponse struct {
	Page               int            `json:"page"`
	TotalPages         int            `json:"totalPages"`
	ItemsOnPage        int            `json:"itemsOnPage"`
	Address            string         `json:"address"`
	Balance            string         `json:"balance"`
	UnconfirmedBalance string         `json:"unconfirmedBalance"`
	UnconfirmedTxs     int            `json:"unconfirmedTxs"`
	Txs                int            `json:"txs"`
	NonTokenTxs        int            `json:"nonTokenTxs"`
	Transactions       []Transaction  `json:"transactions"`
	Nonce              string         `json:"nonce"`
	Tokens             []Tokens       `json:"tokens"`
	AddressAliases     AddressAliases `json:"addressAliases"`
}
