package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	//"fmt"
	"time"
)




func display_text_Anim() {

	var x int32 = 300
	var y int32 = 0
	
	var r uint8 = 0
	var g uint8 = 0
	var b uint8 = 0
	var a uint8 = 255
	var bcolor = rl.NewColor(r,g,b,a)

	
	for i := 0; i < 255; i++ {
		rl.BeginDrawing()
		bcolor = rl.NewColor(r,g,b,a)
		rl.DrawFPS(0,0)
		rl.ClearBackground(bcolor)
		r++
		g++
		b++
		rl.EndDrawing()
	}
	time.Sleep(2 * time.Second)

	for i := 0; i < 200; i++ {
		rl.BeginDrawing()
		rl.DrawFPS(0,0)
		rl.DrawText("SUDOKU", x, y, 40, rl.Black)
		rl.ClearBackground(bcolor)
		y++
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
	var fcolor = rl.NewColor(r, g, b, a)
	

	rl.BeginDrawing()
	rl.ClearBackground(fcolor)
	rl.EndDrawing()
	

	time.Sleep(1 * time.Second)
	rl.BeginDrawing()
	rl.DrawFPS(0,0)
	rl.DrawText("SUDOKU", x, y, 60, rl.Black)
	rl.EndDrawing()
	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		rl.BeginDrawing()
		rl.ClearBackground(fcolor)
		rl.DrawText("SUDOKU", x, y, 60, rl.Black)
		rl.ClearBackground(fcolor)
		y--
		rl.EndDrawing()
	}

	r = 255
	g = 255
	b = 255

	for i := 0; i < 250; i++ {
		rl.BeginDrawing()
		rl.DrawFPS(0,0)
		fcolor = rl.NewColor(r, g, b, a)
		rl.DrawText("PLAY", x, y+80, 40, fcolor)
		r--
		g--
		b--
		rl.EndDrawing()
	}

	r = 255	
	g = 255
	b = 255

	for i := 0; i < 250; i++ {
		rl.BeginDrawing()
		rl.DrawFPS(0,0)
		fcolor = rl.NewColor(r, g, b, a)
		rl.DrawText("SETTING", x, y+120, 40, fcolor)
		r--
		g--
		b--
		rl.EndDrawing()
	}

	r = 255
	g = 255
	b = 255

	for i := 0; i < 250; i++ {
		rl.BeginDrawing()
		rl.DrawFPS(0,0)
		fcolor = rl.NewColor(r, g, b, a)
		rl.DrawText("QUIT", x, y+160, 40, fcolor)
		r--
		g--
		b--
		rl.EndDrawing()
	}

}


func selec_menu(){

	const X float32 = 0
	const Y float32 = 0
	const Z float32 = 0

	

	

	rl.BeginDrawing()
	rl.DrawBoundingBox(rl.BoundingBox{},rl.Black)
	rl.EndDrawing()


}