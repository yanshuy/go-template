package server

import (
	"net/http"
	"time"

	"github.com/yanshuy/http-web-server/pkg"
)

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"status":    "available",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	pkg.RespondWithJSON(w, http.StatusOK, resp)
}

func (s *Server) itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	idStr := r.PathValue("id")

	// Return sample response
	resp := map[string]interface{}{
		"id":   idStr,
		"name": "Sample Item",
	}

	pkg.RespondWithJSON(w, http.StatusOK, resp)
}
