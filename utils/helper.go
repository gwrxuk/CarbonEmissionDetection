package utils

import (
  "context"
  "fmt"
  "log"
  "os"
	"io"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
  "github.com/google/generative-ai-go/genai"
  "google.golang.org/api/option"
  "strings"
  "github.com/joho/godotenv"
)

func uploadToGemini(path, mimeType string, client *genai.Client, ctx context.Context) string {
  file, err := os.Open(path)
  if err != nil {
    log.Fatalf("Error opening file: %v\n", err)
  }
  defer file.Close()

  options := genai.UploadFileOptions{
    DisplayName: path,
    MIMEType:    mimeType,
  }
  fileData, err := client.UploadFile(ctx, "", file, &options)
  if err != nil {
    log.Fatalf("Error uploading file: %v\n", err)
  }

  log.Printf("Uploaded file %s as: %s\n", fileData.DisplayName, fileData.URI)
  return fileData.URI
}

func Predict(url string) string {
  ctx := context.Background()
   // Load environment variables from .env file
   err := godotenv.Load()
   if err != nil {
       log.Fatal("Error loading .env file")
   }

  apiKey := os.Getenv("GEMINI_API_KEY")
  if apiKey == "" {
    log.Fatal("Environment variable GEMINI_API_KEY not set\n")
  }

  fileName := download(url)
  option := option.WithAPIKey(apiKey)

  client, err := genai.NewClient(ctx, option)
  if err != nil {
    log.Fatalf("Error creating client: %v\n", err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")

  model.SetTemperature(0.9)
  model.SetTopK(64)
  model.SetTopP(0.95)
  model.SetMaxOutputTokens(1024)
  model.ResponseMIMEType = "text/plain"

  // model.SafetySettings = Adjust safety settings
  // See https://ai.google.dev/gemini-api/docs/safety-settings

  // TODO Make these files available on the local file system
  // You may need to update the file paths
  fileUris := []string{
    uploadToGemini(fileName, "image/jpeg", client, ctx),
  }

  err = os.RemoveAll(fileName)
  if err != nil {
      fmt.Println("Error deleting directory:", err)
  } else {
      fmt.Println("Directory successfully deleted")
  }

  session := model.StartChat()
  session.History = []*genai.Content{
    {
      Role: "user",
      Parts: []genai.Part{
        genai.Text("From the image, are there carbon emission activities? "),
        genai.FileData{URI: fileUris[0]},
      },
    },
  }

  resp, err := session.SendMessage(ctx, genai.Text("Describe the detail."))
  if err != nil {
    log.Fatalf("Error sending message: %v\n", err)
  }

  var responseText strings.Builder 

  for _, part := range resp.Candidates[0].Content.Parts {
    fmt.Printf("%v\n", part)
   responseText.WriteString(fmt.Sprintf("%v\n", part)) 
  }

  return responseText.String();
}


func download(imageURL string) string {

	// Fetch the image
	response, err := http.Get(imageURL)
	if err != nil {
		log.Fatalf("Failed to fetch image: %v", err)
	}
	defer response.Body.Close()

	hash := sha256.New()

	// Write the input string to the hash
	hash.Write([]byte(imageURL))

	// Calculate the SHA256 checksum
	hashInBytes := hash.Sum(nil)

	// Convert the checksum to a hexadecimal string
	fileName := hex.EncodeToString(hashInBytes)+".jpg";
	// Create a file to save the image
	file, err := os.Create("static/images/"+fileName);
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Copy the data from the response to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	log.Println("Image saved successfully!")
  fileName = "static/images/" + fileName
  return fileName
}
