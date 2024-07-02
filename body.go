package ebitenpkg

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

type BodyType int

type body struct {
	Controller
	attachedCtr Controller
	id          ID
	t           BodyType
	space       Space
}

func newCollidable(s Space, t BodyType, collideCtr Controller, attachedCtr ...Controller) Collidable {
	attached := NewController(0, 0)
	if len(attachedCtr) != 0 && attachedCtr[0] != nil {
		attached = attachedCtr[0]
	}

	return &body{
		Controller:  collideCtr,
		attachedCtr: attached,
		id:          ID(uuid.New()),
		t:           t,
		space:       s,
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

func (b body) IsCollide(p Vector) bool {
	return isInside(b.vertexes(), p)
}

func (b body) GetCollided() []Collidable {
	return b.space.GetCollided(b.ID())
}

func (b *body) controller() Controller {
	return b.Controller
}

func (b *body) CollideCenter() Vector {
	return b.rotationCenter()
}
