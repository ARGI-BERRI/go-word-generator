package main

import (
	"fmt"
)

func main() {
	config, err := loadConfig("resources/config.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("INFO: Successfully loaded the config!")
	fmt.Printf("INFO: %d sounds(s) are defined in the config.\n", len(config.Sounds))
	fmt.Printf("INFO: %d pattern(s) are defined in the config.\n", len(config.Pattern))
}
