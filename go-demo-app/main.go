package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	// Set up the Spin HTTP handler
	spinhttp.Handle(handleRequest)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	w.Header().Set("Content-Type", "text/plain")

	// Write the response message
	fmt.Fprintln(w, "This app is running as WASM!")

	// Initialize the logger
	logger := log.New(os.Stderr, "INFO: ", log.LstdFlags)
	logger.Println("Handled request for:", r.URL.Path)
}

func main() {
	// main function remains empty as Spin handles the HTTP server initialization
}
