package main

import (
	grille "TPPUISSANCE4/Grille"
	"TPPUISSANCE4/player"
	"TPPUISSANCE4/tour"
	"TPPUISSANCE4/variables"
	"fmt"
)

func main() {
	player1 := player.NewPlayer("Joueur 1", "Croix")
	player2 := player.NewPlayer("Joueur 2", "Rond")

	fmt.Printf("\t=== Bienvenue dans le jeu de Puissance 4 ! ===\n")
	fmt.Printf("\t\t=== %s vs %s ===\n", player1.Name, player2.Name)
	fmt.Printf("\n\t\t %s est %s et %s est %s\n", player1.Name, player1.Type, player2.Name, player2.Type)

	fmt.Println("\nVoici la grille de jeu :")
	grille.GetGrille(variables.Thegrille)
	fmt.Println("Le joueur 1 commence !")
	tour.TourJoueur1()
}
