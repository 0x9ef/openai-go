// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestCompletion(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	r, err := e.Completion(context.Background(), &CompletionOptions{
		Model: DefaultModel,
		Prompt: []string{`Write a thorough blog post outline with at least 8 sections and a unique structure for a blog post titled “The global recession cases & consequences”. 
		Avoid lists. Do not be repetitive. The tone should be educational. The audience is intermediate in the subject. Other information to know is “Will be a global recession in 2023? 
		Or has this process already started?”. In your response, do not say anything additional, just provide the outline.`},
	})
	if err != nil {
		log.Fatal(err)
	}
	if b, err := json.MarshalIndent(r, "", "  "); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}
