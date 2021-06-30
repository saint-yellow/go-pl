package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Latitude, Longitude float64
	}

	curiosity := location{Latitude: -4.5895, Longitude: 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}