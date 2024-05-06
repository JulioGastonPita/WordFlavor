package functions

import (
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/ynqa/wego/pkg/embedding"
	"github.com/ynqa/wego/pkg/search"

	"gonum.org/v1/gonum/floats"
)

// LevenshteinDistance calcula la distancia de Levenshtein entre dos cadenas de texto.
// La distancia de Levenshtein es una medida de la diferencia entre dos cadenas de texto,
// representada por el número mínimo de operaciones requeridas para transformar una cadena en la otra.
// Esta función utiliza el algoritmo de programación dinámica para calcular la distancia de Levenshtein.
// Recibe dos parámetros de tipo string, s1 y s2, que son las cadenas de texto a comparar.
// Devuelve un entero que representa la distancia de Levenshtein entre las dos cadenas.
func LevenshteinDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	// Create a 2D matrix to store the distances
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Initialize the first row and column
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Calculate the distances
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	// Return the distance
	return dp[m][n]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

// CountIncludedWords cuenta la cantidad de palabras incluidas en s2 que también están en s1.
// Recibe dos cadenas de texto, s1 y s2, y devuelve un entero que representa la cantidad de palabras incluidas.

func CountIncludedWords(s1, s2 string, removeaccents bool) int {
	count := 0
	words := strings.Fields(s1)
	for _, word := range words {
		if removeaccents && strings.Contains(s2, removeAccents(word)) {
			count++
		}

		if !removeaccents && strings.Contains(s2, word) {
			count++
		}

	}
	return count
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// devuelve un valor entre 0 y 1 que representa el porcentaje de palabras incluidas en s2 que también están en s1.
func PercentageIncludedWords(s1, s2 string) float64 {
	totalWords := len(strings.Fields(s2))
	includedWords := CountIncludedWords(s1, s2, true)
	percentage := float64(includedWords) / float64(totalWords)
	return percentage
}

// ExactPhrase compara dos cadenas de texto y devuelve true si s1 es una frase exacta en s2.
func ExactPhrase(s1, s2 string) bool {
	regex := regexp.MustCompile(`\b` + regexp.QuoteMeta(s1) + `\b`)
	return regex.MatchString(s2)
}

func CompareAll(entidad, comparador string, sensPercentageIncludeWords float64, sensLevenshteinDistance int) bool {

	// si es exacto la frase devuelve truen
	if ExactPhrase(entidad, comparador) {
		return true
	}

	// si el porcentaje de palabras incluidas es mayor a 0.5 devuelve true
	if PercentageIncludedWords(entidad, comparador) >= sensPercentageIncludeWords {
		return true
	}

	// si la distancia es meno a 10 devuelve true
	if LevenshteinDistance(entidad, comparador) <= sensLevenshteinDistance {
		return true
	}

	return false

}

// VectorSimilarity calcula la similitud de coseno entre dos vectores de palabras.
// Recibe dos cadenas de texto s1 y s2 que representan las palabras a comparar.
// También recibe la ruta al archivo de vectores pathtofilevectors.
// Devuelve un valor float64 que representa la similitud de coseno entre los vectores.
func VectorSimilarity(s1, s2 string, pathtofilevectors string) float64 {

	input, err := os.Open(pathtofilevectors)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	embs, err := embedding.Load(input)
	if err != nil {
		log.Fatal(err)
	}
	searcher, err := search.New(embs...)
	if err != nil {
		log.Fatal(err)
	}
	//	neighbors, err := searcher.SearchInternal("tugurios", 10)

	vec1, _ := searcher.Items.Find(s1)
	vec2, _ := searcher.Items.Find(s2)
	if err != nil {
		log.Fatal(err)
	}

	// Calcular el producto punto de los dos vectores
	dotProduct := floats.Dot(vec1.Vector, vec2.Vector)

	// Calcular la similitud del coseno
	cosineSimilarity := dotProduct / (vec1.Norm * vec2.Norm)

	// Imprimir la similitud del coseno
	//	fmt.Println(cosineSimilarity)

	//	neighbors.Describe()

	return cosineSimilarity

}
