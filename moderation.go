// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type ModerationResponse struct {
	Id      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Categories struct {
			// Content that expresses, incites, or promotes hate based on race, gender, ethnicity,
			// religion, nationality, sexual orientation, disability status, or caste.
			Hate bool `json:"hate"`
			// Hateful content that also includes violence or serious harm towards the targeted group.
			HateThreatening bool `json:"hate/threatening"`
			// Content that promotes, encourages, or depicts acts of self-harm, such as suicide,
			// cutting, and eating disorders.
			SelfHarm bool `json:"self-harm"`
			// Content meant to arouse sexual excitement, such as the description of sexual activity,
			// or that promotes sexual services (excluding sex education and wellness).
			Sexual bool `json:"sexual"`
			// Sexual content that includes an individual who is under 18 years old.
			SexualMinors bool `json:"sexual/minors"`
			// Content that promotes or glorifies violence or celebrates the suffering or humiliation of others.
			Violence bool `json:"violence"`
			// Violent content that depicts death, violence, or serious physical injury in extreme graphic detail.
			ViolenceGraphic bool `json:"violence/graphic"`
		} `json:"categories"`
		CategoryScores struct {
			Hate            float64 `json:"hate"`
			HateThreatening float64 `json:"hate/threatening"`
			SelfHarm        float64 `json:"self-harm"`
			Sexual          float64 `json:"sexual"`
			SexualMinors    float64 `json:"sexual/minors"`
			Violence        float64 `json:"violence"`
			ViolenceGraphic float64 `json:"violence/graphic"`
		} `json:"category_scores"`
		Flagged bool `json:"flagged"`
	} `json:"results"`
}

// Moderate classifies if text violates OpenAI's Content Policy
//
// Docs: https://platform.openai.com/docs/api-reference/moderations/create
func (e *Engine) Moderate(ctx context.Context, input string) (*ModerationResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(struct {
		Input string `json:"input"`
	}{Input: input})
	if err != nil {
		return nil, err
	}

	uri := e.apiBaseURL + "/moderations"
	req, err := e.newReq(ctx, http.MethodPost, uri, "json", &buf)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp ModerationResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}
