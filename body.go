package ebitenpkg

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

type BodyType int

type body struct {
	Controller
	id    ID
	t     BodyType
	space Space
}

func newCollidable(s Space, t BodyType, ctr Controller) Collidable {
	return &body{
		Controller: ctr,
		id:         ID(uuid.New()),
		t:          t,
		space:      s,
	}
}

func (b *body) ID() ID {
	return b.id
}

func (b body) Type() BodyType {
	return b.t
}

func (b body) IsCollided() bool {
	return b.space.IsCollided(b.ID())
}

func (b body) IsCollide(p vector) bool {
	return isInside(b.vertexes(), p)
}

func (b body) GetCollided() []Collidable {
	return b.space.GetCollided(b.ID())
}

func (b *body) controller() Controller {
	return b.Controller
}
