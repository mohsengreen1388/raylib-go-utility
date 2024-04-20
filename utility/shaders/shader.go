package shaders

type Shaders struct {
	Desktop map[string]string
	Mobail  map[string]string
}

var desktopLists = map[string]string{
	"pbrVs":     PbrVsDesktop,
	"pbrFs":     PbrFsDesktop,
	"fogFs":     FogFsDesktop,
	"lightVs":   LightVsDesktop,
	"lightFs":   LightFsDesktop,
	"combineVs": CombineVsDesktop,
	"combineFs": CombineFsDesktop,
	"waveFs":    WaveFsDesktop,
	"skyboxVs":  SkyboxVsDesktop,
	"skyboxFs":  SkyboxFsDesktop,
	"maskFs":    MaskFsDesktop,
	"light2dFs": Light2dFsDesktop,
	"outlineFs":   OutlineFsDesktop,
}

var mobailLists = map[string]string{
	"pbrVs":     PbrVsMobail,
	"pbrFs":     PbrFsMobail,
	"fogFs":     FogFsMobail,
	"lightVs":   LightVsMobail,
	"lightFs":   LightFsMobail,
	"combineVs": CombineVsMobail,
	"combineFs": CombineFsMobail,
	"waveFs":    WaveFsMobail,
	"skyboxVs":  SkyboxVsMobail,
	"skyboxFs":  SkyboxFsMobail,
	"maskFs":    MaskFsMobail,
	"light2dFs": Light2dFsMobail,
	"outlineFs":   OutlineFsMobail,
}

func NewShader() Shaders {
	return Shaders{
		Desktop: desktopLists,
		Mobail:  mobailLists,
	}
}
