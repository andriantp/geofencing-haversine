package main

import (
	"log"
	"os"
	"sistem-zonasi/zonasi"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("usage: go run . <user_latitude> <user_longitude>")
		return
	}

	userLat, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Failed to Parse user_latitude [%s]:%v", os.Args[1], err)
	}

	userLong, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("Failed to Parse user_longitude [%s]:%v", os.Args[2], err)
	}

	// example school location
	schoolLocation := zonasi.Location{
		Latitude:  -6.123,
		Longitude: 106.456,
	}

	userLocation := zonasi.Location{
		Latitude:  userLat,
		Longitude: userLong,
	}

	repo := zonasi.NewRepository(2) // x km radius
	result := repo.IsInZone(schoolLocation, userLocation)

	log.Printf("Distance: %f km, In Zone: %t", result.Distance, result.InZone)
}
