package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/crabmustard/crab_maintenance/database"
)

func createPtacList(numRooms int64) error {
	brands := []string{"Amana", "HotPoint", "Trane"}
	models := []string{"12000", "15000", "12000B"}

	// ptacs := []database.Ptac{}

	for room := range numRooms {
		brandIndex := rand.Intn(3)
		modelsIndex := rand.Intn(3)
		year := (2025 - rand.Intn(3))
		month := rand.Intn(12)
		day := rand.Intn(30)
		lastServiceString := fmt.Sprintf("%d-%d-%d", month, day, year)

		pt := database.CreatePtacParams{
			Room:        int64(room + 1),
			Brand:       brands[brandIndex],
			Model:       models[modelsIndex],
			LastService: lastServiceString,
		}
		_, err := cfg.db.CreatePtac(context.Background(), pt)
		if err != nil {
			log.Fatal("unable to add to db")
		}

	}
	return nil

}
