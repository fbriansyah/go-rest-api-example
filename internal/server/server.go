package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"

	"api-example/internal/service"
)

type Server struct {
	port        int
	userHandler *UserHandler
}

type Services struct {
	UserService service.IUserService
}

func NewServer(srv Services) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	validate := validator.New(validator.WithRequiredStructEnabled())

	NewServer := &Server{
		port: port,

		userHandler: NewUserHandler(validate, srv.UserService),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	slog.Info("Server start", slog.Int("Port", port))

	return server
}
