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
// Docs: https://beta.openai.com/docs/models
type Model string

const (
	ModelAda            Model = "ada"
	ModelTextAda001     Model = "text-ada-001"
	ModelBabbage        Model = "babbage"
	ModelTextBabbage001 Model = "text-babbage-001"
	ModelCurie          Model = "curie"
	ModelTextCurie001   Model = "text-curie-001"
	ModelDavinci        Model = "davinci"
	ModelTextDavinci001 Model = "text-davinci-001"
	ModelTextDavinci002 Model = "text-davinci-002"
	ModelTextDavinci003 Model = "text-davinci-003"
	DefaultModel              = ModelDavinci
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
