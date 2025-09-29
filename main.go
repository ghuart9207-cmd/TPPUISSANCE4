package main

import (
	"fmt"

	"example.com/TPPUISSANCE4/etat"
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

		gameOver, status := etat.IsGameOver(board, current.Symbole)
		if gameOver {
			if status == "win" {
				fmt.Printf("🎉 Jeu terminé ! %s (%c) a gagné !\n", current.Nom, current.Symbole)
			} else if status == "draw" {
				fmt.Println("🤝 Jeu terminé ! Match nul !")
			}
			break
		}

		if current.Symbole == grille.X {
			current = joueurO
		} else {
			current = joueurX
		}
	}
}
