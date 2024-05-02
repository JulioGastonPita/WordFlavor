package main

import (
	"fmt"
	"time"

	"github.com/JulioGastonPita/wordflavor/funcs"
)

func main() {
	s1 := "Oesía"

	// List of strings to compare with
	toRelated := []string{"Grupo Oesía",
		"Poesía",
		"Teatro y Poesía",
		"Poesía sentimental",
		"Poesía ingenua",
		"sobre poesia ingenua y poesia sentimental",
		"poesia afroantillana y negrista",
		"Oesia",
		"Cipherbit-Grupo Oesia",
		"Torre Oesia",
		"Tecnobit-Grupo Oesia",
		"Inster-Grupo Oesia",
		"oesia networks",
		"UAV Navigation-Grupo Oesia",
		"Director financiero de Grupo Oesía",
		"Presidencia de Grupo Oesía",
		"Instel-Grupo Oesia",
		"Poesía y elocuencia de las tribus de América y otros textos",
		"Oesia Group",
		"Libro de poesía",
		"Inster-Oesia Group",
		"poesia moderna",
		"Oesia Networks, Tecnobit-Grupo Oesia, Cipherbit-Grupo Oesia, UAV Navigation-Grupo Oesia, Inster-Grupo Oesia",
		"Oesia Networks Colombia S.A.S",
		"poesía clasista",
		"Libro 'Estudio e índices de poesía y poética (1988-1999)'",
		"poesia espanola",
		"Tecnobit-Oesia Group",
		"estudio e indices de poesia y poetica (1988-1999) de isaac canton",
		"poesia del marketing: the party",
		"Tecnobit - Grupo Oesia",
		"inster-grupo oesia, airbus helicopters espana, tecnobit-grupo oesia, anzen engineering",
		"Premio Reina Sofía de Poesía Iberoamericana",
		"poesia social",
		"Proyectos de Defensa de Grupo Oesia",
		"poesia vanguardista",
		"Libro de poesía de Jose Hierro",
		"poesia bucolica",
		"poesia de senectud",
		"Oesia Networks S.L.",
		"poesia y elocuencia de las tribus de america",
		"Oesia Networks SL",
		"poesia epica",
		"leccion de poesia",
		"escribir poesía",
		"reflexiones criticas sobre la poesia y sobre la pintura de jean-baptiste du bos",
		"Financial Director of Grupo Oesia",
		"eloquencia y poesía",
		"Grupa Oesia",
		"xxv premio reina sofia de poesia iberoamericana",
		"Relevancia de Cipherbit-Grupo Oesia",
		"poesia dramatica",
		"la poesia de jose kozer: de la recta a las cajas chinas",
		"poesia lirica",
		"programa de RSC 'Oesia Sostenible'",
		"Luis Furnells (Presidente ejecutivo de Grupo Oesia)",
		"centro tecnológico de Valdepeñas de Tecnobit-Grupo Oesia",
		"uav navigation-grupo oesia’s attitude and heading reference system (ahrs)",
		"UAV Navigation-Oesia Group",
		"Poesía del Medioevo",
		"poesía y poetas",
		"Oesia Networks Sociedad Limitada",
		"1147 specimens of oesia",
		"Luis Furnells (Presidente Ejecutivo de Oesia Group)",
		"ccoo - oesia",
		"Executive President of Grupo Oesia",
		"Emilio Varela (Tecnobit-Grupo Oesia)",
		"Sener Aeroespacial, GMV, Tecnobit-Grupo Oesia",
	}

	start := time.Now()

	for _, item := range toRelated {
		distance := funcs.LevenshteinDistance(s1, item)
		fmt.Printf("Levenshtein distance between %s and %s is %d\n", s1, item, distance)
	}

	elapsed := time.Since(start)

	fmt.Printf("Time taken: %s\n", elapsed)
}
