package shaders

const Light2dFsDesktop = `#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec2 lightPos=vec2(10,10);
uniform float lightRadius=120.5;
uniform vec3 ambientColor=vec3(0.f,0.f,0.f);
uniform float ambientStrong = 1;
// Output fragment color
out vec4 finalColor;

void main()
{
    float distance= distance(gl_FragCoord.xy,vec2(lightPos.xy));
    vec4 color=fragColor;
    
    if(distance<lightRadius){
       discard;
    }else{
         color.rgb = ambientColor;
         color.a = ambientStrong;
    }



    finalColor=vec4(color);
}
`