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

	floar := rl.LoadModel("./model/plane.glb")
	crash := rl.LoadModel("./model/PRD0U20WISTNTZ7EXFS37EFSJ.obj")

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{2, 2, 6}
	cam.Projection = rl.CameraPerspective
	cam.Up = rl.Vector3{0, 1, 0}

	physic := ut.PhysicRender{}
	physic.Init(false)
	physic.MetallicValue(0.3)
	physic.RoughnessValue(0.1)

	physic.TextureMapAlbedo(&floar.Materials[0], rl.LoadTexture("./model/road_a.png"))
	physic.TextureMapMetalness(&floar.Materials[0], rl.LoadTexture("./model/road_mra.png"))
	physic.TextureMapNormal(&floar.Materials[0], rl.LoadTexture("./model/road_n.png"))

	l := physic.CreateLight(0, rl.Vector3{-1.0, 1.0, -4.0}, rl.Vector3{0.0, 1.0, -1.0}, rl.Yellow, 4.0)
	l1 := physic.CreateLight(1, rl.Vector3{0.0, 4.0, -1.0}, rl.Vector3{0.0, 0.0, 0.0}, rl.White, 8.3)

	floar.Materials[0].Shader = *physic.Shader
	crash.Materials[0].Shader = *physic.Shader

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		physic.UpadteByCamera(cam.Position)
		rl.UpdateCamera(&cam, rl.CameraOrbital)
		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(cam)
		rl.DrawModel(floar, rl.Vector3{0, 0, 0}, 6.0, rl.White)
		rl.DrawModel(crash, rl.Vector3{0, 1.2, 0}, 1, rl.RayWhite)
		physic.DrawSphereLoctionLight(l, rl.Yellow)
		physic.DrawSphereLoctionLight(l1, rl.White)
		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
