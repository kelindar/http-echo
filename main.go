package main

import (
	"math/rand"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Create our server
	logger := log.New()
	server := Server{
		logger: logger,
	}

	// Start the server
	server.ListenAndServe()
}

// Server represents our server.
type Server struct {
	logger *log.Logger
}

// ListenAndServe starts the server
func (s *Server) ListenAndServe() {
	s.logger.Info("echo server is starting on port 8080...")
	http.HandleFunc("/", s.echo)
	http.ListenAndServe(":8080", nil)
}

// Echo echos back the request as a response
func (s *Server) echo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	// 30% chance of failure
	if rand.Intn(100) < 30 {
		writer.WriteHeader(500)
		writer.Write([]byte("a chaos monkey broke your server"))
		return
	}

	// Happy path
	writer.WriteHeader(200)
	request.Write(writer)
}
