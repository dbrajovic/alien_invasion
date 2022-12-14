package _map

import (
	"github.com/alien_invasion/game/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewMap(t *testing.T) {
	t.Parallel()

	cwd, _ := os.Getwd()
	m := New(cwd + "/map_test")

	assert.Contains(t, m.cities, types.City("Foo"))
	assert.Contains(t, m.cities, types.City("Boo"))

	assert.True(t, m.cities["Foo"].isNeighbour("Baz"))
	assert.True(t, m.cities["Foo"].isNeighbour("Boo"))
	assert.True(t, m.cities["Foo"].isNeighbour("Honolulu"))

	assert.True(t, m.cities["Boo"].isNeighbour("Foo"))
	assert.True(t, m.cities["Boo"].isNeighbour("Car"))
}

func TestMap_RemoveCity(t *testing.T) {
	t.Parallel()

	t.Run(
		"middle city is removed",
		func(t *testing.T) {
			t.Parallel()

			m := &Map{cities: map[types.City]*neighbourhood{
				"belgrade": {
					"dubrovnik": west,
				},

				"dubrovnik": {
					"belgrade": east,
					"berlin":   north,
				},

				"berlin": {
					"dubrovnik": south,
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
			t.Parallel()

			m := &Map{cities: map[types.City]*neighbourhood{
				"belgrade": {
					"dubrovnik": west,
				},

				"dubrovnik": {
					"belgrade": east,
				},
			}}

			m.RemoveCity("dubrovnik")

			assert.Len(t, m.cities, 1)
			assert.NotContains(t, m.cities, "dubrovnik")

			assert.False(t, m.cities["belgrade"].isNeighbour("dubrovnik"))
		},
	)

}
