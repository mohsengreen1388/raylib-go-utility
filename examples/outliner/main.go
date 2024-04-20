package main

import (
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

func main() {

	screenWidth := int32(900)
	screenHeight := int32(500)
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE
	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - basic lighting")

	tx := rl.LoadTexture("./fudesumi.png")
	out := ut.OutLine{}
	out.Init(&tx,2.0,rl.Green)
	out.OutLineSize(3)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.BeginShaderMode(*out.Shader)
		rl.DrawTexture(tx, 50, 0, rl.RayWhite)
		rl.EndShaderMode()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
