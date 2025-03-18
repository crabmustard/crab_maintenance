package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	_ "modernc.org/sqlite"
)

// type crabConfig struct {
// 	port      string
// 	filesRoot string
// 	db        *database.Queries
// }

// var cfg crabConfig
// var numRooms int64 = 159

func main() {

	program := tea.NewProgram(initialPtacModel())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error encountered %v", err)
		os.Exit(1)
	}
	// godotenv.Load(".env")
	// cfg.filesRoot = os.Getenv("FILES_ROOT")
	// if cfg.filesRoot == "" {
	// 	log.Fatal("STUFF_ROOT env variable not set")
	// }
	// cfg.port = os.Getenv("PORT")
	// if cfg.port == "" {
	// 	log.Fatal("PORT env variable not set")
	// }
	// dbPath := os.Getenv("DB_PATH")
	// db, err := sql.Open("sqlite", dbPath)
	// if err != nil {
	// 	log.Fatal("error opening db")
	// }
	// dbQueries := database.New(db)
	// cfg.db = dbQueries

	// ptacCount, err := cfg.db.GetPtacCount(context.Background())
	// if err != nil {
	// 	log.Fatal("error with inital ptac count")
	// }
	// log.Printf("ptac count inital: %d", ptacCount)

	// err = cfg.db.ClearPtacList(context.Background())
	// if err != nil {
	// 	log.Fatal("error clearing db")
	// }

	// ptacCount, err = cfg.db.GetPtacCount(context.Background())
	// if err != nil {
	// 	log.Fatal("error with inital ptac count")
	// }
	// log.Printf("ptac count should be 0: %d", ptacCount)

	// _, err = createPtacList(numRooms)
	// if err != nil {
	// 	log.Fatal("error making ptac list")
	// }

	// // Checks that the right amount of rooms are in database, move to test later
	// ptacCount, err = cfg.db.GetPtacCount(context.Background())
	// if (err != nil) || (ptacCount != numRooms) {
	// 	log.Fatal("error with ptac count")
	// }
	// log.Printf("ptac count end: %d", ptacCount)

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
