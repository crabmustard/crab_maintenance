package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/crabmustard/crab_maintenance/database"
	"github.com/joho/godotenv"

	_ "modernc.org/sqlite"
)

type crabConfig struct {
	port      string
	filesRoot string
	db        *database.Queries
}

var cfg crabConfig

func main() {
	godotenv.Load(".env")
	cfg.filesRoot = os.Getenv("FILES_ROOT")
	if cfg.filesRoot == "" {
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

	err = cfg.db.ClearPtacList(context.Background())
	if err != nil {
		log.Fatal("error clearing db")
	}

	createPtacList()

	program := tea.NewProgram(NewPtacService())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error encountered %v", err)
		os.Exit(1)
	}
	// mux := http.NewServeMux()
	// mux.Handle("/", templ.Handler(ptacList(ptacs)))
	// stuffHandler := http.StripPrefix("/stuff", http.FileServer(http.Dir(cfg.filesRoot)))
	// mux.Handle("/stuff/", stuffHandler)

	// srv := &http.Server{
	// 	Addr:    ":" + cfg.port,
	// 	Handler: mux,
	// }

	// fmt.Printf("listenning on port %s\n", cfg.port)
	// log.Fatal(srv.ListenAndServe())
}
