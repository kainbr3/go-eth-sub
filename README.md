# Project: Go Ethereum Subscriptions

More details [here](/doc/README.md)

### Preamble
*This solution offers a RESTful API that utilizes a public node and a public indexer to provide blockchain information and manage a list of subscriptions.*

*As per your request, this solution does not utilize any external packages other than the built-in Golang packages. Note that popular frameworks often provide built-in error handling and some abstraction. To replicate this functionality using only standard packages, the code may be a bit more verbose*

### The Architecture
*The solution is structured following the standard Golang directory layout. The doc directory contains documentation assets, the internal directory contains private solution packages, and the pkg directory contains general packages that can be reused in other solutions.*

#### Internal:
📂internal
┣ 📂api
┣ 📂mappedErrors
┗ 📂types

* The folder structure here is self-explanatory.

### PKG
 📂pkg
 ┣ 📂blockchain
 ┃ ┣ 📂blockbook
 ┃ ┗ 📂eth
 ┣ 📂storage
 ┗ 📂utils
   ┣ 📂converter
   ┗ 📂requests

* blockchain/eth: Client to retrieve blockchain information from the eth node.

* blockchain/blockbook: Client for the public Ethereum indexer to retrieve all transactions from a given address. Since there is no indexation on the node, the only way to retrieve transactions for an address is through the GetLogs Method, which has a small default limit.

* storage: Storage manager to save subscribed addresses. It utilizes an interface and an in-memory implementation but can be easily ported to use Redis, Memcached, or other shared memory solutions.

* utils/converter: Parses values into human-readable format. Blockchain stores values as big integers strings, and each token has its own precision (e.g., ETH with 18 decimals and USDT ERC20 with 6).

* utils/requests: HTTP request handler abstraction.


Since there is no specification for the transaction model, I am using this struct as a Data Transfer Object (DTO) to represent all relevant information for a transaction.

``` 
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
```

PS: If this solution were to be made production-ready, it would consist of three parts:

A RESTful API to accept subscriptions via a websocket.
A service to pull blockchain transactions, storing them on a Kafka topic.
A consumer to read the topics, check transactions for subscribed addresses, and call the notification route to push notifications to users. Additionally, it could offer features such as real-time balance updates.

## API Documentation 
##### PS: Can be imported to insomnia/post using the file API Collection.json

![](/doc/images/api/routes.png "api")

URL: (GET - no parameters required)
http://localhost:8080/v1/eth/height

Response: 

![alt text](/doc/images/api/height.png)

URL: (GET - query param: address )
http://localhost:8080/v1/eth/transactions

ex.: http://localhost:8080/v1/eth/transactions?address=0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC

Response: 
![alt text](/doc/images/api/transactions.png)


URL: (GET - no parameters required)
http://localhost:8080/v1/subscriptions


Response: 
![alt text](/doc/images/api/get-subs.png)

URL: (POST - body { "address": "xxxxx"})
http://localhost:8080/v1/subscriptions

ex.: 1
```
{
	"address":"0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC"
}
```

ex.: 2
```
{
	"address":"0x7D34bF994459c831d73A98bFC55B01ABA5768A34"
}
```

Response: 
![alt text](/doc/images/api/post-subs.png)


URL: (DEL - url path param: address)
http://localhost:8080/v1/subscriptions/{address}

ex.: http://localhost:8080/v1/subscriptions/0x7D34bF994459c831d73A98bFC55B01ABA5768A34


Response: 
![alt text](/doc/images/api/del-subs.png)
