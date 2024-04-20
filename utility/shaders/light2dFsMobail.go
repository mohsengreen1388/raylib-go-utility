package shaders

const Light2dFsMobail = `#version 100

precision mediump float;

// Input vertex attributes (from vertex shader)
varying vec2 fragTexCoord;
varying vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec2 lightPos;
uniform float lightRadius;
uniform vec3 ambientColor;
uniform float ambientStrong;
// Output fragment color


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



    gl_FragColor = vec4(color);
}
`