package dictionary

import (
	"example/dictionary"
	"fmt"
)

func main() {
	// Utilisation de la classe Dictionnaire
	dict := dictionary.NewDictionnaire()

	dict.Ajouter("estiam", "ecole")
	dict.Ajouter("ZIAD", "etudiant")

	fmt.Println("Get 'estiam':", dict.Get("estiam"))

	dict.Lister()

	dict.Supprimer("ZIAD")
	dict.Lister()
}
