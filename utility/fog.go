package utility

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

type Fog struct {
	Shader        *rl.Shader
	combineStatus bool
	FogDensity    float32
	fogDensityLoc int32
}

// fogDensity 0.0 -> 0.16
func (fo *Fog) Init(ambient rl.Color, fogDensity float32) {
	if !fo.combineStatus {
		fo.configShader()
	}
	fo.updateUniformInTheShader(ambient, fogDensity)
}

func (fo *Fog) UpdateFogDensity(fogDensity float32) {
	rl.SetShaderValue(*fo.Shader, fo.fogDensityLoc, []float32{float32(fogDensity)}, rl.ShaderUniformFloat)
}

// take inside loop
func (fo *Fog) Update(cameraPos rl.Vector3) {
	rl.SetShaderValue(*fo.Shader, rl.GetShaderLocation(*fo.Shader, "viewPos"), []float32{cameraPos.X}, rl.ShaderUniformVec3)
}

func (fo *Fog) AddTheMaterialToFog(models []*rl.Material, textures []*rl.Texture2D) {
	for index, material := range models {
		material.GetMap(rl.MapDiffuse).Texture = *textures[index]
	}
}

func (fo *Fog) updateUniformInTheShader(ambient rl.Color, fogDensity float32) {

	fo.Shader.UpdateLocation(rl.RL_SHADER_LOC_MATRIX_MODEL, rl.GetShaderLocation(*fo.Shader, "matModel"))
	fo.Shader.UpdateLocation(rl.RL_SHADER_LOC_VECTOR_VIEW, rl.GetShaderLocation(*fo.Shader, "viewPos"))

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(*fo.Shader, "ambient")
	rl.SetShaderValue(*fo.Shader, ambientLoc, []float32{float32(ambient.R) / 255, float32(ambient.G) / 255, float32(ambient.B) / 255, float32(ambient.A) / 255}, rl.ShaderUniformVec4)

	fo.FogDensity = fogDensity
	fo.fogDensityLoc = rl.GetShaderLocation(*fo.Shader, "fogDensity")
	rl.SetShaderValue(*fo.Shader, fo.fogDensityLoc, []float32{float32(fogDensity)}, rl.ShaderUniformFloat)
}

func (fo *Fog) configShader() {
	if checkPlatformIsMobail() {
		fo.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Mobail["fogVs"], ShaderFile.Mobail["fogFs"]))
	} else {
		fo.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Desktop["fogVs"], ShaderFile.Desktop["fogFs"]))
	}
}

// exce before init or set manually
func (fo *Fog) SetCombineShader(CombineShader *rl.Shader) {
	fo.combineStatus = true
	fo.Shader = CombineShader
}

func (fo *Fog) Unload() {
	rl.UnloadShader(*fo.Shader)
}
