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

	model := rl.LoadModel("./greenman.glb")
	sh := rl.LoadShader("./instanced.vs", "")

	in := ut.Instances{}
	in.MaxInstance = 15000
	in.Material = model.Materials[1]
	in.InitInstance(sh)

	rl.SetTargetFPS(60)

	for i := 0; i < in.MaxInstance; i++ {
		in.SetValueToMatrix(int32(i), float32(1*i), rl.NewVector3(float32(1*i), 0, 0), rl.NewVector3(0, 1, 0), rl.NewVector3(1, 1, 1))
	}

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)

		rl.BeginMode3D(camera)
		rl.DrawGrid(100, 1)
		in.Draw(*model.Meshes)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	println(model.Meshes.VertexCount)
	rl.CloseWindow()
}
