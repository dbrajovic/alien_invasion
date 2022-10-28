package _map

import (
	"bufio"
	"github.com/alien_invasion/game/types"
	"log"
	"os"
	"strings"
)

type Map struct {
	cities map[types.City]*neighbourhood
}

func New(filename string) *Map {
	return &Map{
		cities: loadCities(filename),
	}
}

func loadCities(filename string) map[types.City]*neighbourhood {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("cannot close file", "err=", err)
		}
	}()

	sc := bufio.NewScanner(f)

	cities := make(map[types.City]*neighbourhood)
	for sc.Scan() {
		line := sc.Text()

		words := strings.Fields(line)

		city := types.City(words[0])

		cities[city] = generateNeighbourhood(words[1:]...)
	}

	return cities
}

func (m *Map) Display() {

}

func (m *Map) Cities() []types.City {
	return nil
}

func (m *Map) RandomNeighbourCity(city types.City) types.City {
	return m.cities[city].getRandomNeighbour()
}

func (m *Map) RemoveCity(city types.City) {
	//	remove from map
	delete(m.cities, city)

	//	remove from all neighbrouhoods
	for _, neighbourhood := range m.cities {
		neighbourhood.remove(city)
	}
}
