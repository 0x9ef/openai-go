// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Size represents X*Y size wide of image.
type Size string

const (
	Size256    = "256x256"
	Size512    = "512x512"
	Size1024   = "1024x1024"
	SizeSmall  = Size256
	SizeMedium = Size512
	SizeBig    = Size1024
)

// ResponseFormat represents image format of response.
// It can be encoded as URL, or with base64+json.
type ResponseFormat string

const (
	ResponseFormatUrl     = "url"
	ResponseFormatB64Json = "b64_json"
)

type ImageCreateOptions struct {
	Prompt string `json:"prompt" binding:"required"`
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `json:"n,omitempty" binding:"omitempty,min=1,max=10"`
	// The size of the generated images.
	// Must be one of 256x256, 512x512, or 1024x1024.
	Size string `json:"size,omitempty" binding:"oneof=256x256 512x512 1024x1024"`
	// The format in which the generated images are returned.
	// Must be one of url or b64_json
	ResponseFormat string `json:"response_format,omitempty" binding:"omitempty,oneof=url b64_json"`
}

type ImageCreateResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// ImageCreate given a prompt and/or an input image, the model will generate a new image.
//
// Docs: https://beta.openai.com/docs/api-reference/images/create
func (e *Engine) ImageCreate(ctx context.Context, opts *ImageCreateOptions) (*ImageCreateResponse, error) {
	if err := e.validate.StructCtx(ctx, opts); err != nil {
		return nil, err
	}
	url := e.apiBaseURL + "/images/generations"
	if len(opts.Size) == 0 {
		opts.Size = SizeSmall
	}
	if len(opts.ResponseFormat) == 0 {
		opts.ResponseFormat = ResponseFormatUrl
	}
	r, err := marshalJson(opts)
	if err != nil {
		return nil, err
	}
	req, err := e.newReq(ctx, http.MethodPost, url, "json", r)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp ImageCreateResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

type ImageEditOptions struct {
	// The image to edit. Must be a valid PNG file, less than 4MB, and square.
	// If mask is not provided, image must have transparency, which will be used as the mask.
	Image string `binding:"required"`
	// An additional image whose fully transparent areas (e.g. where alpha is zero)
	// indicate where image should be edited. Must be a valid PNG file, less than 4MB,
	// and have the same dimensions as image.
	Mask string `binding:"omitempty"`
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `binding:"required,max=1000"`
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `binding:"min=1,max=10"`
	// The size of the generated images.
	// Must be one of 256x256, 512x512, or 1024x1024.
	Size string `binding:"omitempty,oneof=256x256 512x512 1024x1024"`
	// The format in which the generated images are returned.
	// Must be one of url or b64_json
	ResponseFormat string `binding:"omitempty,oneof=url b64_json"`
}

type ImageEditResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// ImageEdit creates an edited or extended image given an original image and a prompt.
//
// Docs: https://beta.openai.com/docs/api-reference/images/create-edit
func (e *Engine) ImageEdit(ctx context.Context, opts *ImageEditOptions) (*ImageEditResponse, error) {
	if err := e.validate.StructCtx(ctx, opts); err != nil {
		return nil, err
	}
	uri := e.apiBaseURL + "/images/edits"
	if opts.N == 0 {
		opts.N = 1
	}
	if len(opts.Size) == 0 {
		opts.Size = SizeSmall
	}
	if len(opts.ResponseFormat) == 0 {
		opts.ResponseFormat = ResponseFormatUrl
	}
	postValues := url.Values{
		"image":           []string{opts.Image},
		"mask":            []string{opts.Mask},
		"prompt":          []string{opts.Prompt},
		"n":               []string{strconv.Itoa(opts.N)},
		"size":            []string{opts.Size},
		"response_format": []string{opts.ResponseFormat},
	}
	req, err := e.newReq(ctx, http.MethodPost, uri, "formData", strings.NewReader(postValues.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp ImageEditResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

type ImageVariationOptions struct {
	// The image to edit. Must be a valid PNG file, less than 4MB, and square.
	// If mask is not provided, image must have transparency, which will be used as the mask.
	Image string `binding:"required"`
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `binding:"min=1,max=10"`
	// The size of the generated images.
	// Must be one of 256x256, 512x512, or 1024x1024.
	Size string `binding:"oneof=256x256 512x512 1024x1024"`
	// The format in which the generated images are returned.
	// Must be one of url or b64_json
	ResponseFormat string `binding:"omitempty,oneof=url b64_json"`
}

type ImageVariationResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// ImageVariation creates a variation of a given image.
//
// Docs: https://beta.openai.com/docs/api-reference/images/create-variation
func (e *Engine) ImageVariation(ctx context.Context, opts *ImageVariationOptions) (*ImageCreateResponse, error) {
	if err := e.validate.StructCtx(ctx, opts); err != nil {
		return nil, err
	}
	uri := e.apiBaseURL + "/images/variations"
	if opts.N == 0 {
		opts.N = 1
	}
	if len(opts.Size) == 0 {
		opts.Size = SizeSmall
	}
	if len(opts.ResponseFormat) == 0 {
		opts.ResponseFormat = ResponseFormatUrl
	}
	postValues := url.Values{
		"model":           []string{opts.Image},
		"n":               []string{strconv.Itoa(opts.N)},
		"size":            []string{opts.Size},
		"response_format": []string{opts.ResponseFormat},
	}
	req, err := e.newReq(ctx, http.MethodPost, uri, "formData", strings.NewReader(postValues.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp ImageCreateResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}
