package main

// An Item is anything that can be obtained from a Location
type Item struct {
	Index byte   `json:"index"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}
