// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package chatgpt

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestCompletion(t *testing.T) {
	e := New(os.Getenv("OPENAPI_KEY"))
	r, err := e.Completion(context.Background(), &CompletionOptions{
		Model:  ModelTextDavinci001,
		Prompt: "Write a little bit of Wikipedia. What is that?",
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
