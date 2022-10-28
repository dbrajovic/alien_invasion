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
		cities: loadMap(filename),
	}
}

func (m *Map) Display() {
	for city, nbhd := range m.cities {
		log.Println(city, nbhd)
	}
}

func (m *Map) Cities() []types.City {
	cities := make([]types.City, 0, len(m.cities))
	for city := range m.cities {
		cities = append(cities, city)
	}

	return cities
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

func loadMap(filename string) map[types.City]*neighbourhood {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("cannot open file:", filename, "err=", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("cannot close file:", "err=", err)
		}
	}()

	return loadCities(f)
}

func loadCities(file *os.File) map[types.City]*neighbourhood {
	cities := make(map[types.City]*neighbourhood)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		words := strings.Fields(line)

		city := types.City(words[0])
		cities[city] = generateNeighbourhood(words[1:]...)
	}

	if err := sc.Err(); err != nil {
		log.Fatal("scanner error:", err)
	}

	return cities
}
