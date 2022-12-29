// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"context"
	"net/http"
)

type EditOptions struct {
	// ID of the model to use.
	Model Model `json:"model" binding:"required"`
	// The input text to use as a starting point for the edit.
	Input string `json:"input" binding:"required"`
	// The instruction that tells the model how to edit the prompt.
	Instruction string `json:"instruction" binding:"required"`
	// How many edits to generate for the input and instruction.
	// Defaults to 1.
	N int `json:"n,omitempty"`
	// What sampling temperature to use. Higher values means the model will take more risks.
	// Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.
	Temperature int `json:"temperature,omitempty"`
}

type EditResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
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

// Edit given a prompt and an instruction, the model will return an edited version of the prompt.
//
// Docs: https://beta.openai.com/docs/api-reference/edits
func (e *Engine) Edit(ctx context.Context, opts *EditOptions) (*EditResponse, error) {
	if err := e.validate.StructCtx(ctx, opts); err != nil {
		return nil, err
	}
	url := e.apiBaseURL + "/edits"
	req, err := e.newReq(ctx, http.MethodPost, url, opts)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp EditResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}
