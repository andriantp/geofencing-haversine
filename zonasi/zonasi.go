package zonasi

import (
	"math"
)

type (
	Location struct {
		Latitude  float64
		Longitude float64
	}

	Result struct {
		Distance float64
		InZone   bool
	}
)
type repo struct {
	zoneRadius float64
}

type RepositoryI interface {
	IsInZone(schoolLocation, userLocation Location) Result
}

func NewRepository(zoneRadius float64) RepositoryI {
	zoneRadiusMeters := zoneRadius * 1000 // convert km to meters
	return &repo{
		zoneRadius: zoneRadiusMeters,
	}
}

func (r *repo) IsInZone(schoolLocation, userLocation Location) Result {
	distance := r.distanceMeters(schoolLocation, userLocation)

	if distance <= r.zoneRadius {
		return Result{
			Distance: distance / 1000, // convert meters to km
			InZone:   true,
		}
	}

	return Result{
		Distance: distance / 1000, // convert meters to km
		InZone:   false,
	}
}

func (r *repo) distanceMeters(schoolLocation, userLocation Location) float64 {
	const R = 6371000.0

	lat1Rad := schoolLocation.Latitude * math.Pi / 180
	lon1Rad := schoolLocation.Longitude * math.Pi / 180

	lat2Rad := userLocation.Latitude * math.Pi / 180
	lon2Rad := userLocation.Longitude * math.Pi / 180

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(
		math.Sqrt(a),
		math.Sqrt(1-a),
	)

	return R * c
}
