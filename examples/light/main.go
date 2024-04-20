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

	cube := rl.LoadModel("./model/PRD0U20WISTNTZ7EXFS37EFSJ.obj")
	//tx := rl.LoadTexture("./examples/texel_checker.png")

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{2, 2, 6}
	cam.Projection = rl.CameraPerspective
	cam.Up = rl.Vector3{0, 1, 0}

	
	light := ut.Light{}
	light.Init(0.5,rl.Vector3{1,1,1})
	lighshape := light.NewLight(ut.LightTypeDirectional,rl.Vector3{2,3,0},rl.Vector3{0,2,0},rl.Yellow,1,&light.Shader)
	
	cube.Materials[0].Shader = light.Shader

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&cam, rl.CameraOrbital)
		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(cam)
		rl.DrawModel(cube, rl.Vector3{0, 1, 0}, 1, rl.RayWhite)
		light.DrawSpherelight(&lighshape)
		rl.DrawGrid(100, 1)
		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

