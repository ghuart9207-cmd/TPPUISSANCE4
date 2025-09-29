package main

import (
	"fmt"

	"example.com/TPPUISSANCE4/grille"
	"example.com/TPPUISSANCE4/player"
	"example.com/TPPUISSANCE4/tour"
)

func main() {
	board := grille.NewBoard()

	joueurX := player.NewJoueur("Joueur 1", grille.X)
	joueurO := player.NewJoueur("Joueur 2", grille.O)

	fmt.Println("=== Bienvenue dans Puissance 4 ===")
	fmt.Println(board)

	current := joueurX

	for {
		tour.PlayTurn(board, current)
		fmt.Println(board)

		// alterner joueurs
		if current.Symbole == grille.X {
			current = joueurO
		} else {
			current = joueurX
		}
	}
}
