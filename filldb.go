package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/crabmustard/crab_maintenance/database"
)

func createPtacList(numRooms int64) error {
	brands := []string{"Amana", "HotPoint", "Trane"}
	models := []string{"12000", "15000", "12000B"}

	// ptacs := []database.Ptac{}

	for room := range numRooms {
		brandIndex := rand.Intn(3)
		modelsIndex := rand.Intn(3)
		pt := database.CreatePtacParams{
			Room:        int64(room),
			Brand:       brands[brandIndex],
			Model:       models[modelsIndex],
			LastService: time.Now().Unix(),
		}
		_, err := cfg.db.CreatePtac(context.Background(), pt)
		if err != nil {
			log.Fatal("unable to add to db")
		}

	}
	return nil

}
