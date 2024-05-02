package main

import (
	"fmt"
	"time"

	funcs "github.com/juliogastonpita/WordFlavor/functions"

	"github.com/agnivade/levenshtein"

	texttheater "github.com/texttheater/golang-levenshtein/levenshtein"
)

func main() {
	s1 := "Pedro Sánchez"

	// List of strings to compare with
	toRelated := []string{"Pedro Sanchez",
		"Gobierno de Pedro Sánchez",
		"Esposa de Pedro Sánchez",
		"pantuflas para los pies de pedro sanchez",
		"Presidente Pedro Sánchez",
		"Spanish Prime Minister Pedro Sanchez",
		"Criticas contra Pedro Sanchez",
		"Pedro Sánchez Pérez-Castejón",
		"Pedro Sánchez, Kajsa Ollongren y mandatarios de Europa",
		"Partido de Pedro Sánchez",
		"estrategia de Pedro Sánchez",
		"allies of bruised spanish pm pedro sanchez",
		"Pedro Sánchez Fernández",
		"Педро Санчес Перес-Кастехон (Pedro Sanchez Perez-Castejon)",
		"桑切斯 (Pedro Sánchez)",
		"Wife of Pedro Sanchez",
		"Spain's Prime Minister Pedro Sanchez",
		"otros pedro sanchez en España",
		"Педро Санчес (Pedro Sánchez)",
		"Partido Político de Pedro Sánchez",
		"Pedro Sanchez's wife, Begona Gomez",
		"Pedro Sanchez, Begona Gomez",
		"Pedro Sanchez and his wife",
		"西班牙首相佩德罗·桑切斯 (Pedro Sanchez)",
		"桑杰士 (Pedro Sánchez)",
		"总理桑杰士 (Pedro Sanchez)",
		"Mujer de Pedro Sánchez",
		"Spanish Leader Pedro Sanchez",
		"Pedro Sanchez's wife",
		"Supporters of Spain’s Prime Minister Pedro Sanchez",
		"Spain PM Pedro Sanchez",
		"Dimisión de Pedro Sánchez",
		"Longevidad del gobierno de Pedro Sánchez",
		"Prime Minister Pedro Sanchez",
		"Spain’s Prime Minister Pedro Sanchez",
		"Supporters of Spain's Prime Minister Pedro Sanchez",
		"Pedro Sanchez und seine Frau",
		"allegations against Pedro Sánchez's wife",
		"Hija del juez que aceptó querella contra la mujer de Pedro Sánchez",
		"Spanish PM Pedro Sanchez",
	}

	start := time.Now()

	for _, item := range toRelated {

		exact := funcs.ExactPhrase(s1, item)

		countWords := funcs.CountIncludedWords(s1, item)

		distance1 := funcs.LevenshteinDistance(s1, item)

		distance2 := levenshtein.ComputeDistance(s1, item)

		distance3 := texttheater.DistanceForStrings([]rune(s1), []rune(item), texttheater.DefaultOptions)

		fmt.Printf("%s %t %d %d %d %d\n", item, exact, countWords, distance1, distance2, distance3)

	}

	elapsed := time.Since(start)

	fmt.Printf("Time taken: %s\n", elapsed)
}
