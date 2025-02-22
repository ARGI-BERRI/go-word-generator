package main

import (
	"log"
	"math/rand"
)

type RandSource interface {
	get(n int) int
}

type DefaultRandSource struct{}

func (d DefaultRandSource) get(n int) int {
	return rand.Intn(n)
}

func choiceSound(sounds []string, r RandSource) string {
	if len(sounds) == 0 {
		return ""
	}

	return sounds[r.get(len(sounds))]
}

func generateWords(config Config, randSource RandSource) []string {
	sounds := config.Sounds
	patterns := config.Pattern

	var generatedWords []string

	for _, pattern := range patterns {
		log.Printf("Generating pattern: %s\n", pattern.Label)

		var word string

		for _, syllable := range pattern.Syllable {
			if len(sounds[syllable]) == 0 {
				log.Fatalf("'%s' is not defined in Sounds\n", syllable)
			}

			word += choiceSound(sounds[syllable], randSource)
		}

		if len(word) > 0 {
			generatedWords = append(generatedWords, word)
		}
	}

	return generatedWords
}
