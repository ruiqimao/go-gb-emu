package main

const vertexShader = `
#version 330 core

layout (location = 0) in vec2 position;

out vec2 pos;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	pos = position;
}
`

const fragmentShader = `
#version 330 core

out vec4 fragColor;

in vec2 pos;

uniform sampler2D tex;

void main() {
	// Get the pixel in the texture from the fragment position.
	float x = (pos.x + 1.0) * 0.5;
	float y = (1.0 - pos.y) * 0.5;
	float p = 1.0 - texture(tex, vec2(x, y)).r * 85.0;

	fragColor = vec4(p, p, p, 1.0);
}
`
