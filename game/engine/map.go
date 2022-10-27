package engine

import "github.com/alien_invasion/game/types"

type Map interface {
	Cities() []types.City
	RandomNeighbourCity(types.City) types.City
	RemoveCity(city types.City)
}
