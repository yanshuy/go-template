package server

import (
	"net/http"

	"github.com/yanshuy/http-web-server/internal/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logger,
		middleware.Recover,
	)

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("GET /health", s.healthCheckHandler)
	mux.HandleFunc("GET /panic", s.handlePanic)

	v1 := http.NewServeMux()
	v1.HandleFunc("GET /item/{id}", s.itemsHandler)

	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	return stack(mux)
}

func (s *Server) handlePanic(w http.ResponseWriter, r *http.Request) {
	var p *int

	*p += 1
}
