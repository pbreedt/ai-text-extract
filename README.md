# ai-text-extract
> A package to wrap services from different cloud providers in a common, simplified interface

## General info
The main purpose of this package was to explore the SDK's of different AI providers.  
There are many other possibile tools to also perform this task, like [these](https://cloud.google.com/document-ai/docs/enterprise-document-ocr)

## Technologies
* Go - version 1.22

SDK's currently covered:  
- [Google generative AI](https://ai.google.dev/gemini-api/docs) (Google AI API) [SDK](https://github.com/google/generative-ai-go)
- [Google Vertex AI](https://cloud.google.com/vertex-ai/generative-ai/docs/learn/overview) (Google Cloud project-based) [SDK](https://github.com/google/generative-ai-go)
- [Ollama](https://ollama.com/) (for hosting LLM's locally) [SDK](https://github.com/ollama/ollama)

## Setup
Like with all Go modules, you can simply "go get" it.  
```go get github.com/pbreedt/ai-text-extract```

This package is not very flexible like taking CLI params (mainly because it is just a simple proof-of-concept and the incompatible model names between SDK's)

### Google generative AI
1. Next step creates a new key and project in your Google Cloud account
2. Obtain an API key, see [this link](https://aistudio.google.com/app/apikey)
3. export GOOGLE_AI_API_KEY=your_api_key

### Google Vertex AI
1. Login to GCP console
2. Create service account
3. Create service account key (storing key in /path/to/sa-json.json file)
4. export GOOGLE_APPLICATION_CREDENTIALS=/path/to/sa-json.json
5. export GOOGLE_PROJECT_ID=your_project_id
6. export GOOGLE_REGION=google_region

### Ollama
1. Download and install from the [Ollama website](https://ollama.com/download)
2. Run command 'ollama run llava'

## Status
Project is: _in progress_

## Credits
Thanks to [ritaly](https://github.com/ritaly/README-cheatsheet) for a quick readme template

## Contact
Created by [@pbreedt](mailto:petrus.breedt@gmail.com)
