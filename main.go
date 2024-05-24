package main

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	

)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - bmfont and ttf sprite fonts loading")

	
	
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}



	rl.CloseWindow()
}
