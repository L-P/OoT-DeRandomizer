package main

// A Location is any place that can hold an Item
type Location struct {
	Name     string `json:"name"`
	Address  uint32 `json:"address"`
	Address2 uint32 `json:"address2"`
	Default  uint16 `json:"default"`
	Scene    byte   `json:"scene"`
	Type     string `json:"type"`
}
