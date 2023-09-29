package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	//"fmt"
	"time"
	
)








func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(240)


	display_text_Anim()
	time.Sleep(1 * time.Second)
	menu()
	time.Sleep(1 * time.Second)



	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(0,0)
		rl.EndDrawing()
	}
}