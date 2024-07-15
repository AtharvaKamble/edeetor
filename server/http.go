package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const (
	PORT = ":8081"
)

type handler struct {
	// gateway
}

type Resp struct {
	code string
}

func NewHTTPHandler() *handler {
	return &handler{}
}

func (h *handler) HTTPInit() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", h.CheckHealth)
	mux.HandleFunc("POST /api/code/go", h.HandleGo)

	fmt.Printf("Listening on http://localhost%s for connections\n", PORT)

	if err := http.ListenAndServe(PORT, mux); err != nil {
		log.Fatalf("There was an error in HTTPInit: %s", err)
	}
}

func (h *handler) HandleGo(w http.ResponseWriter, r *http.Request) {
	var body map[string]string

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Fatalf("Error in decoding JSON: %s\n", err)
	}

	fmt.Println(body["code"])

	out, err := exec.Command("ls").Output()

	if err != nil {
		log.Fatalf("There was an error executing Go code: %s\n", err)
	}

	fmt.Println(string(out[:]))

}

func (h *handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("https://localhost%s is working perfectly fine\n", PORT)
}
