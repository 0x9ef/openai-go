// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"context"
	"net/http"
)

type CompletionOptions struct {
	// ID of the model to use.
	Model Model `json:"model" binding:"required"`
	// Prompt to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
	Prompt []string `json:"prompt" binding:"required"`
	// The maximum number of tokens to generate in the completion.
	// The token count of your prompt plus max_tokens cannot exceed the model's context length.
	// Most models have a context length of 2048 tokens (except for the newest models, which support 4096).
	MaxTokens int `json:"max_tokens,omitempty" binding:"omitempty,max=4096"`
	// What sampling temperature to use. Higher values means the model will take more risks.
	// Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.
	Temperature int `json:"temperature,omitempty"`
	// How many completions to generate for each prompt.
	N int `json:"n,omitempty"`
	// Up to 4 sequences where the API will stop generating further tokens.
	// The returned text will not contain the stop sequence.
	Stop []string `json:"stop,omitempty"`
}

type CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   Model  `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Completion given a prompt, the model will return one or more predicted completions,
// and can also return the probabilities of alternative tokens at each position.
//
// The default number of tokens to complete is 1024.
// Docs: https://beta.openai.com/docs/api-reference/completions
func (e *Engine) Completion(ctx context.Context, opts *CompletionOptions) (*CompletionResponse, error) {
	if err := e.validate.StructCtx(ctx, opts); err != nil {
		return nil, err
	}
	uri := e.apiBaseURL + "/completions"
	if opts.MaxTokens == 0 {
		opts.MaxTokens = defaultMaxTokens
	}
	r, err := marshalJson(opts)
	if err != nil {
		return nil, err
	}
	req, err := e.newReq(ctx, http.MethodPost, uri, "json", r)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	// Umarshal response to CompletionResponse
	var jsonResp CompletionResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}
