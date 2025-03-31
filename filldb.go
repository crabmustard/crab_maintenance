package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/crabmustard/crab_maintenance/database"
)

func createPtacList() {
	roomNumbers := [80]int{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
		117, 118, 119, 120, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212,
		213, 214, 215, 216, 217, 218, 219, 220, 301, 302, 303, 304, 305, 306, 307, 308,
		309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 401, 402, 403, 404,
		405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420}
	brands := []string{"amana", "hotpoint", "trane", "distinctions"}
	models := []string{"12000", "15000"}
	// ptacs := []database.Ptac{}

	for _, room := range roomNumbers {
		brandIndex := rand.Intn(3)
		modelsIndex := rand.Intn(2)
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
		_, err := cfg.db.CreatePtac(context.Background(), pt)
		if err != nil {
			log.Fatal("unable to add to db")
		}
		// ptacs = append(ptacs, ptdb)

	}

}
