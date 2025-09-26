package grille

import (
	"TPPUISSANCE4/variables"

	"fmt"
)

func GetGrille(grilleactive variables.Grille) {
	fmt.Println(grilleactive.Ligne0)
	fmt.Println(grilleactive.Ligne1)
	fmt.Println(grilleactive.Ligne2)
	fmt.Println(grilleactive.Ligne3)
	fmt.Println(grilleactive.Ligne4)
	fmt.Println(grilleactive.Ligne5)
	fmt.Println(grilleactive.Ligne6)
}
