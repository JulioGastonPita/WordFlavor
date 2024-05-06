package main

import (
	"fmt"
	"time"

	texttheater "github.com/texttheater/golang-levenshtein/levenshtein"
	//"github.com/ynqa/word2vec"
)

func testlevenshtein(entidad string, torelated []string) {

	// s1 := "Pedro Sánchez"

	// // List of strings to compare with
	// toRelated := NewData().Relateds

	start := time.Now()

	for _, item := range torelated {

		distance3 := texttheater.DistanceForStrings([]rune(entidad), []rune(item), texttheater.DefaultOptions)

		fmt.Printf("%s         Distancia: %d\n", item, distance3)

	}

	elapsed := time.Since(start)

	fmt.Printf("Tiempo insumido: %s\n", elapsed)
}
