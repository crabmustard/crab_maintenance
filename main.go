package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/crabmustard/crab_maintenance/database"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

type crabConfig struct {
	port      string
	stuffRoot string
	db        *database.Queries
}

var cfg crabConfig
var red templ.Component

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
	dbPath := os.Getenv("DB_PATH")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("error opening db")
	}
	dbQueries := database.New(db)
	cfg.db = dbQueries

	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(indexPage()))
	stuffHandler := http.StripPrefix("/stuff", http.FileServer(http.Dir(cfg.stuffRoot)))
	mux.Handle("/stuff/", stuffHandler)

	srv := &http.Server{
		Addr:    ":" + cfg.port,
		Handler: mux,
	}

	fmt.Printf("listenning on port %s\n", cfg.port)
	log.Fatal(srv.ListenAndServe())
}
