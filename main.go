package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"server/database"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func init() {
  if err := godotenv.Load(); err!= nil {
    panic(err)
  }
}

func main() {
  dbURL := os.Getenv("DB_URL")
  port := os.Getenv("PORT")
  db := database.GetDBConnectionPool(dbURL)
  _ = db

  r := chi.NewRouter()
  	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", port),
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Println("Server started on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
