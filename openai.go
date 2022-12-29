// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Engine struct {
	apiKey     string
	apiBaseURL string
	client     *http.Client
	validate   *validator.Validate
	n          int
}

func New(apiKey string) *Engine {
	e := &Engine{
		apiKey:     apiKey,
		apiBaseURL: "https://api.openai.com/v1",
		client:     &http.Client{},
		validate:   validator.New(),
	}
	v := validator.New()
	v.SetTagName("binding")
	e.validate = v
	return e
}

func (e *Engine) newReq(ctx context.Context, method string, url string, body any) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background() // prevent nil context error
	}
	r := new(bytes.Reader)
	if body != nil {
		jsonb, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(jsonb)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, r)
	if err != nil {
		return nil, err
	}
	// Setup Content-Type=application/json header only on POST operation
	if body != nil && method == http.MethodPost {
		req.Header.Set("Content-type", "application/json")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", e.apiKey))
	return req, err
}

func (e *Engine) doReq(req *http.Request) (*http.Response, error) {
	e.n++ // increment number of requests
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for valid status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp, nil
	}

	// If we have not-success HTTP status code, unmarshal to APIError
	var apiErr APIError
	if err := unmarshal(resp, &apiErr); err != nil {
		return nil, err
	}
	if apiErr.Err.StatusCode == 0 {
		// Overwrite apiErr status code if it's zero
		apiErr.Err.StatusCode = resp.StatusCode
	}
	return resp, apiErr
}

func unmarshal(resp *http.Response, v any) error {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
