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

	cube :=  rl.LoadModel("./model/PRD0U20WISTNTZ7EXFS37EFSJ.obj") //rl.LoadModelFromMesh(rl.GenMeshCube(2, 2, 2))
	//tx := rl.LoadTexture("./examples/texel_checker.png")

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{2, 2, 6}
	cam.Projection = rl.CameraPerspective
	cam.Up = rl.Vector3{0, 1, 0}

	sx := ut.Cobineshader()
	fogs := ut.Fog{}
	fogs.SetCombineShader(&sx)
	fogs.Init(rl.Blue, 0.16)

	l := ut.Light{}
	l.SetCombineShader(&sx)
	l.Init(0.2, rl.Vector3{1, 1, 1})
	l1 := l.NewLight(ut.LightTypePoint, rl.Vector3{0, 3, 0}, rl.Vector3{}, rl.Yellow, 2,&l.Shader)
	
	p := ut.PhysicRender{}
	p.SetCombineShader(&sx)
	p.Init()
	p.MetallicValue(0.3)
	p.RoughnessValue(0.1)
	p.AmbientColor(rl.Vector3{1, 1, 1}, 0.2)

	cube.GetMaterials()[0].Shader = sx

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
		p.UpadteByCamera(cam.Position)
		rl.BeginDrawing()
		rl.DrawFPS(10, 20)
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(cam)
		rl.DrawModel(cube, rl.Vector3{0, 1, 0}, 1, rl.RayWhite)
		l.DrawSpherelight(&l1)
		rl.DrawGrid(100, 1)
		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

