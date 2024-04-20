package utility

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	"unsafe"
)

type SkyBox struct {
	cube *rl.Model
}

func (sky *SkyBox) Init(doGamma, vflipped bool, file string) {
	mesh := rl.GenMeshCube(2.0, 2.0, 2.0)
	SkyBox := rl.LoadModelFromMesh(mesh)
	sky.cube = &SkyBox
	Gamma := 0
	flipped := 0

	SkyBox.Materials[0].Shader = *sky.configShader()

	if doGamma {
		Gamma = 1
	}
	if vflipped {
		flipped = 1
	}

	rl.SetShaderValue(SkyBox.Materials[0].Shader, rl.GetShaderLocation(SkyBox.Materials[0].Shader, "environmentMap"), sky.convertIntUseForShader(rl.MapCubemap), rl.ShaderUniformInt)
	rl.SetShaderValue(SkyBox.Materials[0].Shader, rl.GetShaderLocation(SkyBox.Materials[0].Shader, "doGamma"), sky.convertIntUseForShader(Gamma), rl.ShaderUniformInt)
	rl.SetShaderValue(SkyBox.Materials[0].Shader, rl.GetShaderLocation(SkyBox.Materials[0].Shader, "vflipped"), sky.convertIntUseForShader(flipped), rl.ShaderUniformInt)

	img := rl.LoadImage(file)
	SkyBox.Materials[0].GetMap(rl.MapCubemap).Texture = rl.LoadTextureCubemap(img, rl.CubemapLayoutAutoDetect)

	defer rl.UnloadImage(img)
}

// in the Golang not work int we need convert
func (sky *SkyBox) convertIntUseForShader(value int) []float32 {
	return unsafe.Slice((*float32)(unsafe.Pointer(&value)), 1)
}

func (sky *SkyBox) Draw() {
	// We are inside the cube, we need to disable backface culling!
	rl.DisableBackfaceCulling()
	rl.DisableDepthMask()
	rl.DrawModel(*sky.cube, rl.Vector3{0, 0, 0}, 1.0, rl.White)
	rl.EnableBackfaceCulling()
	rl.EnableDepthMask()
}

func (sky *SkyBox) configShader() *rl.Shader {
	if checkPlatformIsMobail() {
		return ptr(rl.LoadShaderFromMemory(ShaderFile.Mobail["skyboxVs"], ShaderFile.Mobail["skyboxFs"]))
	} else {
		return ptr(rl.LoadShaderFromMemory(ShaderFile.Desktop["skyboxVs"], ShaderFile.Desktop["skyboxFs"]))
	}
}

func (sky *SkyBox) Unload() {
	rl.UnloadModel(*sky.cube)
}
