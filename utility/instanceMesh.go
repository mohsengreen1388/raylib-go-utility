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
	shader.UpdateLocation(rl.LocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))
	in.Material.Shader = shader
}

// use loop and MaxInstance for dynamic 
func (in *Instances)SetValueToMatrix(iteam int32,angle float32,translate,rotate,scale rl.Vector3){
		in.Matrixs[iteam] = rl.MatrixMultiply(rl.MatrixMultiply(rl.MatrixTranslate(translate.X,translate.Y,translate.Z), rl.MatrixRotate(rotate,angle)),rl.MatrixScale(scale.X,scale.Y,scale.Z))
}

// draw mesh
func (in *Instances)Draw(mesh rl.Mesh){
	rl.DrawMeshInstanced(mesh,in.Material,in.Matrixs,in.MaxInstance)
}