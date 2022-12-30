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

func TestImageCreate(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	r, err := e.ImageCreate(context.Background(), &ImageCreateOptions{
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

func TestImageEdit(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	r, err := e.ImageEdit(context.Background(), &ImageEditOptions{
		Image:  "000test.png",
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

func TestImageVariation(t *testing.T) {
	e := New(os.Getenv("OPENAI_KEY"))
	r, err := e.ImageVariation(context.Background(), &ImageVariationOptions{
		Image: "000test.png",
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
