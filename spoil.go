package main

import "fmt"

func outputLocationToItem(
	rom *ROM,
	items map[byte]Item,
	locations map[string]Location,
) error {
	for _, v := range rom.Overrides {
		l, i, err := findLocationItem(v, items, locations)
		if err != nil {
			return nil
		}

		fmt.Printf("%s: %s\n", l.Name, i.Name)
	}

	return nil
}

func findLocationItem(
	o ItemOverride,
	items map[byte]Item,
	locations map[string]Location,
) (Location, Item, error) {
	for _, location := range locations {
		def := byte(location.Default & 0x00FF)
		if location.Type == "Chest" {
			def &= 0x1F
		}

		if o.Scene == location.Scene && o.Default == def {
			return location, items[o.ItemIndex], nil
		}
	}

	return Location{}, Item{}, nil
}
