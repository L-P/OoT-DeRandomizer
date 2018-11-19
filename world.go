package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// World holds the rules, available locations and available items.
type World struct {
	locations []Location
	items     map[byte]Item
	symbols   map[string]uint32

	isMQ bool

	settings Settings
	rom      *ROM
}

// NewWorld creates a new World with its data already parsed and ready.
func NewWorld(settings Settings, rom *ROM) (*World, error) {
	world := &World{
		locations: make([]Location, 0, 256),
		items:     make(map[byte]Item, 256),
		symbols:   make(map[string]uint32),
		rom:       rom,
		settings:  settings,
	}

	if err := world.loadLocations("data/locations.json"); err != nil {
		return nil, err
	}
	if err := world.loadItems("data/items.json"); err != nil {
		return nil, err
	}
	if err := world.loadSymbols("data/symbols.json"); err != nil {
		return nil, err
	}
	if err := world.setFlags(); err != nil {
		return nil, err
	}

	return world, nil
}

func (w *World) setFlags() error {
	mq, err := w.rom.ReadUInt32(w.symbols["cfg_dungeon_is_mq"])
	if err != nil {
		return err
	}
	if mq > 0 {
		w.isMQ = true
	}

	return nil
}

// MapLocationToItem tries to guess where items go.
func (w *World) MapLocationToItem() (map[Location]Item, error) {
	ret := make(map[Location]Item)
	for _, o := range w.rom.Overrides {
		for _, location := range w.locations {
			if strings.Index(location.Name, "MQ") > -1 && !w.isMQ {
				continue
			}

			if strings.Index(location.Name, "GS") > -1 && w.settings.IgnoreGS {
				continue
			}

			def := byte(location.Default & 0x00FF)
			if location.Type == "Chest" {
				def &= 0x1F
			}

			if o.Scene == location.Scene && o.Default == def {
				item, ok := w.items[o.ItemIndex]
				if !ok {
					log.Printf("item %02X was not found", o.ItemIndex)
					continue
				}

				if item.Type == "Token" && w.settings.IgnoreGS {
					continue
				}

				ret[location] = item
			}
		}
	}

	if len(ret) > 0 {
		return ret, nil
	}

	return ret, errors.New("no matches")
}

func (w *World) loadLocations(path string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	if err := json.NewDecoder(r).Decode(&w.locations); err != nil {
		return err
	}

	return nil
}

func (w *World) loadItems(path string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	items := make([]Item, 0, 256)
	if err := json.NewDecoder(r).Decode(&items); err != nil {
		return err
	}

	for _, v := range items {
		w.items[v.Index] = v
	}

	return nil
}

func (w *World) loadSymbols(path string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	raw := make(map[string]string)
	if err := json.NewDecoder(r).Decode(&raw); err != nil {
		return err
	}

	for k, v := range raw {
		decoded, err := strconv.ParseUint(v, 16, 32)
		if err != nil {
			return err
		}
		w.symbols[k] = uint32(decoded)
	}

	return nil
}
