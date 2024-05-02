package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func runGenAI(model string, prompt string, file string, fileMime string) {
	if fileMime == mimePdf {
		log.Fatalln("PDF format is not (currently) supported by Google generative-ai library")
	}

	ctx := context.Background()
	aiApiKey := os.Getenv("GOOGLE_AI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(aiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fd, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("read file error", err)
	}

	m := client.GenerativeModel(model) //"gemini-1.5-pro-latest"
	m.SetTemperature(0.0)

	parts := []genai.Part{
		genai.Blob{MIMEType: fileMime, Data: fd},
		genai.Text(prompt),
	}
	resp, err := m.GenerateContent(ctx, parts...)
	if err != nil {
		log.Fatal("generate content error", err)
	}

	bs, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println("Gen AI response:")
	fmt.Println(string(bs))
}

func listModels(client *genai.Client) {
	it := client.ListModels(context.Background())
	for model, err := it.Next(); err != iterator.Done; model, err = it.Next() {
		log.Println(model.Name, model.DisplayName, model.BaseModelID)
	}
}

func listFiles(client *genai.Client) {
	it := client.ListFiles(context.Background())
	for file, err := it.Next(); err != iterator.Done; file, err = it.Next() {
		log.Println(file.Name, file.URI)
	}
}

func uploadFile(client *genai.Client, filepath string) (*genai.File, error) {
	// Use Client.UploadFile to Upload a file to the service.
	// Pass it an io.Reader.
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	log.Println("File read")

	// You can choose a name, or pass the empty string to generate a unique one.
	file, err := client.UploadFile(context.Background(), "", f, nil)
	if err != nil {
		return nil, err
	}
	log.Println("File uploaded")

	return file, nil
}

func getFile(client *genai.Client) *genai.File {
	file, err := client.GetFile(context.Background(), "3oq83pynllhd")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File downloaded")

	return file
}

func deleteFiles(client *genai.Client) {
	ctx := context.Background()
	it := client.ListFiles(ctx)
	for file, err := it.Next(); err != iterator.Done; file, err = it.Next() {
		log.Println(file.Name, file.URI)
		client.DeleteFile(ctx, file.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Files deleted")
}
