package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lfcifuentes/clean-arquitecture/internal/handler"
	"github.com/lfcifuentes/clean-arquitecture/internal/repository"
	"github.com/lfcifuentes/clean-arquitecture/internal/usecase"
	"github.com/lfcifuentes/clean-arquitecture/pkg/db"
	"github.com/spf13/cobra"
)

func init() {
	// Add the serve command to the root command
	rootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the REST API server",
	Long:  `Start the REST API server that listens on the specified host and port`,
	Run: func(_ *cobra.Command, _ []string) {
		// Call the serve function
		serve()
	},
}

func serve() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	dns := db.GeneratePostgresConnString()

	conn, err := db.NewPostgresConn(dns)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	userRepo := repository.NewUserPostgresRepository(conn)
	userCase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userCase)

	r.Get("/users", userHandler.UserList)
	r.Post("/users", userHandler.CreateUser)

	log.Println("Server is running on :8080")
	log.Fatal("Server error: ", http.ListenAndServe(":8080", r))
}
