package _map

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/alien_invasion/game/types"
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
	println("Remaining cities:")
	for city, nbhd := range m.cities {
		println(city, nbhd.String())
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
	nbhd := m.cities[city]
	if nbhd.empty() {
		return city
	}

	return nbhd.getRandomNeighbour()
}

func (m *Map) RemoveCity(city types.City) {
	delete(m.cities, city)

	for _, neighbourhood := range m.cities {
		neighbourhood.remove(city)
	}
}

func loadMap(filename string) map[types.City]*neighbourhood {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v", filename, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("cannot close %s: %v", filename, err)
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
		cities[city] = parseNeighbourhood(words[1:]...)
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return cities
}
