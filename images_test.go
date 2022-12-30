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
		Size:   SizeSmall,
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
		Image:  "https://oaidalleapiprodscus.blob.core.windows.net/private/org-eSqeiZkMEDQsrwhfqU2bgbwd/user-KNsBZQnL5BwvITLR3rihZvNx/img-R0loYA2FUFBUGvkewj2E52fJ.png?st=2022-12-30T21%3A07%3A15Z\u0026se=2022-12-30T23%3A07%3A15Z\u0026sp=r\u0026sv=2021-08-06\u0026sr=b\u0026rscd=inline\u0026rsct=image/png\u0026skoid=6aaadede-4fb3-4698-a8f6-684d7786b067\u0026sktid=a48cca56-e6da-484e-a814-9c849652bcb3\u0026skt=2022-12-30T20%3A27%3A06Z\u0026ske=2022-12-31T20%3A27%3A06Z\u0026sks=b\u0026skv=2021-08-06\u0026sig=Freo5NF9iVtQvs0KXhJRhIjzRmV7CAmfDvaVEHJqo3w%3D",
		Prompt: "Write a little bit of Wikipedia. What is that?",
		Size:   SizeSmall,
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
		Image: "https://oaidalleapiprodscus.blob.core.windows.net/private/org-eSqeiZkMEDQsrwhfqU2bgbwd/user-KNsBZQnL5BwvITLR3rihZvNx/img-R0loYA2FUFBUGvkewj2E52fJ.png?st=2022-12-30T21%3A07%3A15Z\u0026se=2022-12-30T23%3A07%3A15Z\u0026sp=r\u0026sv=2021-08-06\u0026sr=b\u0026rscd=inline\u0026rsct=image/png\u0026skoid=6aaadede-4fb3-4698-a8f6-684d7786b067\u0026sktid=a48cca56-e6da-484e-a814-9c849652bcb3\u0026skt=2022-12-30T20%3A27%3A06Z\u0026ske=2022-12-31T20%3A27%3A06Z\u0026sks=b\u0026skv=2021-08-06\u0026sig=Freo5NF9iVtQvs0KXhJRhIjzRmV7CAmfDvaVEHJqo3w%3D",
		Size:  SizeSmall,
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
