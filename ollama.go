package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ollama/ollama/api"
)

// Requires ollama running locally: https://ollama.com/
// ollama run llava
/* NOTE: llava model does not perform nearly as well as gemini which is (currently) not available via ollama.
Output for the sample data form:
```
Standard Data Collection Form

Patient Information
----------------------

Name: _______________________
Date of Birth: _______________________
Age: _______________________
Gender: Male/Female/Other: _______________________
Race: White/Black/Hispanic/Other: _______________________
Insurance Provider: _______________________
Policy Number: _______________________
Group Number: _______________________
Patient ID: _______________________

Physician Information
----------------------------

Name: _______________________
Phone: _______________________
Fax: _______________________
Email: _______________________

Hospital Information
-----------------------------

Name: _______________________
Address: _______________________
Phone: _______________________
Fax: _______________________

Patient's Signature: _______________________
Date: _______________________
````
*/

func runOllama(model string, prompt string, file string, fileMime string) {
	// new [Client] using env var OLLAMA_HOST <scheme>://<host>:<port>
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	fd, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("read file error", err)
	}

	fileData := api.ImageData(fd)

	req := &api.GenerateRequest{
		Model:  model, // "llava"
		Images: []api.ImageData{fileData},
		Prompt: prompt,
		Stream: new(bool), // disable streaming
		Options: map[string]interface{}{
			"temperature": 0.0,
			"top_p":       0.95,
		},
	}

	ctx := context.Background()
	var response string
	respFunc := func(resp api.GenerateResponse) error {
		response = resp.Response
		resp.Summary()
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ollama response:")
	fmt.Println(response)
}

func sample() {
	modelName := flag.String("model", "", "ollama model name")
	flag.Parse()

	// new [Client] using env var OLLAMA_HOST <scheme>://<host>:<port>
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GenerateRequest{
		Model:  *modelName,
		Prompt: flag.Args()[0],
		Stream: new(bool), // disable streaming
	}

	ctx := context.Background()
	var response string
	respFunc := func(resp api.GenerateResponse) error {
		response = resp.Response
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:\n", response)
}
