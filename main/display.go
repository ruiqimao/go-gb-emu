package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ruiqimao/go-gfx/gfx"
)

// Handles all interaction with the user.
type Display struct {
	window *glfw.Window // Display window.
}

// Create a new Display.
func NewDisplay() (*Display, error) {
	d := &Display{}

	// Initialize the window.
	var err error
	d.window, err = gfx.NewWindow(320, 288, "Game Boy", false)
	if err != nil {
		return nil, err
	}

	// Attach a key callback.
	d.window.SetKeyCallback(d.keyCallback)

	// Init and start the display loop.
	d.init()
	go d.run()

	return d, nil
}

// Graphics initialization.
func (d *Display) init() {
	// TODO.
}

// Display loop.
func (d *Display) run() {
	for !d.window.ShouldClose() {
		gfx.Do(func() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			// TODO.

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
