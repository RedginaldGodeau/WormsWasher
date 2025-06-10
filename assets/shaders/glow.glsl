vec4 effect(vec4 color, Image tex, vec2 texture_coords, vec2 screen_coords)
{
    // Couleur bleue de base
    vec4 baseColor = vec4(0.0, 0.0, 1.0, 1.0);
    
    // Effet de glow simple
    vec2 center = vec2(0.5, 0.5);
    float distance = length(texture_coords - center);
    float glowAmount = 1.0 - smoothstep(0.0, 0.7, distance);
    
    // Augmenter la luminosité autour du centre
    vec4 glowColor = baseColor + vec4(glowAmount * 0.5, glowAmount * 0.5, glowAmount * 0.5, 0.0);
    
    return glowColor;
}