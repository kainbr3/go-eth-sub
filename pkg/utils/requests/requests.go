package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	me "github.com/kainbr3/go-eth-sub/internal/mappedErrors"
)

func New(ctx context.Context, method, url string, target, payload any) error {
	var req *http.Request
	var body io.Reader

	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			log.Printf("error trying to marshal payload, with error: %s", err.Error())
			body = nil
			return me.ErrJsonMarshal.BuildErrorMessage(me.UTILS, err)
		}

		body = bytes.NewReader(data)
	}

	req, _ = http.NewRequestWithContext(ctx, method, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "trustWallet")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		if err != nil {
			return me.ErrRequest.BuildErrorMessage(me.UTILS, fmt.Errorf("error trying execute a http request with statuscode: %d - endpoint: %s - with error: %s", resp.StatusCode, url, err.Error()))
		} else {
			return me.ErrRequest.BuildErrorMessage(me.UTILS, fmt.Errorf("error trying execute a http request with statuscode: %d - endpoint: %s", resp.StatusCode, url))
		}
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&target)

	return nil
}
