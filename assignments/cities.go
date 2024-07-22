package main

import (
	"math/rand"
	"strings"
)

var (
	prefixes = []string{"New", "Old", "Fort", "Port", "Lake", "Mount", "West", "East", "North", "South"}
	roots    = []string{"Spring", "Wood", "Lake", "River", "Valley", "Hill", "Field", "Land", "Burg", "Town", "City", "Village", "Haven", "Ford", "Bridge", "Creek"}
	suffixes = []string{"ton", "ville", "burg", "berg", "shire", "port", "field", "haven", "dale", "moor", "chester", "mouth", "ford", "wick", "gate", "wood"}
)

func generateCityName() string {

	// Decide on the structure of the name
	structure := rand.Intn(4)

	var parts []string

	switch structure {
	case 0:
		// Prefix + Root
		parts = append(parts, prefixes[rand.Intn(len(prefixes))])
		parts = append(parts, roots[rand.Intn(len(roots))])
	case 1:
		// Root + Suffix
		parts = append(parts, roots[rand.Intn(len(roots))])
		parts = append(parts, suffixes[rand.Intn(len(suffixes))])
	case 2:
		// Prefix + Root + Suffix
		parts = append(parts, prefixes[rand.Intn(len(prefixes))])
		parts = append(parts, roots[rand.Intn(len(roots))])
		parts = append(parts, suffixes[rand.Intn(len(suffixes))])
	case 3:
		// Just Root
		parts = append(parts, roots[rand.Intn(len(roots))])
	}

	return strings.Join(parts, "")
}
