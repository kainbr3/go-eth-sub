package types

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	me "github.com/kainbr3/go-eth-sub/internal/mappedErrors"
)

type SubscribeRequest struct {
	Address string `json:"address" validate:"required" example:"0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC"`
}

func (s *SubscribeRequest) FromBody(r *http.Request) *me.MappedError {
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return me.ErrParseBodyParam.BuildErrorMessage(me.API, err)
	}

	return nil
}

func (s *SubscribeRequest) FromUrlPath(r *http.Request) *me.MappedError {
	address := strings.Split(r.URL.Path, "/subscriptions/")[1]

	if address == "" {
		return me.ErrParseURLPathParam
	}

	s.Address = address
	return nil
}

func (s *SubscribeRequest) IsValid() *me.MappedError {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return me.ErrStructValidation.BuildErrorMessage(me.API, err)
	}

	if !isValidEthereumAddress(s.Address) {
		return me.ErrInvalidEthereumAddress.BuildErrorMessage(me.API, err)
	}

	return nil
}

type TransactionsRequest struct {
	Address string `json:"address" validate:"required" example:"0xbd0fCcdC19bC3b979e8E256b7B88AAe7C77A5BEC"`
}

func (t *TransactionsRequest) FromQuerystring(r *http.Request) *me.MappedError {
	err := r.ParseForm()
	if err != nil {
		return me.ErrParseQueryParam.BuildErrorMessage(me.API, err)
	}

	t.Address = r.Form.Get("address")

	return nil
}

func (t *TransactionsRequest) IsValid() *me.MappedError {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return me.ErrStructValidation.BuildErrorMessage(me.API, err)
	}

	if !isValidEthereumAddress(t.Address) {
		return me.ErrInvalidEthereumAddress.BuildErrorMessage(me.API, err)
	}

	return nil
}

func isValidEthereumAddress(address string) bool {
	// Ethereum addresses are 20 bytes long, or 40 characters in hex.
	// They always start with '0x'.
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
