//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world password-checker --out gen ./wit
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	gopasswordvalidator "github.com/wagslane/go-password-validator"
	"go.wasmcloud.dev/component/net/wasihttp"

	"github.com/protochron/http-password-checker-go/gen/wasi/blobstore/v0.2.0-draft/blobstore"
	blobstoreTypes "github.com/protochron/http-password-checker-go/gen/wasi/blobstore/v0.2.0-draft/types"
)

type CheckRequest struct {
	Value string `json:"value"`
}

type CheckResponse struct {
	Valid   bool   `json:"valid"`
	Length  int    `json:"length,omitempty"`
	Message string `json:"message,omitempty"`
}

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/check", checkHandler)
	wasihttp.Handle(mux)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResponseJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req CheckRequest
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		errResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(b, &req); err != nil {
		errResponseJSON(w, http.StatusBadRequest, fmt.Sprintf("error with json input: %s", err.Error()))
		return
	}

	badPasswordsList, err := getPasswordList(w)
	if err != nil {
		errResponseJSON(w, http.StatusInternalServerError, fmt.Sprintf("failed to get bad passwords list: %s", err.Error()))
		return
	}

	for pw := range badPasswordsList {
		if req.Value == badPasswordsList[pw] {
			errResponseJSON(w, http.StatusBadRequest, "password is in the list of 500 worst passwords")
			return
		}
	}

	err = gopasswordvalidator.Validate(req.Value, 60)
	if err != nil {
		errResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := CheckResponse{Valid: true, Length: len(req.Value)}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		errResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func errResponseJSON(w http.ResponseWriter, code int, message string) {
	msg, _ := json.Marshal(CheckResponse{Valid: false, Message: message})
	http.Error(w, string(msg), code)
	w.Header().Set("Content-Type", "application/json")
}

func getPasswordList(w http.ResponseWriter) ([]string, error) {
	storeResult := blobstore.GetContainer("")
	st := storeResult.OK()
	if st == nil {
		return nil, fmt.Errorf("failed to get blobstore container")
	}

	if res := st.HasObject("500-worst-passwords.txt"); res.IsErr() {
		return nil, fmt.Errorf("failed to get bad passwords file")
	}

	info := st.ObjectInfo("500-worst-passwords.txt")
	if info.IsErr() {
		return nil, fmt.Errorf("failed to get bad passwords file metadata")
	}

	badPasswordsSize := info.OK().Size
	badPasswords := st.GetData("500-worst-passwords.txt", 0, uint64(badPasswordsSize))
	if badPasswords.IsErr() {
		return nil, fmt.Errorf("failed to get bad passwords file data")
	}

	passwords := badPasswords.OK()
	stream := blobstoreTypes.IncomingValueIncomingValueConsumeAsync(*passwords)
	if stream.IsErr() {
		return nil, fmt.Errorf("failed to consume bad passwords file data")
	}

	data := stream.OK()
	var buf []byte
	for {
		res := data.BlockingRead(4096)
		if err := res.Err(); err != nil {
			if err.Closed() {
				break
			}
			return nil, fmt.Errorf("failed to read bad passwords file data: %s", err.LastOperationFailed().ToDebugString())
		}
		buf = append(buf, res.OK().Slice()...)
	}
	data.ResourceDrop()

	scanner := bufio.NewScanner(strings.NewReader(string(buf)))
	scanner.Split(bufio.ScanLines)
	var badPasswordsList []string
	for scanner.Scan() {
		badPasswordsList = append(badPasswordsList, scanner.Text())
	}

	return badPasswordsList, nil
}

// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}
