package api

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type server struct {
	logger  *slog.Logger
	address string
}

type Server interface {
	RunServer()
}

func NewServer(
	domain string,
	port int,
	logger *slog.Logger,
) (Server, error) {
	// Remove http:// if provided in the address
	address := fmt.Sprintf("%s:%d", domain, port)
	newServer := &server{
		logger:  logger,
		address: strings.TrimPrefix(address, "http://"),
	}
	return newServer, nil
}

func (s *server) RunServer() {
	s.logger.Info("server is starting")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /submit/input", func(response http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			s.logger.Error(
				"server failed to to ready response body",
				slog.String("error", err.Error()),
			)
		}

		s.logger.Info(
			"server has received a request",
			slog.String("endpoint", "/submit/input"),
			slog.String("requestType", "GET"),
			slog.String("responseBody", string(body)),
		)
		response.Write(
			append([]byte("received message: "), body...),
		)
	})

	mux.HandleFunc("POST /submit/input", func(response http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			s.logger.Error(
				"server failed to to ready response body",
				slog.String("error", err.Error()),
			)
		}

		s.logger.Info(
			"server has received a request",
			slog.String("endpoint", "/submit/input"),
			slog.String("requestType", "POST"),
			slog.String("responseBody", string(body)),
		)
		response.Write(
			append([]byte("received message: "), body...),
		)
	})

	err := http.ListenAndServe(s.address, mux)
	if err != nil {
		s.logger.Error(
			"server failed to run",
			slog.String("address", s.address),
			slog.String("error", err.Error()),
		)
	}

}
