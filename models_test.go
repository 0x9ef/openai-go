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

func TestListModels(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	r, err := e.ListModels(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if b, err := json.MarshalIndent(r, "", "  "); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}

func TestRetrieveModel(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	r, err := e.RetrieveModel(context.Background(), &RetrieveModelOptions{
		ID: ModelDavinci,
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
