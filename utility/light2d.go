package utility

import (
	"image/color"

	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

type Lights2d struct {
	Shader           *rl.Shader
	Postion          rl.Vector2
	LightRadius      float32
	AmbientStrong    float32
	AmbientColor     rl.Color
	posLoc           int32
	lightRadiusLoc   int32
	ambientColorLoc  int32
	ambientStrongLoc int32
}

func (li *Lights2d) SpotLightInit(lightRadius, ambientStrong float32, ambientColor rl.Color) {
	li.configShader()
	li.LightRadius = lightRadius
	li.AmbientStrong = ambientStrong
	li.AmbientColor = ambientColor

	li.posLoc = rl.GetShaderLocation(*li.Shader, "lightPos")
	li.lightRadiusLoc = rl.GetShaderLocation(*li.Shader, "lightRadius")
	li.ambientColorLoc = rl.GetShaderLocation(*li.Shader, "ambientColor")
	li.ambientStrongLoc = rl.GetShaderLocation(*li.Shader, "ambientStrong")

	rl.SetShaderValue(*li.Shader, li.lightRadiusLoc, []float32{lightRadius}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*li.Shader, li.ambientColorLoc, []float32{float32(ambientColor.R) / 255, float32(ambientColor.G) / 255, float32(ambientColor.B) / 255}, rl.ShaderUniformVec3)
	rl.SetShaderValue(*li.Shader, li.ambientStrongLoc, []float32{ambientStrong}, rl.ShaderUniformFloat)
}

func (li *Lights2d) UpdatePosition(pos rl.Vector2){
	rl.SetShaderValue(*li.Shader,li.posLoc, []float32{pos.X,pos.Y}, rl.ShaderUniformVec2)
}

func (li *Lights2d) UpdateOptions(lightRadius, ambientStrong float32, ambientColor rl.Color){
	li.LightRadius = lightRadius
	li.AmbientStrong = ambientStrong
	li.AmbientColor = ambientColor
	
	rl.SetShaderValue(*li.Shader, li.lightRadiusLoc, []float32{lightRadius}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*li.Shader, li.ambientColorLoc, []float32{float32(ambientColor.R) / 255, float32(ambientColor.G) / 255, float32(ambientColor.B) / 255}, rl.ShaderUniformVec3)
	rl.SetShaderValue(*li.Shader, li.ambientStrongLoc, []float32{ambientStrong}, rl.ShaderUniformFloat)
}

func (li *Lights2d) SpotLightDraw() {
	rl.BeginShaderMode(*li.Shader)
		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Black)
	rl.EndShaderMode()
}


// you have to use Fade -> []color.RGBA{rl.Fade(rl.Yellow,0.3)
func (li *Lights2d) PointLight(texture *rl.Texture2D, position rl.Vector2, scale float32, color []color.RGBA) {
	rl.BeginBlendMode(rl.BlendAlpha)
	for i := 0; i < len(color); i++ {
		rl.DrawTextureEx(*texture, position, 0, scale, color[i])
	}
	rl.EndBlendMode()
}

func (li *Lights2d) configShader() {
	if checkPlatformIsMobail() {
		li.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Mobail[""], ShaderFile.Mobail["light2dFs"]))
	} else {
		li.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Desktop[""], ShaderFile.Desktop["light2dFs"]))
	}
}

func (li *Lights2d)Unload(){
	rl.UnloadShader(*li.Shader)
}