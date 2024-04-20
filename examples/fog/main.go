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

	cube := rl.LoadModelFromMesh(rl.GenMeshCube(2, 2, 2))
	tx := rl.LoadTexture("./texel_checker.png")

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{2, 2, 6}
	cam.Projection = rl.CameraPerspective
	cam.Up = rl.Vector3{0, 1, 0}


	fogs := ut.Fog{}
	fogs.Init(rl.Blue, 0.16)
	fogs.AddTheMaterialToFog([]*rl.Material{&cube.Materials[0]},[]*rl.Texture2D{&tx})

	cube.Materials[0].Shader = *fogs.Shader

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&cam, rl.CameraOrbital)
		if rl.IsKeyDown(rl.KeyUp) {
			fogs.FogDensity += 0.001
			fogs.UpdateFogDensity(fogs.FogDensity)
		}

		if rl.IsKeyDown(rl.KeyDown) {
			//fogDensity-= 0.001
			fogs.FogDensity -= 0.001
			fogs.UpdateFogDensity(fogs.FogDensity)
		}
		fogs.Update(cam.Position)
		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(cam)
		rl.DrawModel(cube, rl.Vector3{0, 1, 0}, 1, rl.RayWhite)
		rl.DrawGrid(100, 1)
		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

