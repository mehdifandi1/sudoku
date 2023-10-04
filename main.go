package main

import (
	"fmt"
	"time"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	params struct {
		volume           int32
		positionBarreVol int32
		volumeHaut       rl.Texture2D
		volumeMuet       rl.Texture2D
		afficherÀPropos bool
		texteÀPropos    string
		policeÀPropos   rl.Font
		largeurÉcran    int32
		hauteurÉcran    int32
		son             rl.Sound
		CLargeurEcran   int32
		CHauteurEcran   int32
	}

	MenuBtn struct {
		àProposBtn          rl.Rectangle
		playButton          rl.Rectangle
		settingsButton      rl.Rectangle
		quitButton          rl.Rectangle
		playButtonColor     rl.Color
		settingsButtonColor rl.Color
		quitButtonColor     rl.Color
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

const Debug = true

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
	initBtn()

	displayTextAnim()
	time.Sleep(300 * time.Millisecond)
	DrawTittle()
	time.Sleep(300 * time.Millisecond)
	DrawMenu()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(10, 10)

		// Affichez le titre "SUDOKU"
		rl.DrawText("SUDOKU", 300, 100, 60, rl.Black)

		drawButton(MenuBtn.playButton, MenuBtn.playButtonColor, "PLAY", 60)
		drawButton(MenuBtn.settingsButton, MenuBtn.settingsButtonColor, "SETTINGS", 60)
		drawButton(MenuBtn.quitButton, MenuBtn.quitButtonColor, "QUIT", 60)

		VerifBTCol()

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

	for i := 1; i < 300/time_boost1; i++ {
		rl.BeginDrawing()
		rl.DrawText("SUDOKU", x, y, 40, rl.Black)
		rl.ClearBackground(bcolor)
		y += (time_boost2*7)/int32(i)
		rl.EndDrawing()
	}

	if Debug {
		fmt.Println("Fin le l'animation")
	}

}

func DrawTittle() {

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

	//Initialisation des fichier a chargés
	//loadImages()
	//loadSound()

	rl.BeginDrawing()
	rl.ClearBackground(fcolor)
	rl.EndDrawing()

	time.Sleep(300 * time.Millisecond)
	rl.BeginDrawing()
	rl.DrawText("SUDOKU", x, y, 60, rl.Black)
	rl.EndDrawing()
	time.Sleep(1 * time.Second)

	for i := 1; i < 500/time_boost1; i++ {
		rl.BeginDrawing()
		rl.ClearBackground(fcolor)
		rl.DrawText("SUDOKU", x, y, 60, rl.Black)
		rl.DrawFPS(0, 0)
		rl.ClearBackground(fcolor)
		y -= time_boost2 * 4 /int32(i)
		rl.EndDrawing()
	}

	if Debug {
		fmt.Println("titre sudoku en place")
	}

	x = 0
	y = 0

	for i := 1; i < 10/time_boost1; i++ {
		rl.DrawFPS(x, y)
		rl.DrawText("SUDOKU", 300, 100, 60, rl.Black)
		rl.ClearBackground(fcolor)
		x+=time_boost2/(int32(i)*2)
		y+=time_boost2/(int32(i)*2)
	}

}

func loadSound() {
	// Initialisation du périphérique audio
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	// Charger le son MP3 comme un son normal (pas un flux musical)
	params.son = rl.LoadSound("Exyl - MOAI MONEY.mp3")
}

func DrawMenu() {
	var r uint8 = 255
	var g uint8 = 255
	var b uint8 = 255
	var a uint8 = 255
	var bcolor = rl.NewColor(r, g, b, a)

	a = 255

	for i := 1; i < 254/time_boost1; i++ {
		rl.BeginDrawing()
		bcolor = rl.NewColor(r, g, b, a)
		drawButton(MenuBtn.playButton, MenuBtn.playButtonColor, "PLAY", 60) // Augmentez la taille du texte
		rl.DrawRectangle(300, 200, 250, 60, bcolor)
		a -= time_boost*time_boost/(uint8(i))
		rl.EndDrawing()
	}

	a = 255

	for i := 1; i < 254/time_boost1; i++ {
		bcolor = rl.NewColor(r, g, b, a)
		drawButton(MenuBtn.settingsButton, MenuBtn.settingsButtonColor, "SETTINGS", 60)
		rl.DrawRectangle(300, 270, 250, 60, bcolor)
		a -= time_boost*time_boost/(uint8(i))
		rl.EndDrawing()
	}

	a = 255

	for i := 1; i < 254/time_boost1; i++ {
		rl.BeginDrawing()
		bcolor = rl.NewColor(r, g, b, a)
		drawButton(MenuBtn.quitButton, MenuBtn.quitButtonColor, "QUIT", 60)
		rl.DrawRectangle(300, 340, 250, 60, bcolor)
		a -= time_boost*time_boost/(uint8(i))
		rl.EndDrawing()
	}


}

func initBtn() {
	// Initialisation des Rectangles
	MenuBtn.playButton = rl.NewRectangle(300, 200, 250, 60)     // Augmentez la largeur et la hauteur des boutons
	MenuBtn.settingsButton = rl.NewRectangle(300, 270, 250, 60) // Augmentez la position Y et la taille des boutons
	MenuBtn.quitButton = rl.NewRectangle(300, 340, 250, 60)     // Augmentez la position Y et la taille des boutons

	// Initialisation de la couleurs des boutton
	MenuBtn.playButtonColor = rl.RayWhite
	MenuBtn.settingsButtonColor = rl.RayWhite
	MenuBtn.quitButtonColor = rl.RayWhite

}

func VerifBTCol() {

	mousePos := rl.GetMousePosition()

	//Vérifier les collisions de boutons
	if rl.CheckCollisionPointRec(mousePos, MenuBtn.playButton) {
		MenuBtn.playButtonColor = rl.Red
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			if Debug {
				fmt.Println("Boutton Play cliqué")
			}
			Starting_game()
		}
	} else {
		MenuBtn.playButtonColor = rl.RayWhite
	}

	if rl.CheckCollisionPointRec(mousePos, MenuBtn.settingsButton) {
		MenuBtn.settingsButtonColor = rl.Red
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			if Debug {
				fmt.Println("Boutton Setting cliqué")
			}
			setting_window()
		}
	} else {
		MenuBtn.settingsButtonColor = rl.RayWhite
	}

	if rl.CheckCollisionPointRec(mousePos, MenuBtn.quitButton) {
		MenuBtn.quitButtonColor = rl.Red
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			if Debug {
				fmt.Println("Boutton Quit cliqué")
			}
			rl.CloseWindow()
		}
	} else {
		MenuBtn.quitButtonColor = rl.RayWhite
	}
}

func drawButton(rect rl.Rectangle, color rl.Color, text string, fontSize int32) {
	textWidth := rl.MeasureText(text, fontSize) + 20 // Ajoutez un espace pour le texte
	textHeight := fontSize + 20                      // Ajoutez un espace pour le texte

	// Ajustez la taille et la position du rectangle en fonction de la taille du texte
	if Debug {
		fmt.Println("declaration de la position du boutton", text)
	}
	rect.Width = float32(textWidth)
	rect.Height = float32(textHeight)

	if Debug {
		fmt.Println("affichage du boutton")
	}
	rl.DrawRectangleLinesEx(rect, 2, color)
	rl.DrawText(text, int32(rect.X+10), int32(rect.Y+10), fontSize, rl.Black)
}
