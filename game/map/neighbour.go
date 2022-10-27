package _map

import "github.com/alien_invasion/game/types"

type direction uint8

const (
	north direction = iota
	west
	south
	east
)

type neighbourhood map[direction]types.City
