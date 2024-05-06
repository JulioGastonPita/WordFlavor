package main

import (
	"fmt"
	"time"

	funcs "github.com/juliogastonpita/WordFlavor/functions"
	//"github.com/ynqa/word2vec"
)

func testcompareall(entidad string, torelated []string) {

	// s1 := "Pedro Sánchez"

	// // List of strings to compare with
	// toRelated := NewData().Relateds

	start := time.Now()

	for _, item := range torelated {

		booly := funcs.CompareAll(entidad, item, 0.6, 10)

		fmt.Printf("%s        CompareAll: %t\n", item, booly)

	}

	elapsed := time.Since(start)

	fmt.Printf("Tiempo insumido: %s\n", elapsed)
}
