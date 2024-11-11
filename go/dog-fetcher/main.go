//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world fetch --out gen ./wit
package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"go.wasmcloud.dev/component/net/wasihttp"
)

const (
	CONTAINER_NAME          = "doggies"
	RANDOM_DOG_API_ENDPOINT = "https://dog.ceo/api/breeds/image/random"
)

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
)

func init() {
	// Register the handleRequest function as the handler for all incoming requests.
	wasihttp.HandleFunc(handleRequest)
}

type RandomDog struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Response struct {
	Filename   string `json:"filename"`
	PathOnDisk string `json:"path_on_disk"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Step 1:
	// Request random dog image from the API
	resp, err := httpRequest(RANDOM_DOG_API_ENDPOINT)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Decode API response into RandomDog struct
	var dog RandomDog
	err = json.NewDecoder(resp.Body).Decode(&dog)
	if err != nil {
		http.Error(w, "failed to decode dog picture api response", http.StatusBadGateway)
		return
	}

	// Create a new request to be sent to fetch the image returned by the random dog API
	resp, err = httpRequest(dog.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Extract the filename from the URL returned by the API
	dogImageName := dog.Message[strings.LastIndex(dog.Message, "/")+1:]

	err = writeObject(resp.Body, CONTAINER_NAME, dogImageName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Setup the response to be returned
	output := Response{
		Filename:   dogImageName,
		PathOnDisk: filepath.Join("/tmp/blobstore", CONTAINER_NAME, dogImageName),
	}

	json.NewEncoder(w).Encode(output)
}

// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}
