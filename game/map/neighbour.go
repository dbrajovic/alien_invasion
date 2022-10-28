package _map

import "github.com/alien_invasion/game/types"

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
