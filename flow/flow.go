package flow

import (
	"context"
	"fmt"
	"log"

	_ "embed"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/dotprompt"
	"github.com/firebase/genkit/go/plugins/googleai"
	"github.com/invopop/jsonschema"
)

//go:embed summarize.prompt
var promptTemplate string

type promptInput struct {
	URL string `json:"url"`
}

func DefineFlow(ctx context.Context) *genkit.Flow[string, string, struct{}] {
	// Initialize the Google AI plugin
	if err := googleai.Init(ctx, nil); err != nil {
		log.Fatalf("Failed to initialize Google AI plugin: %v", err)
	}

	model := googleai.Model("gemini-1.5-flash")

	// Define the webLoader tool
	webLoader := ai.DefineTool(
		"webLoader",
		"Loads a webpage and returns the textual content.",
		func(ctx context.Context, input struct {
			URL string `json:"url"`
		}) (string, error) {
			return fetchWebContent(input.URL)
		},
	)

	summarizePrompt, err := dotprompt.Define("summarizePrompt",
		promptTemplate,
		dotprompt.Config{
			Model: model,
			Tools: []ai.Tool{webLoader},
			GenerationConfig: &ai.GenerationCommonConfig{
				Temperature: 1,
			},
			InputSchema:  jsonschema.Reflect(promptInput{}),
			OutputFormat: ai.OutputFormatText,
		},
	)
	if err != nil {
		log.Fatalf("Failed to initialize prompt: %v", err)
	}

	// Define a flow that fetches a webpage and summarizes its content
	return genkit.DefineFlow("summarizeFlow", func(ctx context.Context, input string) (string, error) {
		resp, err := summarizePrompt.Generate(ctx,
			&dotprompt.PromptRequest{
				Variables: &promptInput{
					URL: input,
				},
			},
			nil,
		)
		if err != nil {
			return "", fmt.Errorf("failed to generate summary: %w", err)
		}
		return resp.Text(), nil
	})
}
