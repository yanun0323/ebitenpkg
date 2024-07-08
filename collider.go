package ebitenpkg

import (
	"image/color"

	"github.com/google/uuid"
)

var _collidedColor = color.RGBA{R: 100, A: 100}

type ID uuid.UUID

type CollisionType int

type collider struct {
	id ID
	bt CollisionType
}

func newCollider(bt CollisionType) collider {
	return collider{
		id: ID(uuid.New()),
		bt: bt,
	}
}

func (c collider) ID() ID {
	return c.id
}

func (c collider) Type() CollisionType {
	return c.bt
}
