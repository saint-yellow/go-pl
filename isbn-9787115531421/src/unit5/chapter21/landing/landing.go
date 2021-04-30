package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Name string `json:"name"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	locations := []location {
		{Name: "Brabury Landing", Latitude: -4.5895, Longitude:137.4417},
		{Name: "Columbia Memorial Station", Latitude: -14.5684, Longitude: 175.472636},
		{Name: "Challenger Memorial Station", Latitude: -1.9462, Longitude: 354.4734},
	}

	fmt.Println(locations)

	if bytes, e := json.MarshalIndent(locations, "", "	"); e != nil {
		fmt.Println(e)
		os.Exit(1)
	} else {
		fmt.Println(string(bytes))
	}

	

	for _, location := range locations {
		fmt.Println(location)
		if bytes, e := json.MarshalIndent(location, "", "	"); e != nil {
			fmt.Println(e)
			os.Exit(1)
		} else {
			fmt.Println(string(bytes))
		}
	}


}