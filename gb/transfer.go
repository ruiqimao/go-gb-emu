package gb

// Execute a step of pixel transfer.
func (p *PPU) stepTransfer() {
	// TODO.
}

// Push the current frame into the channel.
func (p *PPU) pushFrame() {
	// Make a copy of the frame.
	frame := make([]byte, len(p.frame))
	copy(frame, p.frame[:])

	// Try to push the frame. If the channel is full, drop the frame.
	select {
	case p.gb.F <- frame:
	default:
	}
}
