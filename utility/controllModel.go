package utility

import (
	"github.com/mohsengreen1388/raylib-go-utility/ode"
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

type AxisCamera int8

const (
	AxisCameraX AxisCamera = iota
	AxisCameraY
	AxisCameraZ
)

type ControllerModel struct {
	Model                 *rl.Model
	Body                  *ode.Body
	ModelPosition         *rl.Vector3
	ModelAnimation        []rl.ModelAnimation
	indexAnimation        int32
	CameraPosition        *rl.Vector3
	defaultAnimtionStatus bool
	OnceStatus            bool
	lastAngle             rl.Vector3
	CounterFrame          int32
	Scale                 float64
}

// ModelPosition set the  drawModel for postion
func (con *ControllerModel) Init(model *rl.Model, body *ode.Body, modelAnimation []rl.ModelAnimation, modelPosition, cameraPosition *rl.Vector3,scaleModel float64,enableDefaultAnimtion bool) {
	con.Model = model
	con.Body = body
	con.ModelAnimation = modelAnimation
	con.CameraPosition = cameraPosition
	con.defaultAnimtionStatus = enableDefaultAnimtion
	con.Scale = scaleModel
	con.exitBodyCondition(func() {
		con.Body.SetPosition(ode.NewVector3(float64(modelPosition.X), float64(modelPosition.Y), float64(modelPosition.Z)))
		con.ModelPosition = modelPosition
	}, func() {
		con.ModelPosition = modelPosition
	})
}

func (con *ControllerModel) CameraTargetLockTheModel(camera *rl.Camera) {
	con.exitBodyCondition(func() {
		camera.Target = ode.RlGetPostion(con.Body.Position())
	}, func() {
		camera.Target = *con.ModelPosition
	})
}

func (con *ControllerModel) modelUpdate(model *rl.Model, modelAnimtion *rl.ModelAnimation, angle rl.Vector3) {
	con.exitBodyCondition(func() {
		con.Body.SetRotation(ode.RlSetRotation(angle))
		ode.RlSyncTransformTheModel(con.Body.Position(), con.Body.Rotation(),con.Model,con.Scale)
	}, func() {
		model.Transform = rl.MatrixRotate(con.getAxis(con.lastAngle))
	})
	rl.UpdateModelAnimation(*model, *modelAnimtion, con.CounterFrame)
}

func (con *ControllerModel) cameraMoveByModel(cameraPosition *rl.Vector3, axis AxisCamera, speed rl.Vector2) {
	pos := ode.NewVector3()
	con.exitBodyCondition(func() { pos = con.Body.Position() }, func() {})
	if axis == AxisCameraX {
		cameraPosition.X += speed.X
		con.exitBodyCondition(func() {
			pos[0] += float64(speed.Y)
			con.Body.SetPosition(pos)
			con.ModelPosition.X = float32(pos[0])
		}, func() {
			con.ModelPosition.X += speed.Y
		})
	}
	if axis == AxisCameraY {
		cameraPosition.Y += speed.X
		con.exitBodyCondition(func() {
			pos[1] += float64(speed.Y)
			con.Body.SetPosition(pos)
			con.ModelPosition.Y = float32(pos[1])
		}, func() {
			con.ModelPosition.Y += speed.Y
		})
	}
	if axis == AxisCameraZ {
		cameraPosition.Z += speed.X
		con.exitBodyCondition(func() {
			pos[2] += float64(speed.Y)
			con.Body.SetPosition(pos)
			con.ModelPosition.Z = float32(pos[2])
		}, func() {
			con.ModelPosition.Z += speed.Y
		})
	}
}

// axisCamera is axis position and speedModelAndCamera x for camera speed and y for model speed
// you have to set just a angle
func (con *ControllerModel) Move(indexAnimation int32, speedFrame int32, axisCamera AxisCamera, speedModelAndCamera rl.Vector2, angle rl.Vector3) {
	con.indexAnimation = indexAnimation
	con.defaultAnimtionStatus = false
	if con.CounterFrame > con.ModelAnimation[indexAnimation].FrameCount {
		con.CounterFrame = 0
	}
	con.CounterFrame += speedFrame
	con.lastAngle = angle

	con.cameraMoveByModel(con.CameraPosition, axisCamera, speedModelAndCamera)
	con.modelUpdate(con.Model, &con.ModelAnimation[indexAnimation],angle)
}

// inside loop for run DefaultAnimtion
func (con *ControllerModel) DefaultLoopAnimtion(indexAnimation, speed int32) {
	if con.defaultAnimtionStatus && con.OnceStatus != true {
		con.indexAnimation = indexAnimation
		con.CounterFrame += speed
		con.modelUpdate(con.Model, &con.ModelAnimation[indexAnimation], con.lastAngle)
		if con.CounterFrame > con.ModelAnimation[indexAnimation].FrameCount {
			con.CounterFrame = 0
		}
	}
}

// this is for everything that just once run you should by key event
func (con *ControllerModel) Once() {
	if !con.OnceStatus {
		con.CounterFrame = 0
	}
	con.OnceStatus = true
	con.defaultAnimtionStatus = false
}

// inside loop for Once
func (con *ControllerModel) OnceRun(indexAnimation, speed int32) {
	if con.OnceStatus {
		con.indexAnimation = indexAnimation
		con.CounterFrame += speed
		con.modelUpdate(con.Model, &con.ModelAnimation[indexAnimation], con.lastAngle)
		if con.CounterFrame > con.ModelAnimation[indexAnimation].FrameCount {
			con.CounterFrame = 0
			con.OnceStatus = false
			con.defaultAnimtionStatus = true
		}
	}
}

// use inside rl.IsKeyReleased(...)  if you want enable defualtAnimtion
func (con *ControllerModel) EnableDefaultAnimtion() {
	con.defaultAnimtionStatus = true
}

func (con *ControllerModel) GetStatusEnableDefaultAnimtion() bool {
	return con.defaultAnimtionStatus
}

func (con *ControllerModel) FrameCountRest() {
	con.CounterFrame = 0
}

func (con *ControllerModel) DrawSkeletBindPose(scaleSize float32) {
	axis, angle := con.getAxis(con.lastAngle)
	var cubeSize float32 = 0.2
	for i := 0; i < int(con.Model.BoneCount)-1; i++ {
		rotate := rl.Vector3RotateByAxisAngle(rl.GetBindPose(*con.Model, int32(i)).Translation, axis, angle)
		scale := rl.Vector3Scale(rotate, scaleSize)
		translation := rl.Vector3Add(rl.Vector3{con.ModelPosition.X, con.ModelPosition.Y, con.ModelPosition.Z}, scale)
		rl.DrawCube(translation, cubeSize, cubeSize, cubeSize, rl.Red)

		if rl.GetParentBone(*con.Model, i) >= 0 {
			rotate2Par := rl.Vector3RotateByAxisAngle(rl.GetBindPose(*con.Model, rl.GetParentBone(*con.Model, i)).Translation, axis, angle)
			scale2Par := rl.Vector3Scale(rotate2Par, scaleSize)
			translation2Par := rl.Vector3Add(rl.Vector3{con.ModelPosition.X, con.ModelPosition.Y, con.ModelPosition.Z}, scale2Par)
			rl.DrawLine3D(translation, translation2Par, rl.Green)
		}
	}
}

func (con *ControllerModel) DrawSkeletAnimtion(scaleSize float32) {
	axis, angle := con.getAxis(con.lastAngle)
	var cubeSize float32 = 0.2
	for i := 0; i < int(con.Model.BoneCount)-1; i++ {
		rotate := rl.Vector3RotateByAxisAngle(rl.GetBonePose(con.ModelAnimation[con.indexAnimation], con.CounterFrame, int32(i)).Translation, axis, angle)
		scale := rl.Vector3Scale(rotate, scaleSize)
		translation := rl.Vector3Add(rl.Vector3{con.ModelPosition.X, con.ModelPosition.Y, con.ModelPosition.Z}, scale)
		rl.DrawCube(translation, cubeSize, cubeSize, cubeSize, rl.Blue)

		if rl.GetParentBoneAnimtion(&con.ModelAnimation[con.indexAnimation], i) >= 0 {
			rotateParent := rl.Vector3RotateByAxisAngle(rl.GetBonePose(con.ModelAnimation[con.indexAnimation], con.CounterFrame, rl.GetParentBoneAnimtion(&con.ModelAnimation[con.indexAnimation], i)).Translation, axis, angle)
			scaleParent := rl.Vector3Scale(rotateParent, scaleSize)
			translationParent := rl.Vector3Add(rl.Vector3{con.ModelPosition.X, con.ModelPosition.Y, con.ModelPosition.Z}, scaleParent)
			rl.DrawLine3D(translation, translationParent, rl.Green)
		}
	}
}

func (con *ControllerModel) DrawModelControll(rotationAxis rl.Vector3, rotationAngle float32, scale rl.Vector3, color rl.Color) {
	con.exitBodyCondition(func() {
		rl.DrawModelEx(*con.Model, rl.Vector3Zero(), rotationAxis, rotationAngle, scale, color)
	}, func() {
		rl.DrawModelEx(*con.Model, *con.ModelPosition, rotationAxis, rotationAngle, scale, color)
	})
}

func (con *ControllerModel) GetBoneAnimtionTranslation(boneId int, scaleSize float32) rl.Vector3 {
	axis, angle := con.getAxis(con.lastAngle)
	rotate := rl.Vector3RotateByAxisAngle(rl.GetBonePose(con.ModelAnimation[con.indexAnimation], con.CounterFrame, int32(boneId)).Translation, axis, angle)
	scale := rl.Vector3Scale(rotate, scaleSize)
	translation := rl.Vector3Add(rl.Vector3{con.ModelPosition.X, con.ModelPosition.Y, con.ModelPosition.Z}, scale)
	return translation
}

func (con *ControllerModel) GetBoneBindPoseTranslation(boneId int, scaleSize float32) rl.Vector3 {
	axis, angle := con.getAxis(con.lastAngle)
	rotate := rl.Vector3RotateByAxisAngle(rl.GetBindPose(*con.Model, int32(boneId)).Translation, axis, angle)
	scale := rl.Vector3Scale(rotate, scaleSize)
	translation := rl.Vector3Add(rl.Vector3{con.ModelPosition.X, con.ModelPosition.Y, con.ModelPosition.Z}, scale)
	return translation
}

func (con *ControllerModel) JointModelPostionAndDraw(model *rl.Model, boneId int, scaleSizeBone, scale float32, color rl.Color) {
	rl.DrawModel(*model, con.GetBoneAnimtionTranslation(boneId, scaleSizeBone), scale, color)
}

func (con *ControllerModel) exitBodyCondition(exitBodyFunc, nullBodyFunc func()) {
	if con.Body != nil {
		exitBodyFunc()
		return
	}
	nullBodyFunc()
}

func (con *ControllerModel) getAxis(ang rl.Vector3) (rl.Vector3, float32) {
	if ang.X > 0 {
		return rl.NewVector3(1, 0, 0), ang.X
	}
	if ang.Y > 0 {
		return rl.NewVector3(0, 1, 0), ang.Y
	}
	if ang.Z > 0 {
		return rl.NewVector3(0, 0, 1), ang.Z
	}
	return rl.Vector3Zero(), 0
}

// add inside loop RlStep ode
func (con *ControllerModel) SyncTransformTheModel(scale float64) {
	con.exitBodyCondition(func() {
		ode.RlSyncTransformTheModel(con.Body.Position(), con.Body.Rotation(), con.Model, scale)
	}, func() {})
}

func (con *ControllerModel) AddCollisionByBody(body *ode.Body, kinematic bool, boneId int, scaleSize float32) {
	if kinematic {
		body.SetKinematic(true)
	}
	body.SetAngularVelocity(ode.NewVector3(0, 0, 0))
	postionBoneAnimtion := con.GetBoneAnimtionTranslation(boneId, scaleSize)
	body.SetPosition(ode.NewVector3(float64(postionBoneAnimtion.X), float64(postionBoneAnimtion.Y), float64(postionBoneAnimtion.Z)))
}

func (con *ControllerModel) Unload() {
	rl.UnloadModel(*con.Model)
	rl.UnloadModelAnimations(con.ModelAnimation)
}
