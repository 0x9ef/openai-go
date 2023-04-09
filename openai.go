// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Engine struct {
	apiKey         string
	apiBaseURL     string
	organizationId string
	client         *http.Client
	validate       *validator.Validate
	n              int
}

const (
	defaultMaxTokens = 1024
)

// New is used to initialize engine.
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

// SetApiKey is used to set API key to access OpenAI API.
func (e *Engine) SetApiKey(apiKey string) {
	e.apiKey = apiKey
}

// SetOrganizationId is used to set organization ID if user belongs to multiple organizations.
func (e *Engine) SetOrganizationId(organizationId string) {
	e.organizationId = organizationId
}

func (e *Engine) newReq(ctx context.Context, method string, uri string, postType string, body io.Reader) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background() // prevent nil context error
	}
	if body == nil {
		body = new(bytes.Reader) // prevent nil body error
	}
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", e.apiKey))
	if len(e.organizationId) != 0 {
		req.Header.Set("OpenAI-Organization", e.organizationId)
	}
	// Setup Content-Type depends on postType
	switch {
	case body != nil && postType == "json":
		req.Header.Set("Content-type", "application/json")
	case body != nil && postType == "formData":
		req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	case body != nil:
		req.Header.Set("Content-type", postType)
	}
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

func unmarshal(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}

func marshalJson(body interface{}) (io.Reader, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
