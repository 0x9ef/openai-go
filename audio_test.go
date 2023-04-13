package openai

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAudio(t *testing.T) {
	testCases := []struct {
		name              string
		filename          string
		wantTranscription string
		wantTranslation   string
	}{
		{
			name:              "german",
			filename:          "testdata/german.wav",
			wantTranscription: "Ich will das eben wegbringen und dann mit Karl was trinken gehen.",
			wantTranslation:   "I just want to take this away and go out for a drink with Karl.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New(os.Getenv("OPENAI_KEY"))
			file, err := os.Open(tc.filename)
			assert.NoError(t, err)
			defer file.Close()
			buffer, err := io.ReadAll(file)
			assert.NoError(t, err)
			audioOpts := &AudioOptions{
				AudioFormat: strings.Split(tc.filename, ".")[1],
				Model:       ModelWhisper,
				Temperature: 0,
			}
			t.Run("transcribe", func(t *testing.T) {
				audioOpts.File = bytes.NewBuffer(buffer)
				r, err := e.Transcribe(context.Background(), &TranscribeOptions{AudioOptions: audioOpts})
				assert.NoError(t, err)
				assert.Equal(t, tc.wantTranscription, r.Text)
			})
			t.Run("translate", func(t *testing.T) {
				audioOpts.File = bytes.NewBuffer(buffer)
				r, err := e.Translate(context.Background(), &TranslateOptions{AudioOptions: audioOpts})
				assert.NoError(t, err)
				assert.Equal(t, tc.wantTranslation, r.Text)
			})
		})
	}
}
