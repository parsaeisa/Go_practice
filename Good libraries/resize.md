# Resize

The further package is used for resizing the served image by backend server to the client. 

The package is "github.com/nfnt/resize".

Example:
```go
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func handleImage(w http.ResponseWriter, r *http.Request) {
	// Open the image file
	file, err := os.Open("path/to/your/image.jpg")
	if err != nil {
		http.Error(w, "Error opening image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	// Resize the image to your desired dimensions
	width := 300 // Adjust to your preferred width
	height := 200 // Adjust to your preferred height
	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// Encode the resized image and write it to the response
	w.Header().Set("Content-Type", "image/jpeg")
	if err := jpeg.Encode(w, resizedImg, nil); err != nil {
		http.Error(w, "Error encoding image", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/image", handleImage)
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
```

It takes new size from (can be read from config) 