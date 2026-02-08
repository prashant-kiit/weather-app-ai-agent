package main

import (
	"fmt"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"os"
	"context"
)

func main() {
	fmt.Println("Weather App AI Agent")

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-5.2",
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String("Hello, how are you?"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp.OutputText())
}
