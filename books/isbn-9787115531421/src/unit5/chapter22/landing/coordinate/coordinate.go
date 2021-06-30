package coordinate

import (
	"encoding/json"
	"fmt"
)

type Coordinate struct {
	Degrees, Minutes, Seconds float64
	Hemisphere rune
}

func New(degrees, minutes, seconds float64, hemisphere rune) Coordinate {
	return Coordinate{degrees, minutes, seconds, hemisphere}
}

func (c Coordinate) Decimal() float64 {
	sign := 1.0
	switch c.Hemisphere {
		case 'S', 'W', 's', 'w':
			sign = -1
	}
	return sign * (c.Degrees + c.Minutes / 60 + c.Seconds / 3600)
}

func (c Coordinate) String() string {
	return fmt.Sprintf(`%v°%v'%.1f" %c`, c.Degrees, c.Minutes, c.Seconds, c.Hemisphere)
}

func (c Coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD float64 `json:"decimal"`
		DMS string `json:"dms"`
		D float64 `json:"degrees"`
		M float64 `json:"minutes"`
		S float64 `json:"seconds"`
		H string `json:"hemisphere"`
	}{
		DD: c.Decimal(),
		DMS: c.String(),
		D: c.Degrees,
		M: c.Minutes,
		S: c.Seconds,
		H: string(c.Hemisphere),
	})
}