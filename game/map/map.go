package _map

import "github.com/alien_invasion/game/types"

type Map struct {
	cities map[types.City]*neighbourhood
}

func New(filename string) *Map {
	return &Map{}
}

func (m *Map) Display() {

}

func (m *Map) Cities() []types.City {
	return nil
}

func (m *Map) RandomNeighbourCity(city types.City) types.City {
	return ""
}

func (m *Map) RemoveCity(city types.City) {

}