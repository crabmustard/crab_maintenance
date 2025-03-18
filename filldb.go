package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/crabmustard/crab_maintenance/database"
)

func createPtacList(numRooms int64) ([]database.Ptac, error) {
	brands := []string{"Amana", "HotPoint", "Trane"}
	models := []string{"12000", "15000", "12000B"}

	ptacs := []database.Ptac{}

	for room := range int(numRooms) {
		brandIndex := rand.Intn(3)
		modelsIndex := rand.Intn(3)
		year := (2024 - rand.Intn(3))
		month := rand.Intn(12) + 1
		day := rand.Intn(30) + 1
		lastServiceString := fmt.Sprintf("%d-%02d-%02d", year, month, day)

		pt := database.CreatePtacParams{
			Room:        strconv.Itoa(room),
			Brand:       brands[brandIndex],
			Model:       models[modelsIndex],
			LastService: lastServiceString,
		}
		ptdb, err := cfg.db.CreatePtac(context.Background(), pt)
		if err != nil {
			log.Fatal("unable to add to db")
		}
		ptacs = append(ptacs, ptdb)

	}
	return ptacs, nil

}
