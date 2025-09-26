package tour

import (
	"TPPUISSANCE4/variables"
	"fmt"
)

func VérifierPartieNulle(grille variables.Grille) {
	nombreDeCasesRemplies := 0
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 7; j++ {
			if grille.Ligne1[j] != "    |" {
				nombreDeCasesRemplies++
			}
		}
	}
	if variables.Gagnant > 0 {
		fmt.Printf("\t\t====Le joueur %d a gagné !====\n", variables.Gagnant)
	} else if nombreDeCasesRemplies == 42 {
		fmt.Println("\t\t===La partie est nulle !===")
		variables.PartieNulle = true
	}
}

func PUISSANCE4() {
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 7; j++ {
			if grille.Ligne1[i][j] != "    |" {

			}
		}
	}
}
