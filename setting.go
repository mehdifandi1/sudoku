package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func parseResolution(resolution string) (int, int) {
	var width, height int
	fmt.Sscanf(resolution, "%dx%d", &width, &height)
	return width, height
}

func setting_window() {

	rl.PlaySound(params.son)

	// Save button on the bottom right
	saveBtn = rl.NewRectangle(float32(params.largeurÉcran-largeurBtnSave-marginRight), float32(params.hauteurÉcran-hauteurBtnSave-marginBottom), float32(largeurBtnSave), float32(hauteurBtnSave))
	// About button to the left of the Save button
	

	// FPS button under the Resolution button
	//fpsBtn := rl.NewRectangle(float32(marginLeft), resolutionBtn.Y+resolutionBtn.Height+marginBottom, float32(largeurBtnFPS), float32(hauteurBtnFPS))
	// Initialize current FPS to the default value
	fpsBtn = rl.NewRectangle(float32(marginLeft), resolutionBtn.Y+resolutionBtn.Height+marginBottom, float32(largeurBtnFPS), float32(hauteurBtnFPS))

	MenuBtn.aboutbuttonColor = rl.Black
	MenuBtn.aboutbutton = rl.NewRectangle(float32((params.largeurÉcran/2) + (params.largeurÉcran*2/10)), float32(params.hauteurÉcran/2 + (params.hauteurÉcran*25/100) ),250,50)
	



	for !rl.WindowShouldClose() {

		// Existing code for drawing volume controls, "À Propos" button, and other functionality
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		collBtsSetting()
		soundfunc(params.son)

		
		drawButton(MenuBtn.aboutbutton, MenuBtn.aboutbuttonColor, "About", 60)

		resolutionText := fmt.Sprintf("Resolution : %s", resolutions[numr])
		rl.DrawText(resolutionText, int32(resolutionBtn.X+10), int32(resolutionBtn.Y+resolutionBtn.Height/2-10), 20, rl.Black)
		//rl.DrawRectangleRec(MenuBtn.aboutbutton, rl.DarkGray)
		//rl.DrawText("À Propos", int32(MenuBtn.aboutbutton.X+MenuBtn.aboutbutton.Width/2-30), int32(MenuBtn.aboutbutton.Y+MenuBtn.aboutbutton.Height/2-10), 20, rl.RayWhite)
		//rl.DrawRectangleRec(fpsBtn, rl.DarkGray)
		rl.DrawText(fpsBtnText[fpsBtnState], int32(fpsBtn.X+fpsBtn.Width/2)-30, int32(fpsBtn.Y+fpsBtn.Height/2)-10, 20, rl.Black)

		rl.EndDrawing()

	}
	rl.StopSound(params.son)
	rl.UnloadSound(params.son)

	rl.UnloadTexture(params.volumeHaut)
	rl.UnloadTexture(params.volumeMuet)

	rl.CloseWindow()
}

func collBtsSetting() {
	mousePos := rl.GetMousePosition()

	//bouton à Propos
	if rl.CheckCollisionPointRec(mousePos, MenuBtn.aboutbutton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		params.afficherÀPropos = true
	}

	//bouton resolution
	if rl.CheckCollisionPointRec(mousePos, resolutionBtn) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if numr < len(resolutions)-1 {
			numr++
		} else {
			numr = 0
		}
		width, height = parseResolution(resolutions[numr])

	}

	//FPS button
	if rl.CheckCollisionPointRec(mousePos, fpsBtn) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		numf = (numf + 1) % len(fpsOptions)
		currentFPS = fpsOptions[numf]

	}

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), saveBtn) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		//save on le regle apres!!!!!!!!!!!!
		params.Fps = currentFPS
		params.hauteurÉcran = int32(height)
		params.largeurÉcran = int32(width)
		if Debug {
			fmt.Println("Saving ...")
		}
		BackToMenu = true
		main()

	}
}

func soundfunc(son rl.Sound) {

	// Volume text slightly under the top left
	volumeTextPos := rl.NewVector2(float32(marginLeft), float32(marginTop))
	params.positionBarreVol = int32(float32(params.volume) * (float32(largeurBarreVol) / 5))
	// Draw Volume Text
	texteVolume := fmt.Sprintf("Volume : %d", params.volume)
	rl.DrawText(texteVolume, int32(volumeTextPos.X), int32(volumeTextPos.Y), 20, rl.DarkGray)
	// Volume bar in front of the text with a small space
	volumeBarPos := rl.NewVector2(float32(marginLeft)+float32(largeurBtnSave)+0.5, float32(marginTop))
	// Volume images in front of the volume bar with a small space
	positionVolumeHaut := rl.NewVector2(volumeBarPos.X+float32(largeurBarreVol)+0.5, float32(marginTop))
	// Resolution button under the Volume components
	resolutionBtn = rl.NewRectangle(float32(marginLeft), positionVolumeHaut.Y+float32(hauteurBtnSave)+marginBottom, float32(largeurBtnResolution), float32(hauteurBtnResolution))

	// Draw Volume Bar
	rl.DrawRectangle(int32(volumeBarPos.X-margeBarreVol), int32(volumeBarPos.Y), largeurBarreVol+2*margeBarreVol, hauteurBarreVol, rl.DarkGray)
	couleurBarre := rl.Green
	rl.DrawRectangle(int32(volumeBarPos.X), int32(volumeBarPos.Y), int32(params.positionBarreVol), hauteurBarreVol, couleurBarre)

	// Draw Volume Images
	if params.volume > 0 {
		rl.DrawTextureEx(params.volumeHaut, positionVolumeHaut, 0, float32(tailleIcône)/float32(params.volumeHaut.Width), rl.RayWhite)
	} else {
		rl.DrawTextureEx(params.volumeMuet, positionVolumeHaut, 0, float32(tailleIcône)/float32(params.volumeMuet.Width), rl.RayWhite)
	}

	if rl.IsKeyPressed(rl.KeyRight) && params.volume < 5 {
		params.volume = int32(math.Min(5, float64(params.volume+1)))
		rl.SetSoundVolume(son, float32(params.volume)/5.0)
	} else if rl.IsKeyPressed(rl.KeyLeft) && params.volume > 0 {
		params.volume = int32(math.Max(0, float64(params.volume-1)))
		rl.SetSoundVolume(son, float32(params.volume)/5.0)
	}
}
