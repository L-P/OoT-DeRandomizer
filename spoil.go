package main

import (
	"errors"
	"fmt"
	"strings"
)

func outputLocationToItem(
	settings Settings,
	rom *ROM,
	items map[byte]Item,
	locations map[string]Location,
) error {
	for _, v := range rom.Overrides {
		l, i, err := findLocationItem(settings, v, items, locations)
		if err != nil {
			continue
		}

		fmt.Printf("%s: %s\n", l.Name, i.Name)
	}

	return nil
}

func findLocationItem(
	settings Settings,
	o ItemOverride,
	items map[byte]Item,
	locations map[string]Location,
) (Location, Item, error) {
	for _, location := range locations {
		if strings.Index(location.Name, "MQ") > -1 && settings.Quest == QuestTypeNormal {
			continue
		}

		if strings.Index(location.Name, "GS") > -1 && settings.IgnoreGS {
			continue
		}

		def := byte(location.Default & 0x00FF)
		if location.Type == "Chest" {
			def &= 0x1F
		}

		if o.Scene == location.Scene && o.Default == def {
			if items[o.ItemIndex].Type == "Token" && settings.IgnoreGS {
				continue
			}

			return location, items[o.ItemIndex], nil
		}
	}

	return Location{}, Item{}, errors.New("not match")
}
