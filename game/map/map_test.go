package _map

import (
	"github.com/alien_invasion/game/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap_RemoveCity(t *testing.T) {
	m := &Map{cities: map[types.City]*neighbourhood{
		"belgrade": {
			west: "dubrovnik",
		},

		"dubrovnik": {
			east:  "belgrade",
			north: "berlin",
		},

		"berlin": {
			south: "dubrovnik",
		},
	}}

	m.RemoveCity("dubrovnik")

	assert.Len(t, m.cities, 2)
}
