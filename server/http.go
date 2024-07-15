package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/rs/cors"
)

const (
	PORT = ":8081"
)

type handler struct {
	// gateway
}

func NewHTTPHandler() *handler {
	return &handler{}
}

func (h *handler) HTTPInit() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", h.CheckHealth)
	mux.HandleFunc("POST /api/code/go", h.HandleGo)

	fmt.Printf("Listening on http://localhost%s for connections\n", PORT)

	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(PORT, handler); err != nil {
		log.Fatalf("There was an error in HTTPInit: %s", err)
	}
}

func (h *handler) HandleGo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var body map[string]string

	err := json.NewDecoder(r.Body).Decode(&body)
	check("Error in decoding JSON", err)

	f, err := os.Create("file.go")
	check("Error opening Go file", err)
	defer os.Remove(f.Name())

	err = os.WriteFile(f.Name(), []byte(body["code"]), 0666)
	check("Error in writing to file", err)

	out, _ := exec.Command("go", "run", f.Name()).CombinedOutput()

	cmdOut := string(out[:])

	w.Write([]byte(cmdOut))

}

func check(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

func (h *handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("https://localhost%s is working perfectly fine\n", PORT)
}
