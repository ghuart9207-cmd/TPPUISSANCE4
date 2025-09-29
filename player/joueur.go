package player

import "example.com/TPPUISSANCE4/grille"

type Joueur struct {
	Nom     string
	Symbole grille.Cell
}

// NewJoueur crée un joueur avec son nom et son symbole
func NewJoueur(nom string, symbole grille.Cell) Joueur {
	return Joueur{
		Nom:     nom,
		Symbole: symbole,
	}
}
