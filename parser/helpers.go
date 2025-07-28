package parser

import (
	"fmt"
	"math/rand"
	"time"
)

// ---------------------------------------------------------------------
// Helper: checks that enough bytes are available to read a field
func assertCanRead(data []byte, offset int, need int, context string) error {
	if offset+need > len(data) {
		return fmt.Errorf("buffer too short for %s at offset %d", context, offset)
	}
	return nil
}

// ---------------------------------------------------------------------

// ---------------------------------------------------------------------
// Helpers to convert byte slices to integer types

func bytesToUint8(b byte) uint8 {
	// Use when protocol says "unsigned byte" (0..255)
	// This func is not need it. Since in Go a single byte already fits the uint8 type.
	return uint8(b)
}

func bytesToInt8(b byte) int8 {
	// Use when protocol says "signed byte" (-128..127)
	return int8(b)
}

func bytesToUint16(b []byte) uint16 {
	// Use for 2-byte unsigned (big endian)
	return uint16(b[0])<<8 | uint16(b[1])
}

func bytesToInt16(b []byte) int16 {
	// Use for 2-byte signed (big endian)
	return int16(b[0])<<8 | int16(b[1])
}

func bytesToUint32(b []byte) uint32 {
	// Use for 4-byte unsigned (big endian)
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func bytesToInt32(b []byte) int32 {
	// Use for 4-byte signed (big endian)
	return int32(bytesToUint32(b))
}

func bytesToUint64(b []byte) uint64 {
	// Use for 8-byte unsigned (big endian)
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 |
		uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

// placeholderDebugFunc is a temporary function for simulating random parse errors
// during development or testing. This should be removed or replaced in production code.
func placeholderDebugFunc() error {

	if false {
		// Generate current time (RFC3339 for clarity)
		e1 := fmt.Sprintf("[debugTesting] Now is:%s", time.Now().Format(time.RFC3339Nano))

		// Generate a random int between 0 and 999
		n := rand.Intn(1000)
		e2 := fmt.Sprintf("[debugTesting] Random number:%d", n)

		return fmt.Errorf("%s %s", e1, e2)
	}

	// -----------------------------------------------------------
	// -----------------------------------------------------------

	return nil
}
