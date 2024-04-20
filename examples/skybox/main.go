package main

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
	
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - basic lighting")

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(2.0, 1.0, 4.0)
	camera.Target = rl.NewVector3(4, 1, 4)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	 sky := ut.SkyBox{}
	sky.Init(true,false, "./Daylight Box UV.png")

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawGrid(100, 1)
		sky.Draw()
		rl.DrawCube(rl.Vector3{0,-3,0},2,2,2,rl.Red)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
	sky.Unload()
	rl.CloseWindow()
}

