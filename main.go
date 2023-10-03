package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
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
		largeurÉcran     int32
		hauteurÉcran     int32
		son              rl.Sound
	}
)

const (
	pasVolume         = 1.0 // Ajustez le pas de volume selon vos besoins
	largeurBarreVol   = 200 // Largeur de la barre de volume
	hauteurBarreVol   = 20  // Hauteur de la barre de volume
	margeBarreVol     = 5   // Marge autour de la barre de volume
	tailleIcône       = 32  // Taille des icônes de volume
	largeurBtnÀPropos = 150 // Largeur du bouton "À Propos"
	hauteurBtnÀPropos = 40  // Hauteur du bouton "À Propos"
	
)

const Debug = false
var time_boost uint8 = 8
var time_boost1 int = int(time_boost)
var time_boost2 int32 = int32(time_boost1)



func main() {
	if Debug {
		fmt.Print("Starting...")
	}

	params.largeurÉcran = 1200
	params.hauteurÉcran = 600
	rl.InitWindow(params.largeurÉcran, params.hauteurÉcran, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(144)




	displayTextAnim()
	time.Sleep(300 * time.Millisecond)
	menu()
	time.Sleep(300 * time.Millisecond)

	// Initialisation des Rectangles
	playButton := rl.NewRectangle(300, 200, 250, 60) // Augmentez la largeur et la hauteur des boutons
	settingsButton := rl.NewRectangle(300, 270, 250, 60) // Augmentez la position Y et la taille des boutons
	quitButton := rl.NewRectangle(300, 340, 250, 60) // Augmentez la position Y et la taille des boutons

	// Initialisation de la couleurs des boutton
	playButtonColor := rl.RayWhite
	settingsButtonColor := rl.RayWhite
	quitButtonColor := rl.RayWhite


	


	for !rl.WindowShouldClose() {
		mousePos := rl.GetMousePosition()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)

		// Affichez le titre "SUDOKU"
		rl.DrawText("SUDOKU", 300, 100, 60, rl.Black)

		// Draw buttons with their respective colors and text
		drawButton(playButton, playButtonColor, "PLAY", 60) // Augmentez la taille du texte
		drawButton(settingsButton, settingsButtonColor, "SETTINGS", 60) // Augmentez la taille du texte
		drawButton(quitButton, quitButtonColor, "QUIT", 60) // Augmentez la taille du texte

		//Vérifier les collisions de boutons
		if rl.CheckCollisionPointRec(mousePos, playButton) {
			playButtonColor = rl.Red
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				if Debug {
					fmt.Println("Boutton Play cliqué")
				}
				Starting_game()
			}
		} else {
			playButtonColor = rl.RayWhite
		}

		if rl.CheckCollisionPointRec(mousePos, settingsButton) {
			settingsButtonColor = rl.Red
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				if Debug {
					fmt.Println("Boutton Setting cliqué")
				}
				setting_window()
			}
		} else {
			settingsButtonColor = rl.RayWhite
		}

		if rl.CheckCollisionPointRec(mousePos, quitButton) {
			quitButtonColor = rl.Red
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				if Debug {
					fmt.Println("Boutton Quit cliqué")
				}
				rl.CloseWindow()
			}
		} else {
			quitButtonColor = rl.RayWhite
		}

		if Debug {
			fmt.Println("Running smoothly ...")
		}
		rl.EndDrawing()
	}
}

func loadImages() {
	params.volumeHaut = rl.LoadTexture("volumeup.png") // Remplacez par le chemin de votre image pour le volume haut
	params.volumeMuet = rl.LoadTexture("mute.png")     // Remplacez par le chemin de votre image pour le volume muet
}

func displayTextAnim() {
	var x int32 = 300
	var y int32 = 0

	var r uint8 = 0
	var g uint8 = 0
	var b uint8 = 0
	var a uint8 = 255

	bcolor := rl.NewColor(r, g, b, a)

	if Debug {
		fmt.Println("Setting up the background color")
	}

	for i := 0; i < 255/time_boost1; i++ {
		rl.BeginDrawing()
		bcolor = rl.NewColor(r, g, b, a)
		rl.ClearBackground(bcolor)
		r += time_boost
		g += time_boost
		b += time_boost
		rl.EndDrawing()
	}
	time.Sleep(2 * time.Second)

	if Debug {
		fmt.Println("Moving the logo")
	}

	// Initialisation des fichier a chargés
	loadImages()
	loadSound()
	 


	for i := 0; i < 200/time_boost1; i++ {
		rl.BeginDrawing()
		rl.DrawText("SUDOKU", x, y, 40, rl.Black)
		rl.ClearBackground(bcolor)
		y += time_boost2
		rl.EndDrawing()
	}

	if Debug {
		fmt.Println("Fin le l'animation")
	}

	
}

func menu() {
	var x int32 = 300
	var y int32 = 200

	var r uint8 = 255
	var g uint8 = 255
	var b uint8 = 255
	var a uint8 = 255
	
	if Debug {
		fmt.Println("creation couleur Font")
	}
	fcolor := rl.NewColor(r, g, b, a)

	if Debug {
		fmt.Println("mise en place du titre")
	}

	rl.BeginDrawing()
	rl.ClearBackground(fcolor)
	rl.EndDrawing()

	time.Sleep(300 * time.Millisecond)
	rl.BeginDrawing()
	rl.DrawText("SUDOKU", x, y, 60, rl.Black)
	rl.EndDrawing()
	time.Sleep(1 * time.Second)

	for i := 0; i < 100/time_boost1; i++ {
		rl.BeginDrawing()
		rl.ClearBackground(fcolor)
		rl.DrawText("SUDOKU", x, y, 60, rl.Black)
		rl.ClearBackground(fcolor)
		y -= time_boost2
		rl.EndDrawing()
	}

	if Debug {
		fmt.Println("titre sudoku en place")
	}


}

func drawButton(rect rl.Rectangle, color rl.Color, text string, fontSize int32) {
	textWidth := rl.MeasureText(text, fontSize) + 20 // Ajoutez un espace pour le texte
	textHeight := fontSize + 20                      // Ajoutez un espace pour le texte

	// Ajustez la taille et la position du rectangle en fonction de la taille du texte
	if Debug {
		fmt.Println("ajustement de la position de bouttons")
	}
	rect.Width = float32(textWidth)
	rect.Height = float32(textHeight)


	if Debug {
		fmt.Println("affichage des bouttons")
	}
	rl.DrawRectangleLinesEx(rect, 2, color)
	rl.DrawText(text, int32(rect.X+10), int32(rect.Y+10), fontSize, rl.Black)
}

func loadSound(){
	// Initialisation du périphérique audio
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	// Charger le son MP3 comme un son normal (pas un flux musical)
	params.son = rl.LoadSound("Exyl - MOAI MONEY.mp3")
} 