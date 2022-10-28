package engine

import (
	"fmt"
	"math/rand"

	"github.com/alien_invasion/game/types"
)

/*	Generator helpers */

func spawnRandomAliens(num uint64, gameMap Map) []*types.Alien {
	aliens := make([]*types.Alien, num)
	for i := 0; i < len(aliens); i++ {
		aliens[i] = spawnAlien(i, gameMap.Cities())
	}

	return aliens
}

func spawnAlien(i int, cities []types.City) *types.Alien {
	return &types.Alien{
		Name:     generateName(i),
		Location: generateRandomCity(cities),
		Travels:  0,
	}
}

func generateName(i int) string {
	return fmt.Sprintf("alien_%d", i)
}

func generateRandomCity(cities []types.City) types.City {
	return cities[rand.Intn(len(cities))]
}
