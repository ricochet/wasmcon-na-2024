//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world hello --out gen ./wit
package main

import (
	"io"
	"net/http"

	"go.wasmcloud.dev/component/net/wasihttp"
)

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
)

func init() {
	// Register the handleRequest function as the handler for all incoming requests.
	wasihttp.HandleFunc(handleRequest)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	url := "https://dog.ceo/api/breeds/image/random"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusBadGateway)
		return
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, "failed to make outbound request", http.StatusBadGateway)
		return
	}

	w.Header().Set("X-Custom-Header", "proxied")
	w.WriteHeader(resp.StatusCode)

	_, _ = io.Copy(w, resp.Body)

	// // Get dog picture
	// let dog_picture = make_outgoing_request(&dog_response.message)?;
	// // TODO: blobstore
	// Ok(Response::new(dog_picture))
}

// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}
