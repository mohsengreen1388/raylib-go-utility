package main

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
)

func main() {
	screenWidth := int32(900)
	screenHeight := int32(500)
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - basic lighting")

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(2.0, 4.0, 6.0)
	camera.Target = rl.NewVector3(0.0, 0.5, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	mesh := rl.GenMeshCube(2, 2, 2)
	cube := rl.LoadModelFromMesh(mesh)
	tx := rl.LoadTexture("./space.png")

	wa := ut.Wave{}

	wa.Init([]*rl.Material{&cube.Materials[0]}, []*rl.Texture2D{&tx})

	cube.Materials[0].Shader = *wa.Shader

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		// Update
		//----------------------------------------------------------------------------------
		wa.UpdateWave()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.BeginMode3D(camera)
		rl.DrawGrid(100, 1)

		rl.DrawModel(cube, rl.Vector3{0, 2, 0}, 2, rl.RayWhite)

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
