package utility

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
	sh "github.com/mohsengreen1388/raylib-go-utility/utility/shaders"
	"reflect"
	"runtime"
	"unsafe"
)

var ShaderFile sh.Shaders = sh.NewShader()

func checkPlatformIsMobail() bool {
	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		return true
	}
	return false
}

func Cobineshader() rl.Shader {
	if checkPlatformIsMobail() {
		return rl.LoadShaderFromMemory(ShaderFile.Mobail["combineVs"], ShaderFile.Mobail["combineFs"])
	} else {
		return rl.LoadShaderFromMemory(ShaderFile.Desktop["combineVs"], ShaderFile.Desktop["combineFs"])
	}
}

func ptr(sh rl.Shader)*rl.Shader{
	return &sh
}

func generiteIntForGlsl(value int32) []float32 {
	data := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&value)),
		Len:  4,
		Cap:  4,
	}
	return *(*[]float32)(unsafe.Pointer(data))
}
