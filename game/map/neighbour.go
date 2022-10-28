package _map

import (
	"fmt"
	"github.com/alien_invasion/game/types"
	"math/rand"
	"strings"
)

type direction uint8

const (
	north direction = iota
	west
	south
	east
)

func newDirection(str string) direction {
	switch str {
	case "north":
		return north
	case "west":
		return west
	case "south":
		return south
	case "east":
		return east
	default:
		panic("unknown direction")
	}
}

func (d direction) String() string {
	switch d {
	case north:
		return "north"
	case west:
		return "west"
	case east:
		return "east"
	case south:
		return "south"
	default:
		return ""
	}
}

type neighbourhood map[direction]types.City

func (n *neighbourhood) String() string {
	var result string
	for dir, city := range *n {
		result += fmt.Sprintf("%s=%s ", dir, city)
	}

	return result
}

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

func generateNeighbourhood(neighbours ...string) *neighbourhood {
	neighbourhood := make(neighbourhood)

	for _, neighbour := range neighbours {
		info := strings.Split(neighbour, "=")
		if len(info) != 2 {
			panic("bad format")
		}

		var (
			dir  = info[0]
			city = info[1]
		)

		neighbourhood.add(newDirection(dir), types.City(city))
	}

	return &neighbourhood
}
