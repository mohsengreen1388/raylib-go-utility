package utility

import (
	"unsafe"

	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

type Mask struct {
	Shader             *rl.Shader
	updateAnimationLoc int32
}

func (ma *Mask)Init(){
	ma.configShader()
	
	ma.Shader.UpdateLocation(rl.ShaderLocMapEmission, rl.GetShaderLocation(*ma.Shader, "mask"))
	ma.updateAnimationLoc = rl.GetShaderLocation(*ma.Shader, "frame")
}

func (ma *Mask) SetMapDiffuse(materials []*rl.Material, textures []*rl.Texture2D) {
	for index, material := range materials {
		material.GetMap(rl.MapDiffuse).Texture = *textures[index]
	}
}

func (ma *Mask) SetMapEmission(materials []*rl.Material, texturesMask []*rl.Texture2D) {
	for index, material := range materials {
		material.GetMap(rl.MapEmission).Texture = *texturesMask[index]
		material.Shader = *ma.Shader
	}
}

// take inside loop
func (ma *Mask) UpdateAnimation(x int) {
	rl.SetShaderValue(*ma.Shader,ma.updateAnimationLoc, unsafe.Slice((*float32)(unsafe.Pointer(&x)), 4), rl.ShaderUniformInt)
}

func (ma *Mask) configShader() {
	if checkPlatformIsMobail() {
		ma.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Mobail[""], ShaderFile.Mobail["maskFs"]))
	} else {
		ma.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Desktop[""], ShaderFile.Desktop["maskFs"]))
	}
}

func (ma *Mask) Unload() {
	rl.UnloadShader(*ma.Shader)
}