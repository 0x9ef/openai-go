package openai

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
)

type AudioOptions struct {
	// The audio file to process, in one of these formats:
	// mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	File io.Reader `binding:"required"`
	// The format of the audio file.
	AudioFormat string `binding:"required"`
	// ID of the model to use. Only whisper-1 is currently available.
	Model Model `binding:"required"`
	// An optional text to guide the model's style or continue a previous audio segment.
	// The prompt should match the audio language for transcriptions and English for translations.
	Prompt string
	// The sampling temperature, between 0 and 1.
	// Higher values like 0.8 will make the output more random, while lower values
	// like 0.2 will make it more focused and deterministic.
	// If set to 0, the model will use log probability to automatically increase
	// the temperature until certain thresholds are hit.
	Temperature float32
}

type TranscribeOptions struct {
	*AudioOptions
	// The language of the input audio. Supplying the input language in ISO-639-1
	// format will improve accuracy and latency.
	Language string
}

type TranscribeResponse struct {
	Text string `json:"text"`
}

// Transcribe audio into the input language.
//
// Docs: https://platform.openai.com/docs/api-reference/audio/create
func (e *Engine) Transcribe(ctx context.Context, options *TranscribeOptions) (*TranscribeResponse, error) {
	if err := e.validate.StructCtx(ctx, options); err != nil {
		return nil, err
	}

	url := e.apiBaseURL + "/audio/transcriptions"

	body, contentType, err := newTranscribeBody(options)
	if err != nil {
		return nil, err
	}
	req, err := e.newReq(ctx, "POST", url, contentType, body)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp TranscribeResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

func newTranscribeBody(options *TranscribeOptions) (io.Reader, string, error) {
	body := &bytes.Buffer{}
	writer, err := newMultiPartWriter(body, options.AudioOptions)
	if err != nil {
		return nil, "", err
	}
	if options.Language != "" {
		if err := writer.WriteField("language", options.Language); err != nil {
			return nil, "", fmt.Errorf("write language: %w", err)
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("close writer: %w", err)
	}
	return body, writer.FormDataContentType(), nil
}

type TranslateOptions struct {
	*AudioOptions
}

type TranslateResponse struct {
	Text string `json:"text"`
}

// Translate audio into English.
//
// Docs: https://platform.openai.com/docs/api-reference/audio/create
func (e *Engine) Translate(ctx context.Context, options *TranslateOptions) (*TranslateResponse, error) {
	if err := e.validate.StructCtx(ctx, options); err != nil {
		return nil, err
	}

	url := e.apiBaseURL + "/audio/translations"

	body, contentType, err := newTranslateBody(options)
	if err != nil {
		return nil, err
	}

	req, err := e.newReq(ctx, "POST", url, contentType, body)
	if err != nil {
		return nil, err
	}
	resp, err := e.doReq(req)
	if err != nil {
		return nil, err
	}
	var jsonResp TranslateResponse
	if err := unmarshal(resp, &jsonResp); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

func newTranslateBody(options *TranslateOptions) (io.Reader, string, error) {
	body := &bytes.Buffer{}
	writer, err := newMultiPartWriter(body, options.AudioOptions)
	if err != nil {
		return nil, "", err
	}
	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("close writer: %w", err)
	}
	return body, writer.FormDataContentType(), nil
}

func newMultiPartWriter(body *bytes.Buffer, options *AudioOptions) (*multipart.Writer, error) {
	writer := multipart.NewWriter(body)
	if err := writer.WriteField("model", string(options.Model)); err != nil {
		return nil, fmt.Errorf("write model: %w", err)
	}
	// TODO: support additional response data via verbose_json? what about other formats (vtt, srt)?g
	if err := writer.WriteField("response_format", "json"); err != nil {
		return nil, fmt.Errorf("write response format: %w", err)
	}
	file, err := writer.CreateFormFile("file", "file."+options.AudioFormat)
	if err != nil {
		return nil, fmt.Errorf("create form file: %w", err)
	}
	if _, err := io.Copy(file, options.File); err != nil {
		return nil, fmt.Errorf("copy file: %w", err)
	}
	if options.Prompt != "" {
		if err := writer.WriteField("prompt", options.Prompt); err != nil {
			return nil, fmt.Errorf("write prompt: %w", err)
		}
	}
	if options.Temperature != 0 {
		if err := writer.WriteField("temperature", fmt.Sprintf("%f", options.Temperature)); err != nil {
			return nil, fmt.Errorf("write temperature: %w", err)
		}
	}
	return writer, nil
}
