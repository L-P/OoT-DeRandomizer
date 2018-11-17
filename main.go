package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

func main() {
	settings := Settings{}
	flag.StringVar(&settings.Quest, "quest", QuestTypeNormal, "normal|mixed|master")
	flag.BoolVar(&settings.IgnoreGS, "ignore-gs", false, "ignore gold skulltulas")
	flag.Parse()

	if len(flag.Args()) != 1 {
		os.Exit(1)
	}

	romPath := flag.Args()[0]
	log.Printf("Working on ROM %s", romPath)

	rom, err := NewROM(romPath)
	if err != nil {
		log.Fatal(err)
	}
	defer rom.Close()

	log.Printf("Loading item list")
	items, err := loadItems("data/items.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loading location list")
	locations, err := loadLocations("data/locations.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := outputLocationToItem(settings, rom, items, locations); err != nil {
		log.Fatal(err)
	}
}

func loadItems(path string) (map[byte]Item, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	items := make(map[string]Item, 200)
	ret := make(map[byte]Item)
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
