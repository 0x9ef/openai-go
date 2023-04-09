package openai

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestTranscribe(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	audioFile, err := os.Open("testdata/german.wav")
	if err != nil {
		t.Fatal(err)
	}
	r, err := e.Transcribe(context.Background(), &TranscribeOptions{
		AudioOptions: &AudioOptions{
			File:        audioFile,
			AudioFormat: "wav",
			Model:       ModelWhisper,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if b, err := json.MarshalIndent(r, "", "  "); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}

func TestTranslate(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	audioFile, err := os.Open("testdata/german.wav")
	if err != nil {
		t.Fatal(err)
	}
	r, err := e.Translate(context.Background(), &TranslateOptions{
		AudioOptions: &AudioOptions{
			File:        audioFile,
			AudioFormat: "wav",
			Model:       ModelWhisper,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if b, err := json.MarshalIndent(r, "", "  "); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}
