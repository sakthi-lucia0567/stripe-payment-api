package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// load the .env file
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	v1Router := chi.NewRouter()

	v1Router.Get("/healthcheck", handlerReadiness)

	v1Router.Get("/err", handlerErr)

	v1Router.Post("/create_intent", handlerCreateIntent)
	v1Router.Post("/capture_intent/{id}/capture", handlerCaptureIntent)
	v1Router.Post("/create_refund/{id}", handlerCreateRefund)
	v1Router.Get("/get_intents", handlerGetIntents)

	router.Mount("/api/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is running %v", port)
	err := srv.ListenAndServe()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	if err != nil {
		log.Fatal(err)
	}

}
