package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	largeurÉcran      = 800
	hauteurÉcran      = 600
	pasVolume         = 1.0 // Ajustez le pas de volume selon vos besoins
	largeurBarreVol   = 200 // Largeur de la barre de volume
	hauteurBarreVol   = 20  // Hauteur de la barre de volume
	margeBarreVol     = 5   // Marge autour de la barre de volume
	tailleIcône       = 32  // Taille des icônes de volume
	largeurBtnÀPropos = 150 // Largeur du bouton "À Propos"
	hauteurBtnÀPropos = 40  // Hauteur du bouton "À Propos"
)

var (
	params struct {
		volume           int32
		positionBarreVol int32
		volumeHaut       rl.Texture2D
		volumeMuet       rl.Texture2D
		àProposBtn       rl.Rectangle
		afficherÀPropos  bool
		texteÀPropos     string
		policeÀPropos    rl.Font
	}
)

func loadImages() {
	params.volumeHaut = rl.LoadTexture("C:\\main\\volumeup.png") // Remplacez par le chemin de votre image pour le volume haut
	params.volumeMuet = rl.LoadTexture("C:\\main\\mute.png")     // Remplacez par le chemin de votre image pour le volume muet
}

func main() {
	rl.InitWindow(largeurÉcran, hauteurÉcran, "Contrôle du Son")

	// Initialisation du périphérique audio
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	loadImages()

	// Charger le son MP3 comme un son normal (pas un flux musical)
	son := rl.LoadSound("C:\\main\\smurfcat.mp3")

	// Jouer le son
	rl.PlaySound(son)

	// Définir le rectangle du bouton "À Propos"
	params.àProposBtn = rl.NewRectangle(float32(largeurÉcran/2-largeurBtnÀPropos/2), float32(hauteurÉcran/2+hauteurBarreVol+50), float32(largeurBtnÀPropos), float32(hauteurBtnÀPropos))

	// Définir le texte "À Propos"
	params.texteÀPropos = "                                               C'est le Jeu de Sudoku.\n\nRègles :\nRemplissez la grille de manière à ce que chaque rangée, chaque colonne et chaque boîte de 3x3 contienne les chiffres de 1 à 9.\n\nAmusez-vous bien !"
	// Charger une police personnalisée pour la page "À Propos"
	params.policeÀPropos = rl.LoadFont("C:\\main\\customfont.ttf") // Remplacez par le chemin de votre police personnalisée
	defer rl.UnloadFont(params.policeÀPropos)

	for !rl.WindowShouldClose() {
		// Vérifier les ajustements de volume à l'aide des touches fléchées
		if rl.IsKeyPressed(rl.KeyRight) && params.volume < 5 {
			params.volume = int32(math.Min(5, float64(params.volume+1))) // S'assurer qu'il ne dépasse pas 5
			rl.SetSoundVolume(son, float32(params.volume)/5.0)           // Mettre à jour le volume
		} else if rl.IsKeyPressed(rl.KeyLeft) && params.volume > 0 {
			params.volume = int32(math.Max(0, float64(params.volume-1))) // S'assurer qu'il ne descend pas en dessous de 0
			rl.SetSoundVolume(son, float32(params.volume)/5.0)           // Mettre à jour le volume
		}

		// Calculer la position de la barre de volume en fonction de la valeur du volume
		params.positionBarreVol = int32(float32(params.volume) * (float32(largeurBarreVol) / 5))

		// Vérifier si la souris a cliqué sur le bouton "À Propos"
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), params.àProposBtn) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			params.afficherÀPropos = true // Afficher la fenêtre "À Propos" lorsque le bouton est cliqué
		}

		// Traiter la fenêtre "À Propos"
		if params.afficherÀPropos {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)

			// Dessiner le texte "À Propos" dans une fenêtre séparée avec une police personnalisée
			rl.DrawTextEx(params.policeÀPropos, params.texteÀPropos, rl.NewVector2(50, hauteurÉcran/2-80), float32(params.policeÀPropos.BaseSize), 2, rl.NewColor(0, 128, 255, 255))

			rl.EndDrawing()

			// Vérifier si une demande de fermeture a été faite dans la fenêtre "À Propos"
			if rl.WindowShouldClose() {
				params.afficherÀPropos = false // Fermer la fenêtre "À Propos"
			}
		} else {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)

			// Dessiner le texte du volume centré
			texteVolume := fmt.Sprintf("Volume : %d", params.volume)
			largeurTexte := rl.MeasureText(texteVolume, 20)
			rl.DrawText(texteVolume, largeurÉcran/2-int32(largeurTexte/2), hauteurÉcran/2-50, 20, rl.DarkGray)

			// Dessiner le conteneur de la barre de volume
			rl.DrawRectangle(largeurÉcran/2-int32(largeurBarreVol/2)-margeBarreVol, hauteurÉcran/2-20, largeurBarreVol+2*margeBarreVol, hauteurBarreVol, rl.DarkGray)

			// Dessiner la barre de volume en fonction de la valeur du volume
			couleurBarre := rl.Green // Vous pouvez changer la couleur selon vos préférences
			rl.DrawRectangle(largeurÉcran/2-int32(largeurBarreVol/2), hauteurÉcran/2-20, params.positionBarreVol, hauteurBarreVol, couleurBarre)

			// Dessiner les icônes de volume (volume haut et muet)
			positionVolumeHaut := rl.NewVector2(float32(largeurÉcran/2-tailleIcône/2), float32(hauteurÉcran/2+hauteurBarreVol+10))
			positionVolumeMuet := rl.NewVector2(float32(largeurÉcran/2-tailleIcône/2), float32(hauteurÉcran/2+hauteurBarreVol+10))

			if params.volume > 0 {
				rl.DrawTextureEx(params.volumeHaut, positionVolumeHaut, 0, float32(tailleIcône)/float32(params.volumeHaut.Width), rl.RayWhite)
			} else {
				rl.DrawTextureEx(params.volumeMuet, positionVolumeMuet, 0, float32(tailleIcône)/float32(params.volumeMuet.Width), rl.RayWhite)
			}

			// Dessiner le bouton "À Propos"
			rl.DrawRectangleRec(params.àProposBtn, rl.DarkGray)
			rl.DrawText("À Propos", int32(params.àProposBtn.X+params.àProposBtn.Width/2-30), int32(params.àProposBtn.Y+params.àProposBtn.Height/2-10), 20, rl.RayWhite)

			rl.EndDrawing()
		}
	}

	// Arrêter et décharger le son lorsque vous avez terminé
	rl.StopSound(son)
	rl.UnloadSound(son)

	rl.UnloadTexture(params.volumeHaut)
	rl.UnloadTexture(params.volumeMuet)

	rl.CloseWindow()
}
