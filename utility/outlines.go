package utility

import rl "github.com/mohsengreen1388/raylib-go-custom/raylib"

type OutLine struct {
	Shader          *rl.Shader
	texture         *rl.Texture2D
	textureSize     rl.Vector2
	outlineSizeLoc  int32
	outlineColorLoc int32
	textureSizeLoc  int32
}

// use BeginShaderMode for draw texture
func (out *OutLine) Init(texture *rl.Texture2D, outlineSize float32, outlineColor rl.Color) {
	out.configShader()

	out.texture = texture
	out.textureSize = rl.Vector2{float32(texture.Width), float32(texture.Height)}
	out.outlineSizeLoc = rl.GetShaderLocation(*out.Shader, "outlineSize")
	out.outlineColorLoc = rl.GetShaderLocation(*out.Shader, "outlineColor")
	out.textureSizeLoc = rl.GetShaderLocation(*out.Shader, "textureSize")

	// Set shader values (they can be changed later)
	rl.SetShaderValue(*out.Shader, out.outlineSizeLoc, []float32{outlineSize}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*out.Shader, out.outlineColorLoc, []float32{float32(outlineColor.R) / 255, float32(outlineColor.G) / 255, float32(outlineColor.B) / 255, float32(outlineColor.A) / 255}, rl.ShaderUniformVec4)
	rl.SetShaderValue(*out.Shader, out.textureSizeLoc, []float32{out.textureSize.X, out.textureSize.Y}, rl.ShaderUniformVec2)
}

func (out *OutLine) OutLineSize(size float32) {
	rl.SetShaderValue(*out.Shader, out.outlineSizeLoc, []float32{size}, rl.ShaderUniformFloat)
}

func (out *OutLine) OutLineColor(outlineColor rl.Color) {
	rl.SetShaderValue(*out.Shader, out.outlineColorLoc, []float32{float32(outlineColor.R) / 255, float32(outlineColor.G) / 255, float32(outlineColor.B) / 255, float32(outlineColor.A) / 255}, rl.ShaderUniformVec4)
}

func (out *OutLine) configShader() {
	if checkPlatformIsMobail() {
		out.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Mobail[""], ShaderFile.Mobail["outlineFs"]))
	} else {
		out.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Desktop[""], ShaderFile.Desktop["outlineFs"]))
	}
}

func (out *OutLine) Unload() {
	rl.UnloadTexture(*out.texture)
	rl.UnloadShader(*out.Shader)
}
