package api

import (
	"log"
	"net/http"

	b "github.com/kainbr3/go-eth-sub/pkg/blockchain/blockbook"
	e "github.com/kainbr3/go-eth-sub/pkg/blockchain/eth"
	s "github.com/kainbr3/go-eth-sub/pkg/storage"
)

func Start() {
	mux := http.NewServeMux()
	strg, _ := s.NewStorage()

	handler := &Handler{
		node:    e.NewEther(""),
		indexer: b.NewBlockbook(""),
		storage: strg,
	}

	mux.HandleFunc("GET /v1/eth/height", handler.getBlockHandler)
	mux.HandleFunc("GET /v1/eth/transactions", handler.getTransactionsHandler)
	mux.HandleFunc("GET /v1/subscriptions", handler.getSubsHandler)
	mux.HandleFunc("POST /v1/subscriptions", handler.postSubsHandler)
	mux.HandleFunc("DELETE /v1/subscriptions/{address}", handler.delSubsHandler)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
