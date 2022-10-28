package _map

import (
	"github.com/alien_invasion/game/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	New("game_map")
}

func TestMap_RemoveCity(t *testing.T) {
	t.Run(
		"middle city is removed",
		func(t *testing.T) {
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
			assert.NotContains(t, m.cities, "dubrovnik")

			assert.False(t, m.cities["belgrade"].isNeighbour("dubrovnik"))
			assert.False(t, m.cities["berlin"].isNeighbour("dubrovnik"))
		},
	)

	t.Run(
		"left city is removed",
		func(t *testing.T) {
			m := &Map{cities: map[types.City]*neighbourhood{
				"belgrade": {
					west: "dubrovnik",
				},

				"dubrovnik": {
					east: "belgrade",
				},
			}}

			m.RemoveCity("dubrovnik")

			assert.Len(t, m.cities, 1)
			assert.NotContains(t, m.cities, "dubrovnik")

			assert.False(t, m.cities["belgrade"].isNeighbour("dubrovnik"))
		},
	)

}
