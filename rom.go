package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// ItemOverride contains a list of scenes and their item
type ItemOverride struct {
	Scene     byte
	PlayerID  byte
	Default   byte // (default | 0x1F) for Item.Type == Chest
	ItemIndex byte
}

// ROM represents a decompressed TLoZ:OoT NTSC 1.0 ROM
type ROM struct {
	file *os.File

	Overrides []ItemOverride
}

// NewROM creates a new ROM from a file path
func NewROM(path string) (*ROM, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open ROM: %s", err)
	}

	rom := &ROM{
		file:      file,
		Overrides: make([]ItemOverride, 0, 350),
	}

	if err := rom.loadOverrides(); err != nil {
		rom.file.Close()
		return nil, err
	}

	return rom, nil
}

const overridesOffset = 0x03481000

func (r *ROM) loadOverrides() error {
	r.file.Seek(overridesOffset, io.SeekStart)

	for {
		var v ItemOverride

		err := binary.Read(r.file, binary.BigEndian, &v)
		if err != nil {
			return err
		}
		if (v == ItemOverride{}) {
			break
		}

		r.Overrides = append(r.Overrides, v)
	}

	if len(r.Overrides) == 0 {
		return errors.New("no overrides found, probably not a randomized ROM")
	}

	log.Printf("Loaded %d overrides", len(r.Overrides))

	return nil
}

// Close io.Closer
func (r *ROM) Close() {
	r.file.Close()
}

// ReadUInt8 returns the byte at address
func (r *ROM) ReadUInt8(address uint32) (byte, error) {
	var v byte
	r.file.Seek(int64(address), io.SeekStart)
	err := binary.Read(r.file, binary.BigEndian, &v)
	return v, err
}

// ReadUInt16 returns the UInt16 at address
func (r *ROM) ReadUInt16(address uint32) (uint16, error) {
	var v uint16
	r.file.Seek(int64(address), io.SeekStart)
	err := binary.Read(r.file, binary.BigEndian, &v)
	return v, err
}
