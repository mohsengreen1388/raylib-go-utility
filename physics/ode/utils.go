package ode

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	"math"
)

type CGOMap struct {
	index    int
	pointers map[int]interface{}
}

func NewCGOMap() CGOMap {
	var m CGOMap
	m.pointers = make(map[int]interface{})
	return m

}

func (m CGOMap) Get(index int) interface{} {
	p := m.pointers[index]
	if p == nil {
		panic("couldn't retrieve the pointer")
	}
	return p
}

func (m CGOMap) Set(p interface{}) int {
	m.index += 1
	m.pointers[m.index] = p
	return m.index
}

func (m CGOMap) Delete(index int) {
	delete(m.pointers, index)
}

// the contact have to make by max size -> if you max is 8  contact := make([]*ode.Contact,8)
//
// you have to set mode,Mu
//
// order surfaceParameters Mu2,Bounce,BounceVel,SoftCfm,SoftErp,Slip1,Slip2,Rho,Rho2,RhoN,Motion1,Motion2,MotionN
func RlCollideCallBack(world *World, max int, jointGroup JointGroup, contact []*Contact, mode, Mu int, surfaceParameters ...float64) NearCallback {

	return func(data interface{}, obj1, obj2 Geom) {

		body1, body2 := obj1.Body(), obj2.Body()
		if body1 != 0 && body2 != 0 && body1.ConnectedExcluding(body2, 0) {
			return
		}

		for i := 0; i < max; i++ {
			contact[i] = NewContact()
			contact[i].Surface.Mode = mode
			contact[i].Surface.Mu = math.Inf(Mu)
			contact[i].Surface.Mu2 = 0
			contact[i].Surface.Bounce = 0.1
			contact[i].Surface.BounceVel = 0.1
			contact[i].Surface.SoftCfm = 0.01
			contact[i].Surface.SoftErp = 0
			contact[i].Surface.Slip1 = 0
			contact[i].Surface.Slip2 = 0
			contact[i].Surface.Rho = 0
			contact[i].Surface.Rho2 = 0
			contact[i].Surface.RhoN = 0
			contact[i].Surface.Motion1 = 0
			contact[i].Surface.Motion2 = 0
			contact[i].Surface.MotionN = 0
		}

		contactGeom := obj1.Collide(obj2, uint16(max), 0)

		if len(contactGeom) > 0 {
			for i := 0; i < len(contactGeom); i++ {
				contact[i].Geom = contactGeom[i]
				contactJoint := world.NewContactJoint(jointGroup, contact[i])
				contactJoint.Attach(body1, body2)
			}
		}
	}
}

// AutoDisableFlag 0 | 1
// plane ode.V4(0, 1, 0, 0) for ground
func RlInit(maxJoints int, grav Vector3, autoDisable bool, autoDisableFlag int, plane Vector4) (*World, Space, *JointGroup) {

	Init(0, AllAFlag)
	world := NewWorld()
	world.SetAutoDisable(autoDisable)
	world.WorldSetAutoDisableFlag(autoDisableFlag)
	world.SetGravity(grav)
	world.SetCFM(0.01)
	world.SetERP(0.2)

	space := NilSpace().NewHashSpace()
	space.NewPlane(plane)

	jointGroup := NewJointGroup(maxJoints)

	return &world, space, &jointGroup
}

// inside loop
// stepTime -> 60
func RlStep(world *World, space Space, jointGroup *JointGroup, stepTime float64, nearCallback NearCallback, setOption func()) {
	space.Collide(0, nearCallback)
	world.QuickStep(1.0 / stepTime)
	jointGroup.Empty()
	setOption()
}

// RotateXYZ for ode gem
func RlSetRotation(ang rl.Vector3) Matrix3 {
	matrix := rl.MatrixRotateXYZ(ang)
	matOde := NewMatrix3()

	matOde[0][0] = float64(matrix.M0)
	matOde[1][0] = float64(matrix.M1)
	matOde[2][0] = float64(matrix.M2)
	matOde[0][1] = float64(matrix.M4)
	matOde[1][1] = float64(matrix.M5)
	matOde[2][1] = float64(matrix.M6)
	matOde[0][2] = float64(matrix.M8)
	matOde[1][2] = float64(matrix.M9)
	matOde[2][2] = float64(matrix.M10)

	return matOde
}

