package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){

	// load the .env file
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}