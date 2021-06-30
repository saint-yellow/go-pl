package main

import (
	"encoding/json"
	"fmt"
	"isbn-9787115531421/src/unit5/chapter22/landing/coordinate"
	"isbn-9787115531421/src/unit5/chapter24/marshal/location"
	"os"
)

func main() {
	elysium := location.New("Elysium Planitia", coordinate.New(4, 30, 0.0, 'N'), coordinate.New(135, 54, 0.0, 'E'))
	bytes, err := json.MarshalIndent(elysium, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
	
}