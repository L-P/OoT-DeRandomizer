package main

// An Item is anything that can be obtained from a Location
type Item struct {
	Index  uint32 `json:"index"`
	Name   string `json:"name"`
	Object uint32 `json:"object"`
	Type   string `json:"type"`
}
