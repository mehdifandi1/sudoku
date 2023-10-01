package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const Debug = false
var time_boost uint8 = 6
var time_boost1 int = int(time_boost)
var time_boost2 int32 = int32(time_boost1)

func main() {
	if Debug {
		fmt.Print("Starting...")
	}
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	displayTextAnim()
	time.Sleep(300 * time.Millisecond)
	menu()
	time.Sleep(300 * time.Millisecond)

	// Initialize button rectangles
	playButton := rl.NewRectangle(300, 200, 250, 60) // Augmentez la largeur et la hauteur des boutons
	settingsButton := rl.NewRectangle(300, 270, 250, 60) // Augmentez la position Y et la taille des boutons
	quitButton := rl.NewRectangle(300, 340, 250, 60) // Augmentez la position Y et la taille des boutons

	// Initialize button colors
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

		// Check button collisions
		if rl.CheckCollisionPointRec(mousePos, playButton) {
			playButtonColor = rl.Red
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				fmt.Println("PLAY button clicked")
				Starting_game()
			}
		} else {
			playButtonColor = rl.RayWhite
		}

		if rl.CheckCollisionPointRec(mousePos, settingsButton) {
			settingsButtonColor = rl.Red
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				fmt.Println("SETTINGS button clicked")
			}
		} else {
			settingsButtonColor = rl.RayWhite
		}

		if rl.CheckCollisionPointRec(mousePos, quitButton) {
			quitButtonColor = rl.Red
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				fmt.Println("QUIT button clicked")
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

	for i := 0; i < 200/time_boost1; i++ {
		rl.BeginDrawing()
		rl.DrawText("SUDOKU", x, y, 40, rl.Black)
		rl.ClearBackground(bcolor)
		y += time_boost2
		rl.EndDrawing()
	}
}

func menu() {
	var x int32 = 300
	var y int32 = 200

	var r uint8 = 255
	var g uint8 = 255
	var b uint8 = 255
	var a uint8 = 255
	fcolor := rl.NewColor(r, g, b, a)

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

}

func drawButton(rect rl.Rectangle, color rl.Color, text string, fontSize int32) {
	textWidth := rl.MeasureText(text, fontSize) + 20 // Ajoutez un espace pour le texte
	textHeight := fontSize + 20                      // Ajoutez un espace pour le texte

	// Ajustez la taille et la position du rectangle en fonction de la taille du texte
	rect.Width = float32(textWidth)
	rect.Height = float32(textHeight)

	rl.DrawRectangleLinesEx(rect, 2, color)
	rl.DrawText(text, int32(rect.X+10), int32(rect.Y+10), fontSize, rl.Black)
}
