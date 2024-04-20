package utility

import (
	"image/color"

	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

type Particle struct {
	Position           rl.Vector3
	PositionInTheWorld rl.Vector3
	Speed              rl.Vector3
	Color              rl.Color
	Lifetime           float32
	Particles          []Particle
	MaxParticles       int32
	CircleParicleSpeed float32
	Oneshot            bool
}

//The color Random set manually Color Option unless not show anything if you choose false.
func (pa *Particle) InitializeParticles(max int, position, speed rl.Vector3, lifetimeMin, lifetimeMax int32,colorRandom bool) {
	pa.Particles = make([]Particle, max)
	pa.MaxParticles = int32(max)
	pa.PositionInTheWorld = position

	for i := 0; i < max; i++ {
		pa.Particles[i].Position = pa.PositionInTheWorld
		pa.Particles[i].Speed = speed
		pa.Particles[i].Color = pa.colorRandomCheck(colorRandom)
		pa.Particles[i].Lifetime = float32(rl.GetRandomValue(lifetimeMin, lifetimeMax))
	}
}

func (pa *Particle)colorRandomCheck(status bool)color.RGBA{
	if status{
		return	rl.NewColor(255, uint8(rl.GetRandomValue(100, 255)), 0, uint8(rl.GetRandomValue(200, 255)))
	}

	return pa.Color
}

func (pa *Particle) UpdateParticles(speed rl.Vector3, lifetimeMin, lifetimeMax int32) {

	for i := 0; i < int(pa.MaxParticles); i++ {
		if pa.Particles[i].Lifetime > 0 {
			pa.Particles[i].Position.X += pa.Particles[i].Speed.X
			pa.Particles[i].Position.Y += pa.Particles[i].Speed.Y
			pa.Particles[i].Position.Z += pa.Particles[i].Speed.Z

			pa.Particles[i].Lifetime--

			if pa.Particles[i].Lifetime <= 0 {
				pa.Particles[i].Position = pa.PositionInTheWorld
				pa.Particles[i].Speed = speed
				if !pa.Oneshot {
					pa.Particles[i].Lifetime = float32(rl.GetRandomValue(lifetimeMin, lifetimeMax))
				}
			}

		}
	}
}

func (pa *Particle) DrawParticles(model *rl.Model, texture *rl.Texture2D, scale rl.Vector3) {
	rl.DisableDepthTest()
	rl.DisableBackfaceCulling()
	model.Materials[0].Maps.Texture = *texture
	for i := 0; i < int(pa.MaxParticles); i++ {
		if pa.Particles[i].Lifetime > 0 {
			rl.BeginBlendMode(rl.BlendAdditive)
			rl.DrawModelEx(*model, rl.Vector3{pa.Particles[i].Position.X, pa.Particles[i].Position.Y, pa.Particles[i].Position.Z}, rl.Vector3{1, 0, 0}, 90, scale, pa.Particles[i].Color)
			rl.EndBlendMode()
		}
	}
	defer rl.EnableDepthTest()
	defer rl.EnableBackfaceCulling()
}

func (pa *Particle) DrawCircleParicle(speed float32, model *rl.Model, texture *rl.Texture2D, position, rotationAxis, scale rl.Vector3) {
	rl.DisableDepthTest()
	rl.DisableBackfaceCulling()
	pa.CircleParicleSpeed += speed
	moveForeachOther := 0
	model.Materials[0].Maps.Texture = *texture
	if pa.CircleParicleSpeed >= 10000 {
		pa.CircleParicleSpeed = 100
	}

	for i := 0; i < int(pa.MaxParticles); i++ {
		moveForeachOther++
		if pa.Particles[i].Lifetime > 0 {
			rl.BeginBlendMode(rl.BlendAdditive)
			rl.DrawModelEx(*model, position, rotationAxis, float32(pa.CircleParicleSpeed)*float32(moveForeachOther), scale, pa.Particles[i].Color)
			rl.EndBlendMode()
		}
	}
	defer rl.EnableDepthTest()
	defer rl.EnableBackfaceCulling()
}

func (pa *Particle) DrawParticles2dCircle(radius float32){
	for i := 0; i < int(pa.MaxParticles); i++ {
		if pa.Particles[i].Lifetime > 0 {
			rl.DrawCircle(int32(pa.Particles[i].Position.X),int32(pa.Particles[i].Position.Y),radius,pa.Particles[i].Color)
		}
	}
}

func (pa *Particle) DrawParticles2dTexture(texture *rl.Texture2D){
	for i := 0; i < int(pa.MaxParticles); i++ {
		if pa.Particles[i].Lifetime > 0 {
			rl.DrawTexture(*texture,int32(pa.Particles[i].Position.X),int32(pa.Particles[i].Position.Y),pa.Particles[i].Color)
		}
	}
}

func (pa *Particle) DrawParticlesBilborid(cam *rl.Camera,texture *rl.Texture2D,size float32){
	rl.DisableDepthTest()
	for i := 0; i < int(pa.MaxParticles); i++ {
		if pa.Particles[i].Lifetime > 0 {
			rl.BeginBlendMode(rl.BlendAdditive)
			rl.DrawBillboard(*cam,*texture,rl.Vector3{pa.Particles[i].Position.X,pa.Particles[i].Position.Y,pa.Particles[i].Position.Z},size,pa.Particles[i].Color)
			rl.EndBlendMode()
		}
	}
	defer rl.EnableDepthTest()
}