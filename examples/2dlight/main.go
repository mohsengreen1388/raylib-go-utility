package main

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
	"image/color"
)

func main() {

	screenWidth := int32(900)
	screenHeight := int32(500)
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - basic lighting")

	lightx := ut.Lights2d{}
	lightx.SpotLightInit(float32(123.2), float32(0.9), rl.Black)
	lighttx := rl.LoadTexture("./x.png")
	tx1 := rl.LoadTexture("./fudesumi.png")
	rl.SetTargetFPS(60)
	x := 120
	for !rl.WindowShouldClose() {
		x += 10
		if x > 900 {
			x = 120
		}
		lightx.UpdatePosition(rl.Vector2{float32(rl.GetMouseX()), float32(rl.GetMouseY())})
		lightx.UpdateOptions(float32(x), float32(0.9), rl.Black)
		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.DrawTexture(tx1, 70, 20, rl.RayWhite)
		lightx.PointLight(&lighttx, rl.Vector2{20, 20}, 0.6, []color.RGBA{rl.Fade(rl.Yellow, 0.3), rl.Fade(rl.Red, 0.2)})
		lightx.SpotLightDraw()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
