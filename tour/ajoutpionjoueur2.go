package tour

import (
	"TPPUISSANCE4/variables"
)

func AjouterPionJoueur2(grille variables.Grille) {
	col := variables.ColonneChoisie
	lig := variables.LigneChoisie
	switch lig {
	case 1:
		if col >= 1 && col <= 6 {
			grille.Ligne1[col] = " O  |"
		} else if col == 7 {
			grille.Ligne1[7] = " O  "
		}
	case 2:
		if col >= 1 && col <= 6 {
			grille.Ligne2[col] = " O  |"
		} else if col == 7 {
			grille.Ligne2[7] = " O  "
		}
	case 3:
		if col >= 1 && col <= 6 {
			grille.Ligne3[col] = " O  |"
		} else if col == 7 {
			grille.Ligne3[7] = " O  "
		}
	case 4:
		if col >= 1 && col <= 6 {
			grille.Ligne4[col] = " O  |"
		} else if col == 7 {
			grille.Ligne4[7] = " O  "
		}
	case 5:
		if col >= 1 && col <= 6 {
			grille.Ligne5[col] = " O  |"
		} else if col == 7 {
			grille.Ligne5[7] = " O  "
		}
	case 6:
		if col >= 1 && col <= 6 {
			grille.Ligne6[col] = " O  |"
		} else if col == 7 {
			grille.Ligne6[7] = " O  "
		}
	}
}
