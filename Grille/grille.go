package grille

import "fmt"

type Grille struct {
	Ligne1 string
	Ligne2 string
	Ligne3 string
	Ligne4 string
	Ligne5 string
	Ligne6 string
	Ligne7 string
}

var grille Grille = Grille{Ligne1: "    |    |    |    |    |    |    ", Ligne2: "    |    |    |    |    |    |    ", Ligne3: "    |    |    |    |    |    |    ", Ligne4: "    |    |    |    |    |    |    ", Ligne5: "    |    |    |    |    |    |    ", Ligne6: "    |    |    |    |    |    |    "}

func GetGrille() Grille {
	fmt.Println(grille.Ligne1)
	fmt.Println(grille.Ligne2)
	fmt.Println(grille.Ligne3)
	fmt.Println(grille.Ligne4)
	fmt.Println(grille.Ligne5)
	fmt.Println(grille.Ligne6)

	return grille
}
