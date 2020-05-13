package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ruiqimao/go-gb-emu/gb"
	"github.com/ruiqimao/go-gfx/gfx"
)

const (
	DisplayScale = 2
)

// Handles all interaction with the user.
type Display struct {
	window *glfw.Window

	// Graphics objects.
	program *gfx.Program   // Shader program.
	quad    *gfx.Vao       // Quad shape.
	texture *gfx.Texture2D // Display texture.
}

// Create a new Display.
func NewDisplay() (*Display, error) {
	d := &Display{}

	// Initialize the window.
	var err error
	d.window, err = gfx.NewWindow(DisplayScale*gb.FrameWidth, DisplayScale*gb.FrameHeight, "Game Boy", false)
	if err != nil {
		return nil, err
	}

	// Attach a key callback.
	d.window.SetKeyCallback(d.keyCallback)

	// Init and start the display loop.
	err = d.init()
	if err != nil {
		return nil, err
	}
	go d.run()

	return d, nil
}

// Draw a frame.
func (d *Display) Draw(frame []uint8) {
	gfx.Do(func() {
		// Update the texture.
		if d.texture != nil {
			d.texture.SetData(gl.Ptr(frame), gl.RED, gl.UNSIGNED_BYTE)
		}
	})
}

// Graphics initialization.
func (d *Display) init() error {
	// Create the shader program.
	var err error
	d.program, err = gfx.NewProgram(vertexShader, "", fragmentShader)
	if err != nil {
		return err
	}

	// Make the VBO and VAO.
	buf := []float32{ // Simple quad.
		-1.0, -1.0,
		-1.0, 1.0,
		1.0, -1.0,
		1.0, 1.0,
	}
	markers := []gfx.AttribMarker{
		gfx.NewAttribMarker(0, gl.FLOAT, false, 2, 0),
	}
	vbo := gfx.NewVbo(buf, markers)
	d.quad = gfx.NewVao(vbo, gl.TRIANGLE_STRIP)

	// Make the texture.
	d.texture = gfx.NewTexture2D(nil, gb.FrameWidth, gb.FrameHeight, gl.RED, gl.UNSIGNED_BYTE)
	d.texture.Bind()
	d.texture.SetParam(gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	d.texture.SetParam(gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	d.texture.SetParam(gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	d.texture.SetParam(gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	d.texture.Unbind()

	return nil
}

// Display loop.
func (d *Display) run() {
	for !d.window.ShouldClose() {
		gfx.Do(func() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			d.program.Bind()

			// Draw the display.
			d.program.SetTexture2D("tex", d.texture)
			d.quad.Bind()
			d.quad.Draw()
			d.quad.Unbind()

			// Refresh the window.
			d.window.SwapBuffers()
			glfw.PollEvents()
		})
	}
	gfx.Halt()
}

// Key callback.
func (d *Display) keyCallback(window *glfw.Window, key glfw.Key, scrollCount int, action glfw.Action, mod glfw.ModifierKey) {
	// TODO.
}
