package location

import "isbn-9787115531421/src/unit5/chapter22/landing/coordinate"

type Location struct {
	Name string `json:"name"`
	Latitude coordinate.Coordinate `json:"latitude"`
	Longitude coordinate.Coordinate `json:"longitude"`
}

func New(name string, latitude, longitude coordinate.Coordinate) Location {
	return Location {
		Name: name,
		Latitude: latitude,
		Longitude: longitude,
	}
}