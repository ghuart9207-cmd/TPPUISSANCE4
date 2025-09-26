package tour

import (
	grille "TPPUISSANCE4/Grille"
	"TPPUISSANCE4/variables"
	"fmt"
)

func TourJoueur1() {
	var colonne string = ""
	var ligne string = ""
	fmt.Println("C'est au tour du Joueur n°1")
	fmt.Println("Choisissez une colonne (1-7) :")
	fmt.Scanln(&colonne)
	switch colonne {
	case "1":
		fmt.Println("Vous avez choisi la colonne 1")
		variables.ColonneChoisie = 1
	case "2":
		fmt.Println("Vous avez choisi la colonne 2")
		variables.ColonneChoisie = 2
	case "3":
		fmt.Println("Vous avez choisi la colonne 3")
		variables.ColonneChoisie = 3
	case "4":
		fmt.Println("Vous avez choisi la colonne 4")
		variables.ColonneChoisie = 4
	case "5":
		fmt.Println("Vous avez choisi la colonne 5")
		variables.ColonneChoisie = 5
	case "6":
		fmt.Println("Vous avez choisi la colonne 6")
		variables.ColonneChoisie = 6
	case "7":
		fmt.Println("Vous avez choisi la colonne 7")
		variables.ColonneChoisie = 7
	default:
		fmt.Println("Colonne invalide, veuillez choisir une colonne entre 1 et 7")
		TourJoueur1()
	}
	fmt.Println("Choisissez une ligne (1-6) :")
	fmt.Scanln(&ligne)
	switch ligne {
	case "1":
		fmt.Println("Vous avez choisi la ligne 1")
		variables.LigneChoisie = 1
	case "2":
		fmt.Println("Vous avez choisi la ligne 2")
		variables.LigneChoisie = 2
	case "3":
		fmt.Println("Vous avez choisi la ligne 3")
		variables.LigneChoisie = 3
	case "4":
		fmt.Println("Vous avez choisi la ligne 4")
		variables.LigneChoisie = 4
	case "5":
		fmt.Println("Vous avez choisi la ligne 5")
		variables.LigneChoisie = 5
	case "6":
		fmt.Println("Vous avez choisi la ligne 6")
		variables.LigneChoisie = 6
	default:
		fmt.Println("Ligne invalide, veuillez choisir une ligne entre 1 et 6")
		TourJoueur1()
	}
	fmt.Printf("Vous avez choisi la colonne %d et la ligne %d\n", variables.ColonneChoisie, variables.LigneChoisie)
	AjouterPionJoueur1(variables.Thegrille)
	grille.GetGrille(variables.Thegrille)
	variables.Tour++
	variables.JoueurActuel = 2
	fmt.Printf("Tour n°%d\n", variables.Tour)
	TourJoueur2()
}
