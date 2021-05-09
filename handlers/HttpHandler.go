package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const httpHandlerLog = "[HTTP HANDLER]"

type HttpHandler struct{}

func (h *HttpHandler) Handling() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(h.getPort(), nil)
	log.Println(httpHandlerLog + "listening on port " + h.getPort())
	if err != nil {
		log.Fatal(err)
	}
}

func (m *HttpHandler) getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println(httpHandlerLog + "INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, httpHandlerLog+"start")
}
