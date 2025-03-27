package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/yanshuy/http-web-server/internal/database"
	"github.com/yanshuy/http-web-server/internal/repository"
	"github.com/yanshuy/http-web-server/internal/service"
)

type Server struct {
	port    int
	db      database.Service
	repo    *repository.Queries
	service *service.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()
	repo := repository.New(db.GetDB())

	NewServer := &Server{
		port:    port,
		db:      db,
		repo:    repo,
		service: service.New(repo),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
