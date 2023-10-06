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
		afficherÀPropos  bool
		texteÀPropos     string
		policeÀPropos    rl.Font
		largeurÉcran     int32
		hauteurÉcran     int32
		son              rl.Sound
		CLarEcran        int32
		CHauEcran        int32
		Fps              int32
	}

	MenuBtn struct {
		aboutbutton    rl.Rectangle
		playButton     rl.Rectangle
		settingsButton rl.Rectangle
		quitButton     rl.Rectangle
		savebutton     rl.Rectangle

		aboutbuttonColor    rl.Color
		playButtonColor     rl.Color
		settingsButtonColor rl.Color
		quitButtonColor     rl.Color
		savebuttoncolor     rl.Color

		volume           int32
		positionBarreVol int32
		volumeHaut       rl.Texture2D
		volumeMuet       rl.Texture2D
		afficherÀPropos  bool
		texteÀPropos     string
		policeÀPropos    rl.Font
	}

	resolutionBtn rl.Rectangle
	numr          int
	resolutions   = []string{"800x600", "1024x768", "1280x720"}
	numf          int
	currentFPS    int32
	fpsOptions    = []int32{30, 60, 120, 240}
	fpsBtnState   int // 0: 30 FPS, 1: 60 FPS, 2: 120 FPS, 3: 240 FPS
	fpsBtnText    = []string{"30 FPS", "60 FPS", "120 FPS", "240 FPS"}
	fpsBtn        rl.Rectangle
	saveBtn       rl.Rectangle
	width         int
	height        int
	BackToMenu    bool = false
)

const (
	pasVolume            = 1.0 // Ajustez le pas de volume selon vos besoins
	largeurBarreVol      = 200 // Largeur de la barre de volume
	hauteurBarreVol      = 20  // Hauteur de la barre de volume
	margeBarreVol        = 5   // Marge autour de la barre de volume
	tailleIcône          = 32  // Taille des icônes de volume
	largeurBtnÀPropos    = 150 // Largeur du bouton "À Propos"
	hauteurBtnÀPropos    = 40  // Hauteur du bouton "À Propos"
	largeurBtnResolution = 200 // Increased width for the resolution button
	hauteurBtnResolution = 40
	largeurBtnSave       = 100
	hauteurBtnSave       = 40
	largeurBtnFPS        = 100 // Width of the FPS button
	hauteurBtnFPS        = 40  // Height of the FPS button
	marginTop            = 20
	marginLeft           = 20
	marginRight          = 20
	marginBottom         = 20
)

const Debug = false

var time_boost uint8 = 8
var time_boost1 int = int(time_boost)
var time_boost2 int32 = int32(time_boost1)

func main() {

	var x int32 = 0
	var y int32 = 0

	if BackToMenu == false {

		if Debug {
			fmt.Print("Starting...")
		}

		params.largeurÉcran = 960
		params.hauteurÉcran = 540

		params.CLarEcran = params.largeurÉcran / 2
		params.CHauEcran = params.hauteurÉcran / 2

		rl.InitWindow(params.largeurÉcran, params.hauteurÉcran, "raylib [core]  - Sudoku window")
		defer rl.CloseWindow()
		rl.SetTargetFPS(144)
		xpi, ypi, xsi, ysi, xqi, yqi := initBtn()

		x, y = TitreDec()
		time.Sleep(300 * time.Millisecond)
		x, y = TritreMont(x, y)
		time.Sleep(300 * time.Millisecond)
		DrawMenu(xpi, ypi, xsi, ysi, xqi, yqi)
	}

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(10, 10)

		// Affichez le titre "SUDOKU"
		rl.DrawText("SUDOKU", x, y, 60, rl.Black)

		if Debug {
			fmt.Println(x, " ", y)
		}

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

func TitreDec() (int32, int32) {

	var x int32 = params.CLarEcran - (12 * params.largeurÉcran / 100)
	var y int32 = params.CHauEcran - (2 * params.hauteurÉcran / 10)

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
		y += (time_boost2 * 2) / int32(i)
		rl.EndDrawing()
	}
	if Debug {
		fmt.Println("Fin le l'animation", x, " ", y)
	}
	return x, y

}

