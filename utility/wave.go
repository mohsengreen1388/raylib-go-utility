package utility

import rl "github.com/mohsengreen1388/raylib-go-custom/raylib"

type Wave struct {
	Shader     *rl.Shader
	ScreenSize rl.Vector2
	Seconds    float32
	SecondsLoc int32
	FreqXLoc   int32
	FreqYLoc   int32
	AmpXLoc    int32
	AmpYLoc    int32
	SpeedXLoc  int32
	SpeedYLoc  int32
}

func (wa *Wave) Init(materals []*rl.Material, texture []*rl.Texture2D) {
	wa.configShader()
	
	for index, material := range materals {
		material.GetMap(rl.MapDiffuse).Texture = *texture[index]
	}

	wa.SecondsLoc = rl.GetShaderLocation(*wa.Shader, "secondes")
	wa.FreqXLoc = rl.GetShaderLocation(*wa.Shader, "freqX")
	wa.FreqYLoc = rl.GetShaderLocation(*wa.Shader, "freqY")
	wa.AmpXLoc = rl.GetShaderLocation(*wa.Shader, "ampX")
	wa.AmpYLoc = rl.GetShaderLocation(*wa.Shader, "ampY")
	wa.SpeedXLoc = rl.GetShaderLocation(*wa.Shader, "speedX")
	wa.SpeedYLoc = rl.GetShaderLocation(*wa.Shader, "speedY")
	wa.ScreenSize = rl.Vector2{float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())}

	wa.UpdateOptions(25.0, 25.0, 5.0, 5.0, 8.0, 8.0)

}

func (wa *Wave) InitFor2D() {
	wa.configShader()

	wa.SecondsLoc = rl.GetShaderLocation(*wa.Shader, "secondes")
	wa.FreqXLoc = rl.GetShaderLocation(*wa.Shader, "freqX")
	wa.FreqYLoc = rl.GetShaderLocation(*wa.Shader, "freqY")
	wa.AmpXLoc = rl.GetShaderLocation(*wa.Shader, "ampX")
	wa.AmpYLoc = rl.GetShaderLocation(*wa.Shader, "ampY")
	wa.SpeedXLoc = rl.GetShaderLocation(*wa.Shader, "speedX")
	wa.SpeedYLoc = rl.GetShaderLocation(*wa.Shader, "speedY")
	wa.ScreenSize = rl.Vector2{float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())}

	wa.UpdateOptions(25.0, 25.0, 5.0, 5.0, 8.0, 8.0)

}

// use in the loop
func (wa *Wave) UpdateWave() {
	wa.Seconds += rl.GetFrameTime()
	rl.SetShaderValue(*wa.Shader, wa.SecondsLoc, []float32{float32(wa.Seconds)}, rl.ShaderUniformFloat)
}

func (wa *Wave) UpdateOptions(freqX, freqY, ampX, ampY, speedX, speedY float32) {
	rl.SetShaderValue(*wa.Shader, rl.GetShaderLocation(*wa.Shader, "size"), []float32{wa.ScreenSize.X, wa.ScreenSize.Y}, rl.ShaderUniformVec2)
	rl.SetShaderValue(*wa.Shader, wa.FreqXLoc, []float32{float32(freqX)}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*wa.Shader, wa.FreqYLoc, []float32{float32(freqY)}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*wa.Shader, wa.AmpXLoc, []float32{float32(ampX)}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*wa.Shader, wa.AmpYLoc, []float32{float32(ampY)}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*wa.Shader, wa.SpeedXLoc, []float32{float32(speedX)}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*wa.Shader, wa.SpeedYLoc, []float32{float32(speedY)}, rl.ShaderUniformFloat)
}

func (wa *Wave) configShader() {
	if checkPlatformIsMobail() {
		wa.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Mobail[""], ShaderFile.Mobail["waveFs"]))
	} else {
		wa.Shader = ptr(rl.LoadShaderFromMemory(ShaderFile.Desktop[""], ShaderFile.Desktop["waveFs"]))
	}
}

func (wa *Wave) Unload() {
	rl.UnloadShader(*wa.Shader)
}
