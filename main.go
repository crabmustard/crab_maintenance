package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

type crabConfig struct {
	port      string
	stuffRoot string
}

var cfg crabConfig

func main() {
	godotenv.Load(".env")

	cfg.stuffRoot = os.Getenv("STUFF_ROOT")
	if cfg.stuffRoot == "" {
		log.Fatal("STUFF_ROOT env variable not set")
	}
	cfg.port = os.Getenv("PORT")
	if cfg.port == "" {
		log.Fatal("PORT env variable not set")
	}
	mux := http.NewServeMux()
	stuffHandler := http.StripPrefix("/stuff", http.FileServer(http.Dir(cfg.stuffRoot)))
	mux.Handle("/stuff/", stuffHandler)
	mux.Handle("/", templ.Handler(homePage()))
	mux.Handle("/tickets", templ.Handler(ticketPage()))

	srv := &http.Server{
		Addr:    ":" + cfg.port,
		Handler: mux,
	}

	fmt.Printf("listenning on port %s\n", cfg.port)
	log.Fatal(srv.ListenAndServe())
}
