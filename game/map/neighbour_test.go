package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeighbourhood_add(t *testing.T) {
	n := make(neighbourhood)

	n.add(south, "belgrade")

	assert.Len(t, n, 1)
	assert.Contains(t, n, south)
	assert.Equal(t, "belgrade", n[south].Name())
}

func TestNeighbourhood_remove(t *testing.T) {
	n := make(neighbourhood)

	n.add(south, "belgrade")
	n.remove("belgrade")

	assert.Len(t, n, 0)
	assert.NotContains(t, n, south)
}
