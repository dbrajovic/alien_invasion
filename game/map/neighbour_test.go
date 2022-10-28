package _map

import (
	"github.com/alien_invasion/game/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeighbourhood_add(t *testing.T) {
	n := make(neighbourhood)

	n.add(south, "belgrade")

	assert.Len(t, n, 1)
	assert.Contains(t, n, types.City("belgrade"))
	assert.Equal(t, n["belgrade"], south)
}

func TestNeighbourhood_remove(t *testing.T) {
	n := make(neighbourhood)

	n.add(south, "belgrade")
	n.remove("belgrade")

	assert.Len(t, n, 0)
	assert.NotContains(t, n, types.City("belgrade"))
}
