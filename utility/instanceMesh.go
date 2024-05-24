package utility

import rl "github.com/mohsengreen1388/raylib-go-custom/raylib"

type Instances struct{
	MaxInstance int
	Matrixs []rl.Matrix
	Material rl.Material
}

//init Instance  and in the shader mat4 mvpi = mvp*instanceTransform; and finall gl_Position = mvpi*vec4(vertexPosition, 1.0);
func (in *Instances)InitInstance(shader rl.Shader){
	in.Matrixs = make([]rl.Matrix,in.MaxInstance)
	shader.UpdateLocation(rl.ShaderLocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))
	in.Material.Shader = shader
}

// use loop and MaxInstance for dynamic 
func (in *Instances)SetValueToMatrix(iteam int32,translate,angle,scale rl.Vector3){
		rotateToMatrix := rl.MatrixRotateXYZ(rl.Vector3{angle.X,angle.Y,angle.Z})
		scaleToMatrix  := rl.MatrixMultiply(rl.MatrixScale(scale.X,scale.Y,scale.Z),rotateToMatrix)
		in.Matrixs[iteam] = rl.MatrixMultiply(scaleToMatrix,rl.MatrixTranslate(translate.X,translate.Y,translate.Z))
}

// draw mesh
func (in *Instances)Draw(mesh rl.Mesh){
	rl.DrawMeshInstanced(mesh,in.Material,in.Matrixs,in.MaxInstance)
}