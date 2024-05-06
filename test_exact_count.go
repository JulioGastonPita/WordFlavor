package main

import (
	"fmt"
	"time"

	funcs "github.com/juliogastonpita/WordFlavor/functions"
	//"github.com/ynqa/word2vec"
)

func testexactcount(entidad string, torelated []string) {

	//s1 := "Pedro Sánchez"

	// arreglo de datos a compara
	//toRelated := NewData().Relateds

	start := time.Now()

	for _, item := range torelated {

		exact := funcs.ExactPhrase(entidad, item)
		countWords := funcs.PercentageIncludedWords(entidad, item)

		fmt.Printf("%s        Exact:%t  CountWords:%v\n", item, exact, countWords)

	}

	elapsed := time.Since(start)

	fmt.Printf("Tiempo insumido: %s\n", elapsed)
}
