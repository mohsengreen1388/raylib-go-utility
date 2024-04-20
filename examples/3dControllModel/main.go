package main

import (
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

var LastFrame float32
var x = 0.0
var y = 0.0
var z = -1

var bound rl.BoundingBox



func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(800, 600, "anm")

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{0, 12, -6}
	cam.Up = rl.Vector3{0, 1, 0}
	cam.Target = rl.Vector3{0, 0, 0}
	cam.Projection = rl.CameraPerspective

	model := rl.LoadModel("./mouse.glb")
	modelanm := rl.LoadModelAnimations("./mouse.glb")

	posModel := rl.Vector3{}
	move := ut.ControllerModel{}

	move.Init(&model, modelanm, &posModel, &cam.Position, true)

	cube2 := rl.LoadModelFromMesh(rl.GenMeshCube(1,1,1))
	cube2Pos := rl.Vector3{0,1,1}


	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

	//	rl.UpdateCamera(&cam,rl.CameraFirstPerson)
		
		 move.CameraTargetLockTheModel(&cam)
		
		if rl.IsKeyDown(rl.KeyUp) {
			move.Move(3, 1, ut.AxisCameraZ, rl.Vector2{0.02, 0.02}, rl.Vector3{0, 0, 0}, 0)
		}

		if rl.IsKeyDown(rl.KeyDown) {
			move.Move(3, 1, ut.AxisCameraZ, rl.Vector2{-0.02, -0.02}, rl.Vector3{0, 0, 1}, 180-0.9)
		}

		if rl.IsKeyReleased(rl.KeyUp) || rl.IsKeyReleased(rl.KeyDown) || rl.IsKeyReleased(rl.KeyRight) || rl.IsKeyReleased(rl.KeyLeft) {
			move.EnableDefaultAnimtion()
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			move.Once()
		}

		move.OnceRun(0, 1, rl.Vector3{0, 0, 1})

		move.DefaultLoopAnimtion(2, 1, rl.Vector3{0, 0, 1})
		
		bound = rl.NewBoundingBox(rl.Vector3{posModel.X,posModel.Y,posModel.Z},rl.Vector3{posModel.X+0.01,posModel.Y+0.01,posModel.Z+0.01})
		ch := rl.CheckCollisionBoxes(bound,rl.NewBoundingBox(rl.Vector3{cube2Pos.X,cube2Pos.Y,cube2Pos.Z},rl.Vector3{cube2Pos.X,cube2Pos.Y,cube2Pos.Z}))
	
		println(ch)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(0, 0)
		rl.BeginMode3D(cam)
		rl.DrawGrid(100, 1)
		rl.DrawModel(cube2,cube2Pos,1,rl.Blue)
		rl.DrawModelEx(model, posModel, rl.Vector3{1, 0, 0}, 90, rl.Vector3{0.01, 0.01, 0.01}, rl.RayWhite)
		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
