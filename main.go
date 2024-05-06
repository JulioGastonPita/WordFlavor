package main

import (
	"fmt"

	funcs "github.com/juliogastonpita/WordFlavor/functions"
)

func main() {

	// entidad para comparar
	s1 := "Pedro Sánchez"

	// lista de entidades que se compararan, ver data.go
	toRelated := NewData().Relateds

	//	Funcion que devuelve si la frase es exacta y la cantidad de palabras exactas
	testexactcount(s1, toRelated)

	//	funcion que devuelve la distancia de levenshtein
	testlevenshtein(s1, toRelated)

	// Hace una comparacion con todas las funciones
	testcompareall(s1, toRelated)

	// Recordar descomprimir words_vectors.txt
	fmt.Println("VectorSimilarity")
	fmt.Println(funcs.VectorSimilarity("tugurio", "noveno", "words_vector9.txt"))

}