// sync Geom  rotate and postion the model
func RlSyncTransformTheModel(pos Vector3, R Matrix3, model *rl.Model,scale float64) {
	model.Transform.M0 = float32(R[0][0])
	model.Transform.M1 = float32(R[1][0])
	model.Transform.M2 = float32(R[2][0])
	model.Transform.M3 = 0
	model.Transform.M4 = float32(R[0][1])
	model.Transform.M5 = float32(R[1][1])
	model.Transform.M6 = float32(R[2][1])
	model.Transform.M7 = 0
	model.Transform.M8 = float32(R[0][2])
	model.Transform.M9 = float32(R[1][2])
	model.Transform.M10 = float32(R[2][2])
	model.Transform.M11 = 0
	model.Transform.M12 = float32(pos[0]/scale)
	model.Transform.M13 = float32(pos[1]/scale)
	model.Transform.M14 = float32(pos[2]/scale)
	model.Transform.M15 = 1
}

func RlSetRotationTheBodyAndModel(body *Body, model *rl.Model, ang rl.Vector3) {
	body.SetRotation(RlSetRotation(ang))
	model.Transform = rl.MatrixRotateXYZ(rl.Vector3{ang.X, ang.Y, ang.Z})
}

func RlCheckCollide(geom Geom, geom2 Geom, maxContacts uint16, flags int) (bool, []ContactGeom) {
	ContactGeom := geom.Collide(geom2, maxContacts, flags)
	if ContactGeom != nil {
		return true, ContactGeom
	}
	return false, ContactGeom
}

func RlNewBox(world *World, space Space, boxLens, pos, lensMass Vector3, massTotal float64, autoDisble, isGravityEnabled bool) (*Body, *Box, *Mass) {
	body := world.NewBody()
	box := space.NewBox(boxLens)
	mass := NewMass()
	mass.SetBoxTotal(massTotal, NewVector3(lensMass[0]/2, lensMass[1]/2, lensMass[2]/2))
	box.SetBody(body)
	body.SetMass(mass)
	body.SetPosition(pos)
	body.SetAutoDisable(autoDisble)
	body.SetGravityEnabled(isGravityEnabled)

	return &body, &box, mass
}

func RlNewSphere(world *World, space Space, radius float64, pos Vector3, massTotal float64, autoDisble, isGravityEnabled bool) (*Body, *Sphere, *Mass) {
	body := world.NewBody()
	Sphere := space.NewSphere(radius)
	mass := NewMass()
	mass.SetSphereTotal(massTotal, radius/2)
	Sphere.SetBody(body)
	body.SetMass(mass)
	body.SetPosition(pos)
	body.SetAutoDisable(autoDisble)
	body.SetGravityEnabled(isGravityEnabled)

	return &body, &Sphere, mass
}

func RlNewCapsule(world *World, space Space, radius, length float64, directionMass int, pos Vector3, massTotal float64, autoDisble, isGravityEnabled bool) (*Body, *Capsule, *Mass) {
	body := world.NewBody()
	capsule := space.NewCapsule(radius, length)
	mass := NewMass()
	mass.SetCapsuleTotal(massTotal, directionMass, radius/2, length)
	capsule.SetBody(body)
	body.SetMass(mass)
	body.SetPosition(pos)
	body.SetAutoDisable(autoDisble)
	body.SetGravityEnabled(isGravityEnabled)

	return &body, &capsule, mass
}

func RlNewCylinder(world *World, space Space, radius, length float64, directionMass int, pos Vector3, massTotal float64, autoDisble, isGravityEnabled bool) (*Body, *Cylinder, *Mass) {
	body := world.NewBody()
	cylinder := space.NewCylinder(radius, length)
	mass := NewMass()
	mass.SetCylinderTotal(massTotal, directionMass, radius/2, length)
	cylinder.SetBody(body)
	body.SetMass(mass)
	body.SetPosition(pos)
	body.SetAutoDisable(autoDisble)
	body.SetGravityEnabled(isGravityEnabled)

	return &body, &cylinder, mass
}

func RlNewRay(space Space, length float64, pos, dir Vector3, geomEnabled bool) *Ray {
	ray := space.NewRay(length)
	ray.SetPosDir(pos, dir)
	ray.SetPosition(pos)
	ray.SetEnabled(geomEnabled)

	return &ray
}

