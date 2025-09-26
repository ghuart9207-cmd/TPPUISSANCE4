package variables

type Grille struct {
	Ligne1 []string
	Ligne2 []string
	Ligne3 []string
	Ligne4 []string
	Ligne5 []string
	Ligne6 []string
	Ligne0 string
}

var Tour int = 1

var JoueurActuel int = 1
var Gagnant int = 0
var PartieNulle bool = false

var ColonneRemplie bool = false
var ColonneValide bool = false
var ligneValide bool = false
var ColonneChoisie int = -1
var LigneChoisie int = -1

var RecommencerPartie bool = false
var QuitterPartie bool = false

var Thegrille Grille = Grille{Ligne1: []string{"1.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne2: []string{"2.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne3: []string{"3.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne4: []string{"4.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne5: []string{"5.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne6: []string{"6.", "    |", "    |", "    |", "    |", "    |", "    |", "    "}, Ligne0: "      1 |  2  |  3  |  4  |  5  |  6  |  7  "}
