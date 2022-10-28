package engine

import (
	"testing"

	"github.com/alien_invasion/game/types"
	"github.com/stretchr/testify/assert"
)

func TestEngine_isDone(t *testing.T) {
	t.Parallel()

	t.Run(
		"aliens are maxed out on travels",
		func(t *testing.T) {
			t.Parallel()

			e := &Engine{
				aliens: []*types.Alien{
					{
						Name:     "alien_1",
						Location: "belgrade",
						Travels:  maxTravels,
					},
				},
				gameMap: nil,
			}

			assert.True(t, e.isDone())
		},
	)

	t.Run(
		"aliens are dead",
		func(t *testing.T) {
			t.Parallel()

			e := &Engine{
				aliens:  []*types.Alien{},
				gameMap: nil,
			}

			assert.True(t, e.isDone())
		},
	)

	t.Run(
		"aliens are not maxed out on travels",
		func(t *testing.T) {
			t.Parallel()

			e := &Engine{
				aliens: []*types.Alien{
					{
						Name:     "alien",
						Location: "belgrade",
						Travels:  0,
					},
				},
				gameMap: nil,
			}

			assert.False(t, e.isDone())
		},
	)
}

func TestEngine_moveAlien(t *testing.T) {
	t.Parallel()

	e := &Engine{
		aliens: []*types.Alien{
			{
				Name:     "alien",
				Location: "belgrade",
				Travels:  0,
			},
		},
		gameMap: mockMap{
			neighbourCallback: func(_ types.City) types.City { return "barcelona" },
		},
	}

	e.moveAliens()

	assert.Equal(t,
		"barcelona",
		e.aliens[0].Location.Name(),
	)

	assert.Equal(t,
		uint64(1),
		e.aliens[0].Travels,
	)
}

func TestEngine_aliensFight(t *testing.T) {
	t.Parallel()

	t.Run(
		"no fight",
		func(t *testing.T) {
			t.Parallel()

			e := &Engine{
				aliens: []*types.Alien{
					{
						Name:     "alien_1",
						Location: "belgrade",
						Travels:  0,
					},
					{
						Name:     "alien_2",
						Location: "barcelona",
						Travels:  0,
					},
				},
				gameMap: nil,
			}

			destroyedCities := e.aliensFight()

			assert.Len(t, e.aliens, 2)
			assert.Nil(t, destroyedCities)
		},
	)

	t.Run(
		"2 aliens die and a city is destroyed",
		func(t *testing.T) {
			t.Parallel()
			
			e := &Engine{
				aliens: []*types.Alien{
					{
						Name:     "alien_1",
						Location: "belgrade",
						Travels:  0,
					},
					{
						Name:     "alien_2",
						Location: "belgrade",
						Travels:  0,
					},
				},
				gameMap: nil,
			}

			destroyedCities := e.aliensFight()

			assert.Len(t, e.aliens, 0)
			assert.Len(t, destroyedCities, 1)
			assert.Equal(t, destroyedCities[0].Name(), "belgrade")
		},
	)
}

type mockMap struct {
	citiesCallback    func() []types.City
	neighbourCallback func(types.City) types.City
	removeCallback    func(city types.City)
}

func (m mockMap) Cities() []types.City {
	if m.citiesCallback == nil {
		return nil
	}

	return m.citiesCallback()
}

func (m mockMap) RandomNeighbourCity(city types.City) types.City {
	if m.neighbourCallback == nil {
		return ""
	}

	return m.neighbourCallback(city)
}

func (m mockMap) RemoveCity(city types.City) {
	if m.removeCallback == nil {
		return
	}

	m.removeCallback(city)
}

func (m mockMap) Display() {

}