func RlNewTriMeshGeom(space Space, mesh *rl.Mesh, scale float32) (*TriMeshData, *TriMesh) {
	trmisdata := RlGeomTriMeshDataBuildSingle(mesh, scale)
	triMesh := RlNewTriMesh(space, &trmisdata)
	return &trmisdata, &triMesh
}

func RlNewTriMeshGeomAnm(space Space, mesh *rl.Mesh) (*TriMeshData, *TriMesh) {
	trmisdata := RlGeomTriMeshDataBuildSingleAnm(mesh)
	triMesh := RlNewTriMesh(space, &trmisdata)
	return &trmisdata, &triMesh
}

func RlMeshToDynamic(world *World, space Space, trimesh *TriMesh, trmisdata *TriMeshData, totalMass float64, pos Vector3) (*Body, Mass) {
	body := world.NewBody()
	trimesh.SetBody(body)
	mass := NewMass()
	mass.SetTriMeshTotal(totalMass, *trimesh)
	body.FirstGeom().SetPosition(NewVector3(-mass.Center[0], -mass.Center[1], -mass.Center[2]))
	mass.Translate(NewVector3(-mass.Center[0], -mass.Center[1], -mass.Center[2]))
	trmisdata.Preprocess()
	body.SetPosition(pos)

	return &body, *mass
}

func RlGetPostion(body Vector3) rl.Vector3 {
	odeVec3 := body
	return rl.NewVector3(float32(odeVec3[0]), float32(odeVec3[1]), float32(odeVec3[2]))
}

type JointTypeDraw interface {
	Body(index int) Body
	NumBodies() int
}

func drawAnchor(joint JointTypeDraw, anchor Vector3, anchor2 Vector3) {
	rl.DrawSphere(rl.Vector3{float32(anchor[0]), float32(anchor[1]), float32(anchor[2])}, 0.2, rl.Red)
	for i := 0; i < joint.NumBodies(); i++ {
		if i == 0 {
			rl.DrawLine3D(rl.Vector3{float32(joint.Body(0).Position()[0]), float32(joint.Body(0).Position()[1]), float32(joint.Body(0).Position()[2])}, rl.Vector3{float32(anchor[0]), float32(anchor[1]), float32(anchor[2])}, rl.Blue)
		}
		if i == 1 {
			rl.DrawLine3D(rl.Vector3{float32(joint.Body(1).Position()[0]), float32(joint.Body(1).Position()[1]), float32(joint.Body(1).Position()[2])}, rl.Vector3{float32(anchor2[0]), float32(anchor2[1]), float32(anchor2[2])}, rl.Blue)
		}
	}
	rl.DrawSphere(rl.Vector3{float32(anchor2[0]), float32(anchor2[1]), float32(anchor2[2])}, 0.2, rl.Blue)
}

func RlDrawAnchorBall(ball *BallJoint, anchor Vector3, anchor2 Vector3) {
	drawAnchor(ball, anchor, anchor2)
}

func RlDrawAnchorHing(hing *HingeJoint, anchor Vector3, anchor2 Vector3) {
	drawAnchor(hing, anchor, anchor2)
}

func RlDrawAnchorHing2(hing2 *Hinge2Joint, anchor Vector3, anchor2 Vector3) {
	drawAnchor(hing2, anchor, anchor2)
}

func DrawSlider(slide *SliderJoint) {
	for i := 0; i < slide.NumBodies(); i++ {
		if i == 0 {
			rl.DrawLine3D(rl.Vector3{float32(slide.Body(0).Position()[0]), float32(slide.Body(0).Position()[1]), float32(slide.Body(0).Position()[2])}, rl.Vector3{float32(slide.Body(1).Position()[0]), float32(slide.Body(1).Position()[1]), float32(slide.Body(1).Position()[2])}, rl.Blue)
		}
		if i == 1 {
			rl.DrawLine3D(rl.Vector3{float32(slide.Body(1).Position()[0]), float32(slide.Body(1).Position()[1]), float32(slide.Body(1).Position()[2])}, rl.Vector3{float32(slide.Body(0).Position()[0]), float32(slide.Body(0).Position()[1]), float32(slide.Body(0).Position()[2])}, rl.Blue)
		}
	}
}

