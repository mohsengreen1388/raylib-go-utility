package main

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
	"github.com/mohsengreen1388/raylib-go-utility/physics/ode"
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
	in.MaxInstance = 20
	in.Material = model.GetMaterials()[1]
	in.InitInstance(sh)

	ode.RlBeginConfig()
	world, space, jointGroup := ode.RlInit(1, ode.NewVector3(0, -1, 0), true, 1, ode.NewVector4(0, 1, 0, 0))
	triMeshData := ode.RlGeomTriMeshDataBuildSingle(model.Meshes, 1)
	triMesh := ode.RlNewTriMesh(space, &triMeshData)
	contact := make([]*ode.Contact, 64)
	mode := ode.SoftCFMCtParam | ode.BounceCtParam
	ode.RlEndConfig()

	rl.SetTargetFPS(60)
	for i := 0; i < in.MaxInstance; i++ {
		in.SetValueToMatrix(int32(i), rl.NewVector3(float32(1*i), 0, 0), rl.NewVector3(0, float32(1*i), 0), rl.NewVector3(1, 1, 1))
	}
	for !rl.WindowShouldClose() {
		ode.RlStep(world, space, jointGroup, 60, ode.RlCollideCallBack(world, 64, *jointGroup, contact, mode, 1), func() {
			triMesh.SetPosition(ode.NewVector3(6, 0, 0))
			r := ode.RlSetRotation(rl.Vector3{0, 6, 0})
			triMesh.SetRotation(r)
		})
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(camera)
		rl.DrawGrid(100, 1)
		in.Draw(*model.Meshes)
		ode.RlBeginDrawingDebug()
		ode.RlDrawTriMesh(&triMesh, rl.Red)
		ode.RlEndDrawingDebug()
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	ode.RlClose(world, space, jointGroup)
	rl.CloseWindow()
}
