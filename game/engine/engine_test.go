package engine

import (
	"testing"

	"github.com/alien_invasion/game/types"
	"github.com/stretchr/testify/assert"
)

func TestEngine_isDone(t *testing.T) {
	t.Run(
		"aliens are maxed out on travels",
		func(t *testing.T) {
			e := New(
				[]*types.Alien{
					{
						Name:     "alien_1",
						Location: "belgrade",
						Steps:    maxTravels},
				},
				nil,
			)

			assert.True(t, e.isDone())
		},
	)

	t.Run(
		"aliens are dead",
		func(t *testing.T) {
			e := New(
				[]*types.Alien{},
				nil,
			)

			assert.True(t, e.isDone())
		},
	)

	t.Run(
		"aliens are not maxed out on travels",
		func(t *testing.T) {
			e := New(
				[]*types.Alien{
					{Name: "alien", Location: "belgrade", Steps: 0},
				},
				nil,
			)

			assert.False(t, e.isDone())
		},
	)
}

func TestEngine_moveAlien(t *testing.T) {
	e := New(
		[]*types.Alien{
			{
				Name:     "alien",
				Location: "belgrade",
				Steps:    0,
			},
		},
		mockMap{
			neighbourCallback: func(_ types.City) types.City { return "barcelona" },
		},
	)

	e.moveAliens()

	assert.Equal(t,
		"barcelona",
		e.aliens[0].Location.Name(),
	)

	assert.Equal(t,
		uint64(1),
		e.aliens[0].Steps,
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
