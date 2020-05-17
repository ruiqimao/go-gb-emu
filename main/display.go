package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ruiqimao/go-gb-emu/gb"
	"github.com/ruiqimao/go-gfx/gfx"
)

const (
	FrameWidth   = 160
	FrameHeight  = 144
	DisplayScale = 2
)

// Handles all interaction with the user.
type Display struct {
	window *glfw.Window

	// Graphics objects.
	program *gfx.Program   // Shader program.
	quad    *gfx.Vao       // Quad shape.
	texture *gfx.Texture2D // Display texture.

	// Input event channel.
	I chan gb.Input
}

// Create a new Display.
func NewDisplay() (*Display, error) {
	d := &Display{
		I: make(chan gb.Input, 16),
	}

	// Initialize the window.
	var err error
	d.window, err = gfx.NewWindow(DisplayScale*FrameWidth, DisplayScale*FrameHeight, "Game Boy", false)
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
		// If the frame is nil, turn off the display.
		if frame == nil {
			// TODO.
			return
		}

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
	d.texture = gfx.NewTexture2D(nil, FrameWidth, FrameHeight, gl.RED, gl.UNSIGNED_BYTE)
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
// Keys are mapped:
//   W -> Up
//   A -> Left
//   S -> Down
//   D -> Right
//   K -> A
//   J -> B
//   N -> Start
//   B -> Select
func (d *Display) keyCallback(window *glfw.Window, key glfw.Key, scrollCount int, action glfw.Action, mod glfw.ModifierKey) {
	// Ignore repeat events.
	if action == glfw.Repeat {
		return
	}

	// Create an input event.
	var button int
	state := action == glfw.Press
	switch key {
	case glfw.KeyW:
		button = gb.JoypadUp
	case glfw.KeyA:
		button = gb.JoypadLeft
	case glfw.KeyS:
		button = gb.JoypadDown
	case glfw.KeyD:
		button = gb.JoypadRight
	case glfw.KeyK:
		button = gb.JoypadA
	case glfw.KeyJ:
		button = gb.JoypadB
	case glfw.KeyN:
		button = gb.JoypadStart
	case glfw.KeyB:
		button = gb.JoypadSelect
	default:
		return
	}
	event := gb.NewInput(button, state)

	// Try to push the event to the channel.
	select {
	case d.I <- event:
	default:
	}
}
