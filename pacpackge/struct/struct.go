package main

type GruruTags struct {
	Tags     map[string]TagConfig `json:"tags"`
	gurusMap map[string][]string
}
