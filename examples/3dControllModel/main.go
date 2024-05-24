package main

import (
	"github.com/mohsengreen1388/raylib-go-utility/ode"
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	ut "github.com/mohsengreen1388/raylib-go-utility/utility"
)

func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(800, 600, "anm")

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{0, 5, -6}
	cam.Up = rl.Vector3{0, 1, 0}
	cam.Target = rl.Vector3{0, 0, 0}
	cam.Projection = rl.CameraPerspective

	model := rl.LoadModel("./greenman.glb")
	modelanm := rl.LoadModelAnimations("./greenman.glb")
	sword := rl.LoadModel("./greenman_sword.glb")

	posModel := rl.Vector3{0, 0, 0}
	move := ut.ControllerModel{}

	cube2 := rl.LoadModelFromMesh(rl.GenMeshCube(1, 1, 1))

	ode.RlBeginConfig()
	world, space, jointGroup := ode.RlInit(1, ode.NewVector3(0, -1, 0), true, 1, ode.NewVector4(0, 1, 0, 0))
	//world.SetContactSurfaceLayer(0.0)
	world.SetCFM(0.01)
	world.SetERP(0.2)
	contact := make([]*ode.Contact, 64)
	mode := ode.SoftCFMCtParam | ode.BounceCtParam
	TriData := ode.RlGeomTriMeshDataBuildSingleAnm(&model.GetMeshes()[0])
	TriMesh := ode.RlNewTriMesh(space, &TriData)
	//	TrimeshBody, _ := ode.RlMeshToDynamic(world, space, &TriMesh, &TriData, 1, ode.NewVector3(0, 1, 0))
	//bodyMouse, _, _ := ode.RlNewBox(world, space, ode.NewVector3(0.5, 0.5, 0.5), ode.NewVector3(0, 2.5, 0), ode.NewVector3(0.5, 0.5, 0.5), 1, false, false)

	bodyBox, _, _ := ode.RlNewBox(world, space, ode.NewVector3(1, 1, 1), ode.NewVector3(2, 2, 0), ode.NewVector3(1, 1, 1), 1, false, false)
	ode.RlEndConfig()
	move.Init(&model, nil, modelanm, &posModel, &cam.Position, true)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		ode.RlStep(world, space, jointGroup, 60, ode.RlCollideCallBack(world, 64, *jointGroup, contact, mode, 1), func() {
			move.AddCollisionByBody(bodyBox, true, 4, 1)
			move.SyncTransformTheModel()
		})
		move.CameraTargetLockTheModel(&cam)

		if rl.IsKeyDown(rl.KeyUp) {
			move.Move(3, 1, ut.AxisCameraZ, rl.Vector2{0.02, 0.02}, rl.Vector3{0, 0, 0})
		}

		if rl.IsKeyDown(rl.KeyDown) {
			move.Move(3, 1, ut.AxisCameraZ, rl.Vector2{-0.02, -0.02}, rl.Vector3{0, 34.5, 0})
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			move.Move(3, 1, ut.AxisCameraX, rl.Vector2{0.02, 0.02}, rl.Vector3{0, 45.5, 0})
		}

		if rl.IsKeyReleased(rl.KeyUp) || rl.IsKeyReleased(rl.KeyDown) || rl.IsKeyReleased(rl.KeyRight) || rl.IsKeyReleased(rl.KeyLeft) {
			move.EnableDefaultAnimtion()
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			move.Once()
		}

		move.OnceRun(1, 3)

		move.DefaultLoopAnimtion(2, 1)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(0, 0)
		rl.BeginMode3D(cam)
		rl.DrawGrid(100, 1)
		move.DrawModelControll(rl.Vector3{0, 0, 0}, 0, rl.Vector3{1, 1, 1}, rl.RayWhite)
		ode.RlBeginDrawingDebug()
		ode.RlDrawTriMesh(&TriMesh, rl.Blue)
		ode.RlDrawBox(&cube2, bodyBox, rl.Vector3{1, 1, 1}, rl.Red)
		ode.RlEndDrawingDebug()

		//move.DrawSkeletBindPose(posModel, rl.Vector3{0, 0, 0}, 0, 2)
		move.DrawSkeletAnimtion(1)
		move.JointModelPostionAndDraw(&sword, 10, 1, 1, rl.Red)
		rl.EndMode3D()
		rl.EndDrawing()
	}
	ode.RlClose(world, space, jointGroup)
	rl.CloseWindow()
}
