package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	settings := Settings{}
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

	world, err := NewWorld(settings, rom)
	if err != nil {
		log.Fatal(err)
	}

	combined, err := world.MapLocationToItem()
	if err != nil {
		log.Fatal(err)
	}

	for location, item := range combined {
		fmt.Printf("%s: %s\n", location.Name, item.Name)
	}
}
