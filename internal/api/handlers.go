package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	me "github.com/kainbr3/go-eth-sub/internal/mappedErrors"
	"github.com/kainbr3/go-eth-sub/internal/types"
	b "github.com/kainbr3/go-eth-sub/pkg/blockchain/blockbook"
	e "github.com/kainbr3/go-eth-sub/pkg/blockchain/eth"
	s "github.com/kainbr3/go-eth-sub/pkg/storage"
)

type Handler struct {
	node    *e.Eth
	indexer *b.Blockbook
	storage *s.MemoryStorage
}

func (h *Handler) getBlockHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{}
	httpCode := http.StatusOK

	result, err := h.node.GetHeight(r.Context())
	if err != nil {
		httpCode = getErrorCode(err)
		response["error"] = result
	} else {
		response["data"] = result
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		payload     = &types.TransactionsRequest{}
		response    = map[string]any{}
		httpCode    = http.StatusOK
		errResponse error
	)

	if err := payload.FromQuerystring(r); err != nil {
		httpCode = getErrorCode(err)
		errResponse = err
	} else {
		if err := payload.IsValid(); err != nil {
			httpCode = getErrorCode(err)
			errResponse = err
		}
	}

	if errResponse == nil {
		if result, err := h.indexer.GetTransactionsByAddress(r.Context(), payload.Address); err != nil {
			httpCode = getErrorCode(err)
			response["error"] = err.Error()
		} else {
			response["data"] = result
		}
	} else {
		response["error"] = errResponse.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getSubsHandler(w http.ResponseWriter, r *http.Request) {
	responseData := &types.AddressList{}
	response := map[string]any{}
	httpCode := http.StatusOK

	result, err := h.storage.Get(r.Context(), "")
	if err != nil {
		httpCode = getErrorCode(err)
		response["error"] = result
	} else {
		responseData.Address = result.([]string)
		response["data"] = responseData
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) postSubsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		payload     = &types.SubscribeRequest{}
		response    = map[string]any{}
		httpCode    = http.StatusOK
		errResponse error
	)

	if err := payload.FromBody(r); err != nil {
		httpCode = getErrorCode(err)
		errResponse = err.BuildErrorMessage(me.API, nil)
	} else {
		if err := payload.IsValid(); err != nil {
			httpCode = getErrorCode(err)
			errResponse = err
		}
	}

	if errResponse == nil {
		key := fmt.Sprintf("subscription/%s", payload.Address)
		if err := h.storage.Set(r.Context(), key, nil, time.Hour*24); err != nil {
			if err == me.ErrStorageSetDuplicate {
				response["data"] = false
				response["error"] = err.Error()
			} else {
				httpCode = http.StatusInternalServerError
				response["error"] = me.ErrStorageSet.BuildErrorMessage(me.API, nil).Error()
			}
		} else {
			response["data"] = true
		}
	} else {
		response["error"] = errResponse.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) delSubsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		payload     = types.SubscribeRequest{}
		response    = map[string]any{}
		httpCode    = http.StatusOK
		errResponse error
	)

	if err := payload.FromUrlPath(r); err != nil {
		httpCode = getErrorCode(err)
		errResponse = err
	} else {
		if err := payload.IsValid(); err != nil {
			httpCode = getErrorCode(err)
			errResponse = err
		}
	}

	if errResponse == nil {
		key := fmt.Sprintf("subscription/%s", payload.Address)
		if err := h.storage.Delete(r.Context(), key); err != nil {
			httpCode = getErrorCode(err)
			response["error"] = err.Error()
		} else {
			response["data"] = "deleted successfully"
		}
	} else {
		response["error"] = errResponse.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func getErrorCode(err error) (httpCode int) {
	switch err {
	case me.ErrUnexpected:
		httpCode = http.StatusInternalServerError

	case me.ErrNotFound:
		httpCode = http.StatusNotFound

	default:
		httpCode = http.StatusBadRequest
	}

	return httpCode
}
