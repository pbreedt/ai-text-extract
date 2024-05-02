package main

const (
	mimePdf string = "application/pdf"
	mimeJpg string = "image/jpeg"
	mimePng string = "image/png"
)

func main() {
	file := "./data/test.png"
	fileMime := mimePdf

	// Useful with PDF doc:
	//   'top to bottom' - provide better output when large spaces between text, like tables or lists
	//   'Do not add any word of your own' - prevents intro text like "## The text from the image ##"
	//   'Detect any text in the provided image' - For some images, Vertex needed to first describe the image before the text can be extracted!
	// prompt := "Provide all the text in the document exactly as it appears, from top to bottom. Do not add any word of your own, only text from the document. Find text in any graphics, images, diagrams, charts, etc."
	// prompt := "Detect any text in the provided image. Provide this text exactly as it appears, from top to bottom. Do not add any word of your own."
	prompt := "Examine this image looking for any text in the image, then provide only the text that came directly from the image, including the layout. Do not add any word of your own, only text from the image."
	model := "gemini-1.5-pro-latest"
	runGenAI(model, prompt, file, fileMime)
	model = "gemini-1.5-pro-preview-0409"
	runVertexAI(model, prompt, file, fileMime)
}
