package tour

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/TPPUISSANCE4/etat"
	"example.com/TPPUISSANCE4/grille"
	"example.com/TPPUISSANCE4/player"
	"example.com/TPPUISSANCE4/variables"
)

func PlayTurn(b *grille.Board, j player.Joueur) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s (%c), choisis une colonne (0-%d) : ",
			j.Nom, j.Symbole, variables.Cols-1)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		col, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("⚠️ Valeur non acceptée, entre un nombre entre 0 et 6.")
			continue
		}

		if err := b.Drop(col, j.Symbole); err != nil {
			fmt.Println("⚠️", err)
			continue
		}

		// Vérifier l'état du jeu après que le joueur ait joué
		gameOver, status := etat.IsGameOver(b, j.Symbole)
		if gameOver {
			if status == "win" {
				fmt.Printf("🎉 %s (%c) a gagné !\n", j.Nom, j.Symbole)
			} else if status == "draw" {
				fmt.Println("🤝 Match nul !")
			}
		}

		break
	}
}
