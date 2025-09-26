package grille

import "fmt"

type Grille struct {
	Ligne1 []string
	Ligne2 []string
	Ligne3 []string
	Ligne4 []string
	Ligne5 []string
	Ligne6 []string
	Ligne0 string
}

var grille Grille = Grille{Ligne1: []string{"1.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne2: []string{"2.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne3: []string{"3.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne4: []string{"4.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne5: []string{"5.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne6: []string{"6.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne0: "      1 |  2  |  3  |  4  |  5  |  6  |  7  "}

func GetGrille() Grille {
	fmt.Println(grille.Ligne0)
	fmt.Println(grille.Ligne1)
	fmt.Println(grille.Ligne2)
	fmt.Println(grille.Ligne3)
	fmt.Println(grille.Ligne4)
	fmt.Println(grille.Ligne5)
	fmt.Println(grille.Ligne6)

	return grille
}
