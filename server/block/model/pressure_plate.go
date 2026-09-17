package model

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

// PressurePlate is a model for pressure plate blocks.
type PressurePlate struct {
	// Powered is whether the pressure plate is in its pressed down state.
	Powered bool
}

// BBox returns the collision/bounding box of the pressure plate.
func (p PressurePlate) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	height := 1.0 / 16.0
	if p.Powered {
		height = 1.0 / 32.0
	}
	return []cube.BBox{cube.Box(1.0/16.0, 0, 1.0/16.0, 15.0/16.0, height, 15.0/16.0)}
}

// FaceSolid always returns false.
func (PressurePlate) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