func RlDrawAnchorUniversal(universa *UniversalJoint, anchor Vector3, anchor2 Vector3) {
	drawAnchor(universa, anchor, anchor2)
}

func RlDrawAnchorPiston(piston *PistonJoint, anchor Vector3, anchor2 Vector3) {
	for i := 0; i < piston.NumBodies(); i++ {
		if i == 0 {
			rl.DrawLine3D(rl.Vector3{float32(piston.Body(0).Position()[0]), float32(piston.Body(0).Position()[1]), float32(piston.Body(0).Position()[2])}, rl.Vector3{float32(piston.Body(1).Position()[0]), float32(piston.Body(1).Position()[1]), float32(piston.Body(1).Position()[2])}, rl.Blue)
		}
		if i == 1 {
			rl.DrawLine3D(rl.Vector3{float32(piston.Body(1).Position()[0]), float32(piston.Body(1).Position()[1]), float32(piston.Body(1).Position()[2])}, rl.Vector3{float32(piston.Body(0).Position()[0]), float32(piston.Body(0).Position()[1]), float32(piston.Body(0).Position()[2])}, rl.Blue)
		}
	}
}

func RlDrawTriMesh(tri *TriMesh, color rl.Color) {
	for i := 0; i < tri.TriangleCount(); i++ {
		v, v1, v2 := tri.Triangle(i)
		rl.DrawLine3D(rl.Vector3{float32(v[0]), float32(v[1]), float32(v[2])}, rl.Vector3{float32(v1[0]), float32(v1[1]), float32(v1[2])}, color)
		rl.DrawLine3D(rl.Vector3{float32(v1[0]), float32(v1[1]), float32(v1[2])}, rl.Vector3{float32(v2[0]), float32(v2[1]), float32(v2[2])}, color)
	}
}

func RlDrawRay(pos, dir Vector3, length float64) {
	dirVec3 := RlGetPostion(dir)
	posVec3 := RlGetPostion(pos)
	var xl, yl, zl float64

	if dirVec3.X == 1 {
		xl = length
	}
	if dirVec3.Y == 1 {
		yl = length
	}
	if dirVec3.Z == 1 {
		zl = length
	}
	rl.DrawLine3D(posVec3, rl.Vector3{posVec3.X + float32(xl), posVec3.Y + float32(yl), posVec3.Z + float32(zl)}, rl.Green)
}

// use rl.LoadModelFromMesh(rl.GenMeshCylinder(radius,height,slices))
func RlDrawCapsule(meshCylinderModel *rl.Model, body *Body,scale float64,color rl.Color) {
	RlSyncTransformTheModel(body.Position(), body.Rotation(), meshCylinderModel,scale)
	rl.DrawModelWires(*meshCylinderModel, rl.Vector3Zero(), 1, color)
}

// use rl.LoadModelFromMesh(rl.GenMeshCylinder(radius,height,slices))
func RlDrawCylinder(meshCylinderModel *rl.Model, body *Body,scale float64,color rl.Color) {
	RlSyncTransformTheModel(body.Position(), body.Rotation(),meshCylinderModel,scale)
	rl.DrawModelWires(*meshCylinderModel, rl.Vector3Zero(), 1, color)
}

// rl.LoadModelFromMesh(rl.GenMeshCube(1, 1, 1))
func RlDrawBox(meshCubeModel *rl.Model, body *Body, boxLens rl.Vector3,scale float64,color rl.Color) {
	RlSyncTransformTheModel(body.Position(), body.Rotation(), meshCubeModel,scale)
	rl.DrawModelWires(*meshCubeModel, rl.Vector3Zero(), 1, color)
}

func RlNewBallJoint(world *World, body1, body2 *Body, jointGroup JointGroup, anchor Vector3) *BallJoint {
	ball := world.NewBallJoint(jointGroup)
	ball.Attach(*body1, *body2)
	ball.SetAnchor(anchor)

	return &ball
}

func RlNewHingJoint(world *World, body1, body2 *Body, jointGroup JointGroup, anchor Vector3, axis Vector3) *HingeJoint {
	hing := world.NewHingeJoint(jointGroup)
	hing.Attach(*body1, *body2)
	hing.SetAnchor(anchor)
	hing.SetAxis(axis)

	return &hing
}

