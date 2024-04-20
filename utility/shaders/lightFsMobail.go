package shaders

const LightFsMobail = `#version 100

precision mediump float;

// Input vertex attributes (from vertex shader)
varying vec3 fragPosition;
varying vec2 fragTexCoord;
//in vec4 fragColor;
varying vec3 fragNormal;

// Input uniform values
uniform sampler2D texture0;
uniform sampler2D texture1;
uniform sampler2D flashlight;

uniform vec4 colDiffuse;



// NOTE: Add here your custom variables

#define MAX_LIGHTS 5
#define LIGHT_DIRECTIONAL 0
#define LIGHT_POINT 1
#define LIGHT_SPOT 2

struct Light{
    int enabled;
    int type;
    float enargy;
    float cutOff ;
    float outerCutOff;
    float constant;
    float linear ;
    float quadratic ;
    float shiny ;
    float specularStr;
    vec3 position;
    vec3 direction;
    vec3 lightColor;
};

// Input lighting values
uniform Light lights[MAX_LIGHTS];
uniform vec3  ambientColor ;
uniform float ambientStrength ;
// uniform float shiny = 32.0;
// uniform float specularStr = 0.9;
// uniform float constant = 1.0;
// uniform float linear = 0.09;
// uniform float quadratic = 0.032;
// uniform float cutOff = 12.5;
// uniform float outerCutOff = 17.5;
uniform vec3  viewPos;

vec3 CalcDirLight(Light light,vec3 normal,vec3 viewDir,vec3 ambient);
vec3 CalcPointLight(Light light,vec3 normal,vec3 viewDir,vec3 ambient);
vec3 CalcSpotLight(Light light,vec3 normal,vec3 viewDir,vec3 ambient);


void main()
{
    
    vec3 norm=normalize(fragNormal);
    vec3 viewDir=normalize(viewPos-fragPosition);
    vec3 ambient = ambientStrength*ambientColor*vec3(texture2D(texture0,fragTexCoord));
    vec3 result = vec3(0);

    for (int i = 0; i < MAX_LIGHTS; i++){
       
        if(lights[i].enabled == 1){

            if(lights[i].type == LIGHT_DIRECTIONAL){
                result += CalcDirLight(lights[i],norm,viewDir,ambient);  
            }

            if(lights[i].type == LIGHT_POINT){
                result += CalcPointLight(lights[i],norm,viewDir,ambient);
            }

            if(lights[i].type == LIGHT_SPOT){
                 result += CalcSpotLight(lights[i],norm,viewDir,ambient);   
            }

        }

    } 
    
     gl_FragColor = vec4(result,1);
}

vec3 CalcDirLight(Light light,vec3 normal,vec3 viewDir,vec3 ambient)
{
    
    // diff
    vec3 lightDir=normalize(-light.direction);
    float diff=max(dot(normal,lightDir),0.0);
    vec3 diffuse=light.lightColor*diff*vec3(texture2D(texture0,fragTexCoord));
    
    //specular
    float specularStrength = light.specularStr;
    vec3 reflectDir=reflect(-lightDir,normal);
    float spec=pow(max(dot(viewDir,reflectDir),0.0),light.shiny);
    vec3 specular=specularStrength*spec*light.lightColor;
    
    return vec3(ambient+diffuse+specular);
}

vec3 CalcPointLight(Light light,vec3 normal,vec3 viewDir,vec3 ambient)
{
    
    // diff
    vec3 lightDir=normalize(light.position-fragPosition);
    float diff=max(dot(normal,lightDir),0.);
    vec3 diffuse=light.lightColor*diff*vec3(texture2D(texture0,fragTexCoord));
    
    //specular
    float specularStrength= light.specularStr;
    vec3 reflectDir=reflect(-lightDir,normal);
    float spec=pow(max(dot(viewDir,reflectDir),0.0),light.shiny);
    vec3 specular=specularStrength*spec*light.lightColor;
    
    float distance=length(light.position-fragPosition);
    float attenuation=light.enargy/(light.constant+light.linear*distance+light.quadratic*(distance*distance));
    ambient*=attenuation;
    diffuse*=attenuation;
    specular*=attenuation;
        
    return vec3(ambient+diffuse+specular);
}
    
    vec3 CalcSpotLight(Light light,vec3 normal,vec3 viewDir,vec3 ambient)
    {
        
        // diff
        vec3 lightDir=normalize(light.position-fragPosition);
        float diff=max(dot(normal,lightDir),0.0);
        vec3 diffuse=light.lightColor*diff*vec3(texture2D(texture0,fragTexCoord));
        
        //specular
        float specularStrength=light.specularStr;
        vec3 reflectDir=reflect(-lightDir,normal);
        float spec=pow(max(dot(viewDir,reflectDir),0.0),light.shiny);
        vec3 specular=specularStrength*spec*light.lightColor;
        
        float distance=length(light.position-fragPosition);
        float attenuation=light.enargy/(light.constant+light.linear*distance+light.quadratic*(distance*distance));
        ambient*=attenuation;
        diffuse*=attenuation;
        specular*=attenuation;
            
        float theta=dot(lightDir,normalize(-light.direction));
        float epsilon=cos(radians(light.cutOff))-cos(radians(light.outerCutOff));
        float intensity=smoothstep(0.0,1.0,(theta-cos(radians(light.outerCutOff)))/epsilon);//clamp((theta-cos(radians(light.outerCutOff)))/epsilon,0.0,1.0);
        intensity*=length(vec3(texture2D(flashlight,(fragTexCoord))));
        diffuse*=intensity;
        specular*=intensity;
            
        return vec3(ambient+diffuse+specular);
}`