func TritreMont(x int32, y int32) (int32, int32) {

	var r uint8 = 255
	var g uint8 = 255
	var b uint8 = 255
	var a uint8 = 255

	var Fx int32 = 0
	var Fy int32 = 0

	if Debug {
		fmt.Println("creation couleur Font")
	}
	fcolor := rl.NewColor(r, g, b, a)

	if Debug {
		fmt.Println("mise en place du titre")
	}

	//Initialisation des fichier a chargés
	loadImages()
	loadSound()

	rl.BeginDrawing()
	rl.ClearBackground(fcolor)
	rl.EndDrawing()

	time.Sleep(300 * time.Millisecond)
	rl.BeginDrawing()
	rl.DrawText("SUDOKU", x, y, 60, rl.Black)
	rl.EndDrawing()
	time.Sleep(300 * time.Millisecond)

	for i := 1; i < 300/time_boost1; i++ {
		rl.BeginDrawing()
		rl.ClearBackground(fcolor)
		rl.DrawText("SUDOKU", x, y, 60, rl.Black)
		rl.DrawFPS(0, 0)
		rl.ClearBackground(fcolor)
		y -= time_boost2 * 2 / int32(i)
		rl.EndDrawing()
	}

	if Debug {
		fmt.Println("titre sudoku en place")
	}

	for i := 1; i < 10/time_boost1; i++ {
		rl.DrawFPS(Fx, Fy)
		rl.DrawText("SUDOKU", x, y, 60, rl.Black)
		rl.ClearBackground(fcolor)
		Fx += time_boost2 / (int32(i) * 2)
		Fy += time_boost2 / (int32(i) * 2)
	}

	if Debug {
		fmt.Println("Fin le l'animation", x, " ", y)
	}
	return x, y
}

func loadSound() {
	// Initialisation du périphérique audio
	rl.InitAudioDevice()

	// Charger le son MP3 comme un son normal (pas un flux musical)
	params.son = rl.LoadSound("music.mp3")
	rl.PlaySound(params.son)
}

func DrawMenu(xpi int32, ypi int32, xsi int32, ysi int32, xqi int32, yqi int32) {
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
		rl.DrawRectangle(xpi, ypi, 250, 60, bcolor)
		a -= time_boost * time_boost / (uint8(i))
		rl.EndDrawing()
	}

	a = 255

	for i := 1; i < 100/time_boost1; i++ {
		bcolor = rl.NewColor(r, g, b, a)
		drawButton(MenuBtn.settingsButton, MenuBtn.settingsButtonColor, "SETTINGS", 60)
		rl.DrawRectangle(xsi, ysi, 350, 60, bcolor)
		a -= time_boost * time_boost / (uint8(i))
		rl.EndDrawing()
	}

	a = 255

	for i := 1; i < 254/time_boost1; i++ {
		rl.BeginDrawing()
		bcolor = rl.NewColor(r, g, b, a)
		drawButton(MenuBtn.quitButton, MenuBtn.quitButtonColor, "QUIT", 60)
		rl.DrawRectangle(xqi, yqi, 250, 60, bcolor)
		a -= time_boost * time_boost / (uint8(i))
		rl.EndDrawing()
	}

}

func initBtn() (int32, int32, int32, int32, int32, int32) {

	var xp float32 = float32(params.CLarEcran - params.largeurÉcran/10)
	var yp float32 = float32(params.CHauEcran - (8 * params.hauteurÉcran / 100))

	var xs float32 = float32(params.CLarEcran - params.largeurÉcran/10)
	var ys float32 = float32(params.CHauEcran + (5 * params.hauteurÉcran / 100))

	var xq float32 = float32(params.CLarEcran - params.largeurÉcran/10)
	var yq float32 = float32(params.CHauEcran + (18 * params.hauteurÉcran / 100))

	// Initialisation des Rectangles
	MenuBtn.playButton = rl.NewRectangle(xp, yp, 250, 50)     // Augmentez la largeur et la hauteur des boutons
	MenuBtn.settingsButton = rl.NewRectangle(xs, ys, 250, 50) // Augmentez la position Y et la taille des boutons
	MenuBtn.quitButton = rl.NewRectangle(xq, yq, 250, 50)     // Augmentez la position Y et la taille des boutons

	// Initialisation de la couleurs des boutton
	MenuBtn.playButtonColor = rl.RayWhite
	MenuBtn.settingsButtonColor = rl.RayWhite
	MenuBtn.quitButtonColor = rl.RayWhite

	xpi := int32(xp)
	ypi := int32(yp)

	xsi := int32(xs)
	ysi := int32(ys)

	xqi := int32(xq)
	yqi := int32(yq)

	return xpi, ypi, xsi, ysi, xqi, yqi

}

func VerifBTCol() {

	mousePos := rl.GetMousePosition()

	//Vérifier les collisions de boutons
	if rl.CheckCollisionPointRec(mousePos, MenuBtn.playButton) {
		MenuBtn.playButtonColor = rl.Gray
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
		MenuBtn.settingsButtonColor = rl.Gray
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
		MenuBtn.quitButtonColor = rl.Gray
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
