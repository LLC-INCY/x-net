package http2

import "testing"

// TestIncyFrameScratchBufferCapped guards the incy 32KB patch in
// frameScratchBufferLen. If an upstream rebase drops or moves the patch,
// this fails loudly instead of silently restoring the 512KB buffer.
func TestIncyFrameScratchBufferCapped(t *testing.T) {
	const want = 32 << 10
	cs := &clientStream{reqBodyContentLength: -1} // -1 = unknown length (gRPC/duplex)
	got := cs.frameScratchBufferLen(1 << 20)      // peer advertises a huge 1MB frame size
	if got > want {
		t.Fatalf("frameScratchBufferLen returned %d, want <= %d (incy 32KB patch missing?)", got, want)
	}
}
