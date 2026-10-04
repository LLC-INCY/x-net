package http2

import "testing"

// Exercise both long-lived XHTTP uploads and finite request bodies after rebase.
func TestIncyFrameScratchBufferCapped(t *testing.T) {
	for _, tc := range []struct {
		length      int64
		frame, want int
	}{
		{-1, 1 << 20, 32 << 10}, {-1, 16 << 10, 16 << 10},
		{0, 1 << 20, 1}, {100, 1 << 20, 101}, {1 << 20, 1 << 20, 32 << 10},
	} {
		cs := &clientStream{reqBodyContentLength: tc.length}
		if got := cs.frameScratchBufferLen(tc.frame); got != tc.want {
			t.Fatalf("length=%d frame=%d: got %d, want %d", tc.length, tc.frame, got, tc.want)
		}
	}
}
