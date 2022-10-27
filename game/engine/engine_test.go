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
				[]types.Alien{
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
				[]types.Alien{},
				nil,
			)

			assert.True(t, e.isDone())
		},
	)

	t.Run(
		"aliens are not maxed out on travels",
		func(t *testing.T) {
			e := New(
				[]types.Alien{
					{Name: "alien", Location: "belgrade", Steps: 0},
				},
				nil,
			)

			assert.False(t, e.isDone())
		},
	)
}