func RlNewHing2Joint(world *World, body1, body2 *Body, jointGroup JointGroup, anchor, axis1, axis2 Vector3) *Hinge2Joint {
	hing2 := world.NewHinge2Joint(jointGroup)
	hing2.Attach(*body1, *body2)
	hing2.SetAnchor(anchor)
	hing2.SetAxis1(axis1)
	hing2.SetAxis2(axis2)

	return &hing2
}

func RlNewSliderJoint(world *World, body1, body2 *Body, jointGroup JointGroup, axis Vector3) *SliderJoint {
	slider := world.NewSliderJoint(jointGroup)
	slider.Attach(*body1, *body2)
	slider.SetAxis(axis)

	return &slider
}

func RlNewPistonJoint(world *World, body1, body2 *Body, jointGroup JointGroup, anchor, axis Vector3) *PistonJoint {
	piston := world.NewPistonJoint(jointGroup)
	piston.Attach(*body1, *body2)
	piston.SetAnchor(anchor)
	piston.SetAxis(axis)

	return &piston
}

func RlNewUniversalJoint(world *World, body1, body2 *Body, jointGroup JointGroup, anchor, axis1 Vector3,axis2 Vector3) *UniversalJoint {
	universal := world.NewUniversalJoint(jointGroup)
	universal.Attach(*body1, *body2)
	universal.SetAnchor(anchor)
	universal.SetAxis1(axis1)
	universal.SetAxis2(axis2)

	return &universal
}

func RlDrawAnchorTransmission(joint *TransmissionJoint, contactPoint bool) {
	contactPoint1 := rl.Vector3{}
	contactPoint2 := rl.Vector3{}
	anchor1 := RlGetPostion(joint.Anchor1())
	anchor2 := RlGetPostion(joint.Anchor2())

	if contactPoint {
		contactPoint1 = RlGetPostion(joint.ContactPoint1())
		contactPoint2 = RlGetPostion(joint.ContactPoint2())
	}
	rl.DrawSphere(rl.Vector3{float32(contactPoint1.X), float32(contactPoint1.Y), float32(contactPoint1.Z)}, 0.1, rl.Blue)
	rl.DrawSphere(rl.Vector3{float32(contactPoint2.X), float32(contactPoint2.Y), float32(contactPoint2.Z)}, 0.1, rl.Green)
	rl.DrawLine3D(rl.Vector3{anchor1.X, anchor1.Y, anchor1.Z}, rl.Vector3{anchor2.X, anchor2.Y, anchor2.Z}, rl.Red)
}

func RlNewTransmission(world *World, jointGroup JointGroup, transmissionMode int) *TransmissionJoint {
	transmission := world.NewTransmissionJoint(jointGroup)
	transmission.SetMode(transmissionMode)

	return &transmission
}

// An angular motor (AMotor) allows the relative angular velocities of two bodies to be controlled. The angular velocity can be controlled on up to three axes.
// allowing torque motors and stops to be set for rotation about those axes (see the "Stops and motor parameters" section below). This is mainly useful in conjunction will ball joints (which do not constrain the angular degrees of freedom at all), but it can be used in any situation where angular control is needed.
// To use an AMotor with a ball joint, simply attach it to the same two bodies that the ball joint is attached to.
func RlNewAngularMotor(world *World, jointGroup JointGroup, numAxes int) *AMotorJoint {
	angularMotor := world.NewAMotorJoint(jointGroup)
	angularMotor.SetNumAxes(numAxes)

	return &angularMotor
}

// A linear motor (LMotor) allows the relative linear velocities of two bodies to be controlled.
// The linear velocity can be controlled on up to three axes
// allowing torque motors and stops to be set for translation along those axes (see the "Stops and motor parameters" section below).
func RlNewLinearMotor(world *World, jointGroup JointGroup, numAxes int) *LMotorJoint {
	linearMotor := world.NewLMotorJoint(jointGroup)
	linearMotor.SetNumAxes(numAxes)

	return &linearMotor
}

func RlClose(world *World, space Space, jointGrop *JointGroup) {
	jointGrop.Empty()
	jointGrop.Destroy()
	space.Destroy()
	world.Destroy()
	Close()
}

func RlDrawDebug(f func()) {}

// for clean code
func RlBeginConfig() {}
func RlEndConfig()   {}

// for clean code you have to inside BeginMode3D(cam)
func RlBeginDrawingDebug() {}
func RlEndDrawingDebug()   {}
