package _map

import (
	"github.com/alien_invasion/game/types"
	"math/rand"
)

type direction uint8

const (
	north direction = iota
	west
	south
	east
)

type neighbourhood map[direction]types.City

func (n *neighbourhood) add(d direction, city types.City) {
	(*n)[d] = city
}

func (n *neighbourhood) remove(city types.City) {
	for d, c := range *n {
		if c.Name() == city.Name() {
			delete(*n, d)
		}
	}
}

func (n *neighbourhood) isNeighbour(city types.City) bool {
	for _, c := range *n {
		if c.Name() == city.Name() {
			return true
		}
	}

	return false
}

func (n *neighbourhood) getRandomNeighbour() types.City {
	neighbours := make([]types.City, 0, len(*n))
	for _, city := range *n {
		neighbours = append(neighbours, city)
	}

	return neighbours[rand.Intn(len(*n))]
}
