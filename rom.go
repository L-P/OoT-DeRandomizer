package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const crcAddress = 0x10
const crcLength = 0x08

// crc for the uncompressed ROM
var crc = []byte{0x93, 0x52, 0x2E, 0x7B, 0xE5, 0x06, 0xD4, 0x27}

// ROM represents a decompressed TLoZ:OoT NTSC 1.0 ROM
type ROM struct {
}

// NewROM creates a new ROM from a file path
func NewROM(path string) (*ROM, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open ROM: %s", err)
	}
	defer r.Close()

	if err := validateRom(r); err != nil {
		return nil, err
	}

	return &ROM{}, nil
}

func validateRom(r io.ReaderAt) error {
	romCRC := make([]byte, crcLength)
	n, err := r.ReadAt(romCRC, crcAddress)
	if n != crcLength {
		return fmt.Errorf("unable to read the %d bytes of CRC", crcLength)
	}
	if err != nil {
		return fmt.Errorf("unable to read file: %s", err)
	}

	if !bytes.Equal(romCRC, crc) {
		return errors.New("CRC does not match TLoZ:OoT NTSC 1.0 ROM")
	}

	log.Print("ROM CRC matches TLoZ:OoT NTSC 1.0 ROM")

	return nil
}
