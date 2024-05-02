package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

func runVertexAI(model string, prompt string, file string, fileMime string) {
	ctx := context.Background()
	projectId := os.Getenv("GOOGLE_PROJECT_ID")
	region := os.Getenv("GOOGLE_REGION")

	client, err := genai.NewClient(ctx, projectId, region)
	if err != nil {
		log.Fatal("client connect error", err)
	}
	defer client.Close()

	fd, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("read file error", err)
	}

	parts := []genai.Part{
		// genai.FileData{MIMEType: fileMime, FileURI: fileURI},
		genai.Blob{MIMEType: fileMime, Data: fd},
		genai.Text(prompt),
	}

	m := client.GenerativeModel(model) // "gemini-1.5-pro-preview-0409"
	m.SetTemperature(0.0)
	m.SetTopP(0.95)

	resp, err := m.GenerateContent(ctx, parts...)
	if err != nil {
		log.Fatal("generate content error", err)
	}

	bs, _ := json.MarshalIndent(resp, "", "    ")
	// out := append([]byte(fmt.Sprintf("INSTRUCTION:\"%s\"\n", instruction)), bs...)
	// os.WriteFile(page+".out", out, 0644)

	fmt.Println("Vertex AI response:")
	fmt.Println(string(bs))
}
