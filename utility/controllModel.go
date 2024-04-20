package utility

import (
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
	ModelAnimation        []rl.ModelAnimation
	ModelPosition         *rl.Vector3
	CameraPosition        *rl.Vector3
	defaultAnimtionStatus bool
	OnceStatus            bool
	lastAngle             float32
	CounterFrame          int32
}

func (con *ControllerModel) Init(model *rl.Model, modelAnimation []rl.ModelAnimation, modelPosition, cameraPosition *rl.Vector3, enableDefaultAnimtion bool) {
	con.Model = model
	con.ModelAnimation = modelAnimation
	con.CameraPosition = cameraPosition
	con.defaultAnimtionStatus = enableDefaultAnimtion
	con.ModelPosition = modelPosition
}

func (con *ControllerModel) CameraTargetLockTheModel(camera *rl.Camera) {
	camera.Target = *con.ModelPosition
}

func (con *ControllerModel) modelUpdate(model *rl.Model, modelAnimtion *rl.ModelAnimation, axis rl.Vector3, angle float32) {
	model.Transform = rl.MatrixRotate(axis, angle)
	rl.UpdateModelAnimation(*model, *modelAnimtion, con.CounterFrame)
}

func (con *ControllerModel) cameraMoveByModel(cameraPosition, modelPosition *rl.Vector3, axis AxisCamera, speed rl.Vector2) {
	if axis == AxisCameraX {
		cameraPosition.X += speed.X
		modelPosition.X += speed.Y
	}
	if axis == AxisCameraY {
		cameraPosition.Y += speed.X
		modelPosition.Y += speed.Y
	}
	if axis == AxisCameraZ {
		cameraPosition.Z += speed.X
		modelPosition.Z += speed.Y
	}
}

// axisCamera is axis position and speedModelAndCamera x for camera speed and y for model speed
func (con *ControllerModel) Move(indexAnimation int32, speedFrame int32, axisCamera AxisCamera, speedModelAndCamera rl.Vector2, axis rl.Vector3, angle float32) {
	con.defaultAnimtionStatus = false
	if con.CounterFrame > con.ModelAnimation[indexAnimation].FrameCount {
		con.CounterFrame = 0
	}
	con.CounterFrame += speedFrame
	con.lastAngle = angle

	con.modelUpdate(con.Model, &con.ModelAnimation[indexAnimation], axis, angle)
	con.cameraMoveByModel(con.CameraPosition, con.ModelPosition, axisCamera, speedModelAndCamera)
}

// inside loop for run DefaultAnimtion
func (con *ControllerModel) DefaultLoopAnimtion(indexAnimation, speed int32, axis rl.Vector3) {
	if con.defaultAnimtionStatus {
		con.CounterFrame += speed
		con.modelUpdate(con.Model, &con.ModelAnimation[indexAnimation], axis, con.lastAngle)
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
func (con *ControllerModel) OnceRun(indexAnimation, speed int32, axis rl.Vector3) {
	if con.OnceStatus {
		con.CounterFrame += speed
		con.modelUpdate(con.Model, &con.ModelAnimation[indexAnimation], axis, float32(con.lastAngle))
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

func (con *ControllerModel) Unload() {
	rl.UnloadModel(*con.Model)
	rl.UnloadModelAnimations(con.ModelAnimation)
}
