package location

import (
	"fmt"
	"isbn-9787115531421/src/unit5/chapter22/landing/coordinate"
)

type Location struct {
	Name string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func New(name string, latitude, longitude coordinate.Coordinate) Location {
	return Location{name, latitude.Decimal(), longitude.Decimal() }
}

func (l Location) String() string {
	return fmt.Sprintf(`%v (%.2f°, %.2f°)`, l.Name, l.Latitude, l.Longitude)
}
