package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

const (
	largeurEcran   = 1000
	hauteurEcran   = 600
	tailleGrille   = 9
	tailleCellule  = 66
	boutonLargeur  = 120
	boutonHauteur  = 40
	boutonX        = largeurEcran - boutonLargeur - 20
	boutonY        = 10
	resultatX      = 10
	resultatY      = 70
	resultatTaille = 24
)

var (
	grille          [tailleGrille][tailleGrille]int
	rangeeSel       = -1
	colonneSel      = -1
	verificationEnCours = false
)

func main() {
	rl.InitWindow(largeurEcran, hauteurEcran, "Exemple de Sudoku")
	rl.SetTargetFPS(60)

	// Initialiser le générateur de nombres aléatoires avec une graine basée sur le temps
	rand.Seed(time.Now().UnixNano())

	// Générer une grille Sudoku valide avec quelques cases vides
	genererSudoku()

	for !rl.WindowShouldClose() {
		gererSaisie()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			sourisX := rl.GetMouseX()
			sourisY := rl.GetMouseY()

			// Vérifier si le bouton de vérification a été cliqué
			if sourisX >= boutonX && sourisX <= boutonX+boutonLargeur && sourisY >= boutonY && sourisY <= boutonY+boutonHauteur {
				verificationEnCours = true
				resultat := verifierGrille()
				fmt.Println("La grille est correcte:", resultat)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Dessiner la grille du sudoku
		dessinerGrille()
		dessinerNombres()

		// Dessiner le bouton de vérification
		rl.DrawRectangle(int32(boutonX), int32(boutonY), int32(boutonLargeur), int32(boutonHauteur), rl.RayWhite)
		rl.DrawRectangleLines(int32(boutonX), int32(boutonY), int32(boutonLargeur), int32(boutonHauteur), rl.Black)
		rl.DrawText("Vérifier", int32(boutonX+10), int32(boutonY+10), 20, rl.Black)

		// Afficher un message de résultat de vérification
		if verificationEnCours {
			message := "La grille est correcte!"
			if !verifierGrille() {
				message = "La grille est incorrecte."
			}
			rl.DrawText(message, int32(resultatX), int32(resultatY), int32(resultatTaille), rl.Black)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func verifierGrille() bool {
	for rangee := 0; rangee < tailleGrille; rangee++ {
		for colonne := 0; colonne < tailleGrille; colonne++ {
			chiffre := grille[rangee][colonne]
			if chiffre != 0 {
				// Vérification de la ligne
				for i := 0; i < tailleGrille; i++ {
					if i != colonne && grille[rangee][i] == chiffre {
						return false
					}
				}

				// Vérification de la colonne
				for i := 0; i < tailleGrille; i++ {
					if i != rangee && grille[i][colonne] == chiffre {
						return false
					}
				}

				// Vérification de la région 3x3
				rangeeDebut, colonneDebut := rangee-rangee%3, colonne-colonne%3
				for i := rangeeDebut; i < rangeeDebut+3; i++ {
					for j := colonneDebut; j < colonneDebut+3; j++ {
						if i != rangee && j != colonne && grille[i][j] == chiffre {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func genererSudoku() {
	// Effacer la grille
	for rangee := 0; rangee < tailleGrille; rangee++ {
		for colonne := 0; colonne < tailleGrille; colonne++ {
			grille[rangee][colonne] = 0
		}
	}

	// Remplir la grille en respectant les règles du Sudoku
	resoudreSudoku()

	// Limiter le nombre de chiffres générés aléatoirement
	restantes := 81 - compterNombresInitiaux()
	for restantes > 0 {
		rangee := rand.Intn(tailleGrille)
		colonne := rand.Intn(tailleGrille)

		if grille[rangee][colonne] == 0 {
			chiffre := rand.Intn(9) + 1
			if estSecuritaire(rangee, colonne, chiffre) {
				grille[rangee][colonne] = chiffre
				restantes--
			}
		}
	}

	// Supprimer certaines valeurs pour créer des cases vides
	casesVides := 45 // Nombre de cases vides (ajustez-le selon vos préférences)
	for casesVides > 0 {
		rangee := rand.Intn(tailleGrille)
		colonne := rand.Intn(tailleGrille)

		if grille[rangee][colonne] != 0 {
			grille[rangee][colonne] = 0
			casesVides--
		}
	}
}

func resoudreSudoku() bool {
	rangeeVide, colonneVide := trouverCaseVide() // Ignorer la deuxième valeur

	if rangeeVide == -1 && colonneVide == -1 {
		return true // La grille est résolue
	}

	for chiffre := 1; chiffre <= 9; chiffre++ {
		if estSecuritaire(rangeeVide, colonneVide, chiffre) {
			grille[rangeeVide][colonneVide] = chiffre

			if resoudreSudoku() {
				return true
			}

			grille[rangeeVide][colonneVide] = 0 // Annuler la tentative
		}
	}

	return false // Aucune solution trouvée pour cette configuration
}

func estSecuritaire(rangee, colonne, chiffre int) bool {
	// Vérifier la ligne
	for i := 0; i < tailleGrille; i++ {
		if grille[rangee][i] == chiffre {
			return false
		}
	}

	// Vérifier la colonne
	for i := 0; i < tailleGrille; i++ {
		if grille[i][colonne] == chiffre {
			return false
		}
	}

	// Vérifier la région 3x3
	rangeeDebut, colonneDebut := rangee-rangee%3, colonne-colonne%3
	for i := rangeeDebut; i < rangeeDebut+3; i++ {
		for j := colonneDebut; j < colonneDebut+3; j++ {
			if grille[i][j] == chiffre {
				return false
			}
		}
	}

	return true
}

func trouverCaseVide() (int, int) {
	for rangee := 0; rangee < tailleGrille; rangee++ {
		for colonne := 0; colonne < tailleGrille; colonne++ {
			if grille[rangee][colonne] == 0 {
				return rangee, colonne
			}
		}
	}
	return -1, -1 // Aucune case vide trouvée
}

func compterNombresInitiaux() int {
	compte := 0
	for rangee := 0; rangee < tailleGrille; rangee++ {
		for colonne := 0; colonne < tailleGrille; colonne++ {
			if grille[rangee][colonne] != 0 {
				compte++
			}
		}
	}
	return compte
}

func gererSaisie() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		sourisX := float32(rl.GetMouseX())
		sourisY := float32(rl.GetMouseY())

		rangeeSel = int(sourisY / tailleCellule)
		colonneSel = int(sourisX / tailleCellule)
	}

	for touche := rl.KeyOne; touche <= rl.KeyNine; touche++ {
		if rl.IsKeyPressed(int32(touche)) && rangeeSel != -1 && colonneSel != -1 {
			if grille[rangeeSel][colonneSel] == 0 {
				grille[rangeeSel][colonneSel] = int(touche - rl.KeyOne + 1)
			}
		}
	}
}

func dessinerGrille() {
	for i := 0; i < tailleGrille; i++ {
		// Dessiner les lignes horizontales et verticales
		rl.DrawLine(int32(float32(i)*tailleCellule), 0, int32(float32(i)*tailleCellule), hauteurEcran, rl.Black)
		rl.DrawLine(0, int32(float32(i)*tailleCellule), largeurEcran, int32(float32(i)*tailleCellule), rl.Black)

		// Dessiner des lignes épaisses pour délimiter les régions 3x3
		if i%3 == 0 && i != 0 {
			rl.DrawLine(int32(float32(i)*tailleCellule), 0, int32(float32(i)*tailleCellule), hauteurEcran, rl.Black)
			rl.DrawLine(0, int32(float32(i)*tailleCellule), largeurEcran, int32(float32(i)*tailleCellule), rl.Black)
		}
	}

	// Dessiner des lignes épaisses autour des régions
	for i := 0; i <= 3; i++ {
		rl.DrawLine(0, int32(i*3*tailleCellule)+1, largeurEcran, int32(i*3*tailleCellule), rl.Black)
		rl.DrawLine(0, int32(i*3*tailleCellule)-1, largeurEcran, int32(i*3*tailleCellule), rl.Black)


		rl.DrawLine(int32(i*3*tailleCellule)+1, 0, int32(i*3*tailleCellule), hauteurEcran, rl.Black)
		rl.DrawLine(int32(i*3*tailleCellule)-1, 0, int32(i*3*tailleCellule), hauteurEcran, rl.Black)
	}
}


func dessinerNombres() {
	for rangee := 0; rangee < tailleGrille; rangee++ {
		for colonne := 0; colonne < tailleGrille; colonne++ {
			chiffre := grille[rangee][colonne]
			if chiffre != 0 {
				// Dessiner le chiffre au centre de la cellule
				x := int32(float32(colonne)*tailleCellule + tailleCellule/2 - 10)
				y := int32(float32(rangee)*tailleCellule + tailleCellule/2 - 10)
				rl.DrawText(fmt.Sprintf("%d", chiffre), x, y, 20, rl.Black)
			}

			// Dessiner la cellule sélectionnée en surbrillance
			if rangee == rangeeSel && colonne == colonneSel {
				rl.DrawRectangleLines(int32(float32(colonne)*tailleCellule), int32(float32(rangee)*tailleCellule), int32(tailleCellule), int32(tailleCellule), rl.Red)
			}
		}
	}
}
