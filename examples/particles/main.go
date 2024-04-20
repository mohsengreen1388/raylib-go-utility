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

	p := ut.Particle{}

	//p.InitializeParticles(100, rl.Vector3{10, 20, 10}, rl.Vector3{10,10,0},50, 150,true)
	p.InitializeParticles(100, rl.Vector3{0, 2, 10}, rl.Vector3{float32(rl.GetRandomValue(-1, 1)) / 64, float32(rl.GetRandomValue(1, 2)) / 264, float32(rl.GetRandomValue(-1, 1)) / 64}, 50, 150, true)
	tx := rl.LoadTexture("./smoke_04.png")
	//mesh := rl.GenMeshPlane(2, 2, 1, 1)
	//model := rl.LoadModelFromMesh(mesh)
	//p.Oneshot = true

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFirstPerson)
		p.UpdateParticles(rl.Vector3{float32(rl.GetRandomValue(-1, 1)) / 64, float32(rl.GetRandomValue(0, 1)) / 24, float32(rl.GetRandomValue(-1, 1)) / 64}, 40, 100)
		//p.UpdateParticles(rl.Vector3{float32(rl.GetRandomValue(100, -100))/74, float32(rl.GetRandomValue(100, -100))/74,1},20, 80)
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.BeginMode3D(camera)
		rl.DrawGrid(100, 1)
		//rl.DrawCube(rl.Vector3{2,2,2}, 2, 2, 2, rl.Red)
		//p.DrawCircleParicle(3, &model, &tx, rl.Vector3{-2,-2,0}, rl.Vector3{0, 1, 0}, rl.Vector3{5, 5, 5})
		//p.DrawParticles(&model, &tx, rl.Vector3{0.5, 0.5, 0.5})
		//rl.DrawCube(rl.Vector3{10,2,2}, 2, 2, 2, rl.Red)
		p.DrawParticlesBilborid(&camera, &tx, 4)
		rl.EndMode3D()
		//p.DrawParticles2dTexture(&tx)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
