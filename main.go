package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	filePath := flag.String("f", "resources/config.json", "file path")
	flag.Parse()

	config, err := loadConfig(*filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Successfully loaded the config from '%s'!\n", *filePath)
	log.Printf("%d sounds(s) are defined in the config.\n", len(config.Sounds))
	log.Printf("%d pattern(s) are defined in the config.\n", len(config.Pattern))

	words := generateWords(config, DefaultRandSource{})

	for _, word := range words {
		fmt.Println(word)
	}
}
