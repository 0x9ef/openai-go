// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"context"
	"net/http"
)

// Generative Pre-trained Transformer (GPT) model.
//
// Learn more: https://beta.openai.com/docs/models
type Model string

// The Codex models are descendants of our GPT-3 models that can understand and generate code.
// Their training data contains both natural language and billions of lines of public code from GitHub.
// They’re most capable in Python and proficient in over a dozen languages including
// JavaScript, Go, Perl, PHP, Ruby, Swift, TypeScript, SQL, and even Shell.
//
// Learn more: https://platform.openai.com/docs/models/codex
const (
	ModelCodexDavinci002 Model = "code-davinci-002"
	ModelCodexCushman001 Model = "code-cushman-001"
)

// GPT-3 models can understand and generate natural language.
// These models were superceded by the more powerful GPT-3.5 generation models.
// However, the original GPT-3 base models (davinci, curie, ada, and babbage) are current the only
// models that are available to fine-tune.
const (
	ModelGPT3Ada            Model = "ada"
	ModelGPT3Babbage        Model = "babbage"
	ModelGPT3TextBabbage    Model = "text-babbage-001"
	ModelGPT3Curie          Model = "curie"
	ModelGPT3TextCurie001   Model = "text-curie-001"
	ModelGPT3Davince        Model = "davinci"
	ModelGPT3TextDavince    Model = "text-davinci-001"
	ModelGPT3TextDavinci002 Model = "text-davinci-002"
	ModelGPT3TextDavinci003 Model = "text-davinci-003"
	ModelGPT3TextAda001     Model = "text-ada-001"
	// DefaultModel is alias to ModelGPT3TextDavinci003
	DefaultModel = ModelGPT3TextDavinci003
)

// GPT-3.5 models can understand and generate natural language or code.
// Our most capable and cost effective model in the GPT-3.5 family is gpt-3.5-turbo which has been
// optimized for chat but works well for traditional completions tasks as well.
//
// Learn more: https://platform.openai.com/docs/models/gpt-3-5
const (
	ModelGPT3Dot5Turbo0301 Model = "gpt-3.5-turbo-0301"
	ModelGPT3Dot5Turbo     Model = "gpt-3.5-turbo"
)

// GPT4 generation models.
// GPT-4 is a large multimodal model (accepting text inputs and emitting text outputs today, with image inputs
// coming in the future) that can solve difficult problems with greater accuracy than any of our previous models,
// thanks to its broader general knowledge and advanced reasoning capabilities.
// Like gpt-3.5-turbo, GPT-4 is optimized for chat but works well for traditional completions tasks.
//
// Learn more: https://platform.openai.com/docs/models/gpt-4
const (
	ModelGPT4        Model = "gpt-4"
	ModelGPT432K0314 Model = "gpt-4-32k-0314"
	ModelGPT432K     Model = "gpt-4-32k"
	ModelGPT40314    Model = "gpt-4-0314"
)

type ListModelsResponse struct {
	Data []struct {
		ID      Model  `json:"id"`
		Object  string `json:"object"`
		OwnedBy string `json:"owned_by"`
	} `json:"data"`
}

// ListModels lists the currently available models, and provides basic information about
// each one such as the owner and availability.
//
// Docs: https://beta.openai.com/docs/api-reference/models/list
func (e *Engine) ListModels(ctx context.Context) (*ListModelsResponse, error) {
	url := e.apiBaseURL + "/models"
	req, err := e.newReq(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp ListModelsResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

type RetrieveModelOptions struct {
	// The ID of the model.
	ID Model `json:"id" binding:"required"`
}

type RetrieveModelResponse struct {
	ID      Model  `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}

// RetrieveModel retrieves a model instance, providing basic information
// about the model such as the owner and permissioning.
//
// Docs: https://beta.openai.com/docs/api-reference/models/retrieve
func (e *Engine) RetrieveModel(ctx context.Context, opts *RetrieveModelOptions) (*RetrieveModelResponse, error) {
	if err := e.validate.StructCtx(ctx, opts); err != nil {
		return nil, err
	}
	url := e.apiBaseURL + "/models/" + string(opts.ID)
	req, err := e.newReq(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp RetrieveModelResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}
