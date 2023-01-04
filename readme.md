# Golang OpenAI API Client

An Golang native implementation to easily interacting with OpenAI API.

## Usage

You can use environment variable to store API secret key
```
export OPENAI_KEY=[YOUR_KEY]
```

To initialize engine, use this:
```go
e := openai.New(os.Getenv("OPENAI_KEY"))
```
 
### Tips 

#### Model
If you want to use the most powerful model to generate text outputs, ensure that you are using "text-davinci-003". This model is defined as constant `openai.ModelTextDavinci003`.

#### Text edition
You can use the bundle Completion+Edit to regenerate the response based on the last context.
```go
e := openai.New(os.Getenv("OPENAI_KEY"))
ctx := context.Background()
completionResp, err := e.Completion(ctx, &openai.CompletionOptions{
  // Choose model, you can see list of available models in models.go file
	Model:  openai.ModelTextDavinci001, 
  // Number of completion tokens to generate response. By default - 1024
  MaxTokens: 1200, 
  // Text to completion
	Prompt: []string{"Write a little bit of Wikipedia. What is that?"},
}) 

editResp, err := e.Edit(ctx, &EditOptions{
	Model:       ModelTextDavinci001,
	Input:       completionResp.Choices[0],
	Instruction: "Please rewrite a bit more and add more information about Wikipedia in different aspects. Please build based on that for 4 topics",
})
```

### Text completion example 
Given a prompt, the model will return one or more predicted completions. 

**Note**: the default number of completion tokens is 1024, if you want to increase or decrease this limit, you should change `MaxTokens` parameter. 
 
```go
e := openai.New(os.Getenv("OPENAI_KEY"))
r, err := e.Completion(context.Background(), &openai.CompletionOptions{
  // Choose model, you can see list of available models in models.go file
	Model:  openai.ModelTextDavinci001, 
  // Number of completion tokens to generate response. By default - 1024
  MaxTokens: 1200, 
  // Text to completion
	Prompt: []string{"Write a little bit of Wikipedia. What is that?"},
})
```

You will get the next output:
```json
{
  "id": "cmpl-6SrcYDLCVT7xyHKVNuSLNuhRvwOJ1",
  "object": "text_completion",
  "created": 1672337322,
  "model": "text-davinci-001",
  "choices": [
    {
      "text": "\n\nWikipedia is a free online encyclopedia, created and edited by volunteers.",
      "index": 0,
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 11,
    "completion_tokens": 15,
    "total_tokens": 26
  }
}
```

To get only the text you should use the next code:
```go
fmt.Println(r.Choices[0].Text)
```

So, the full code will be:
```go
package main 


import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

    "github.com/0x9ef/openai-go"
) 

func main() {
    e := openai.New(os.Getenv("OPENAI_KEY"))
    r, err := e.Completion(context.Background(), &openai.CompletionOptions{
        // Choose model, you can see list of available models in models.go file
        Model:  openai.ModelTextDavinci001, 
        // Text to completion
        Prompt: []string{"Write a little bit of Wikipedia. What is that?"}
    })

	if b, err := json.MarshalIndent(r, "", "  "); err != nil {
		panic(err)
	} else {
		fmt.Println(string(b))
	}

    // Wikipedia is a free online encyclopedia, created and edited by volunteers.
    fmt.Println("What is the Wikipedia?", r.Choices[0].Text)
}
```

### Models list/retrieve 
Lists the currently available models, and provides basic information about each one such as the owner and availability.

```go
e := openai.New(os.Getenv("OPENAI_KEY"))
r, err := e.ListModels(context.Background())
if err != nil {
	log.Fatal(err)
}
```

You will get the next output:
```json
{
  "data": [
    {
      "id": "babbage",
      "object": "model",
      "owned_by": "openai"
    },
    {
      "id": "ada",
      "object": "model",
      "owned_by": "openai"
    },
    {
      "id": "text-davinci-002",
      "object": "model",
      "owned_by": "openai"
    },
    {
      "id": "davinci",
      "object": "model",
      "owned_by": "openai"
    },
    ...
  ]
}
``` 

To retrieve information about specified model instead of all models, you can do this:

```go
e := openai.New(os.Getenv("OPENAI_KEY"))
r, err := e.RetrieveModel(context.Background(), &openai.RetrieveModelOptions{
		ID: openai.ModelDavinci,
	})
if err != nil {
	log.Fatal(err)
}
```

## License

[MIT](./LICENSE)
