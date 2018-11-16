package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	romPath := os.Args[1]
	log.Printf("Working on ROM %s", romPath)

	rom, err := NewROM(romPath)
	if err != nil {
		log.Fatal(err)
	}

	items, err := loadItems("data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	locations, err := loadLocations("data/locations.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", items)
	log.Printf("%+v\n", locations)
	log.Printf("%+v\n", rom)
}

func loadItems(path string) (map[uint32]Item, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	items := make(map[string]Item, 200)
	ret := make(map[uint32]Item)
	if err := json.NewDecoder(r).Decode(&items); err != nil {
		return nil, err
	}

	for k := range items {
		ret[items[k].Index] = items[k]
	}

	return ret, err
}

func loadLocations(path string) (map[string]Location, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	locations := make(map[string]Location, 680)
	if err := json.NewDecoder(r).Decode(&locations); err != nil {
		return nil, err
	}
	return locations, nil
}
