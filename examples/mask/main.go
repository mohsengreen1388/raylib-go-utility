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

	cube := rl.LoadModelFromMesh(rl.GenMeshCube(2, 2, 2))

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{2, 2, 6}
	cam.Projection = rl.CameraPerspective
	cam.Up = rl.Vector3{0, 1, 0}

	txdiff := rl.LoadTexture("./plasma.png")
	txMask := rl.LoadTexture("./mask.png")

	masks := ut.Mask{}

	masks.Init()
	masks.SetMapDiffuse([]*rl.Material{&cube.Materials[0]}, []*rl.Texture2D{&txdiff})

	masks.SetMapEmission([]*rl.Material{&cube.Materials[0]}, []*rl.Texture2D{&txMask})

	rl.SetTargetFPS(60)
	x := 0
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&cam, rl.CameraOrbital)
		x++
		masks.UpdateAnimation(x)
		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(cam)
		rl.DrawGrid(100, 1)
		rl.DrawModel(cube, rl.Vector3{0, 1, 0}, 1, rl.RayWhite)

		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
