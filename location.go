package main

// A Location is any place that can hold an Item
type Location struct {
	Name     string `json:"name"`
	Address  uint32 `json:"address"`
	Address2 uint32 `json:"address2"`
	Default  uint32 `json:"default"`
	Hint     string `json:"hint"`
	Scene    uint32 `json:"scene"`
	Type     string `json:"type"`
}
