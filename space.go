package ebitenpkg

import "sync"

type Space interface {
	GameUpdate()
	AddBody(Collidable) Space
	RemoveBody(Collidable) Space
	IsCollided(Collidable) bool
	GetCollided(Collidable) []Collidable
}

type Collidable interface {
	coords

	ID() ID
	Group() int
	Bounds() (w, h int)
	Parent() attachable
}

type space struct {
	mu       *sync.RWMutex
	bodies   map[ID]Collidable
	collided map[ID][]Collidable
}

func NewSpace() Space {
	return &space{
		mu:     &sync.RWMutex{},
		bodies: map[ID]Collidable{},
	}
}

func (s *space) GameUpdate() {
	s.mu.Lock()
	defer s.mu.Unlock()

	collided := make(map[ID][]Collidable, len(s.bodies))
	bs := make([]Collidable, 0, len(s.bodies))

	for _, b := range s.bodies {
		bs = append(bs, b)
	}

	for i := range bs {
		for j := i; j < len(bs); j++ {
			if bs[i].Group() == bs[j].Group() {
				continue
			}

			iv := s.getVertexes(bs[i])
			jv := s.getVertexes(bs[j])
			if isOverlap(iv, jv) || gjk(iv, jv) {
				collided[bs[i].ID()] = append(collided[bs[i].ID()], bs[j])
				collided[bs[j].ID()] = append(collided[bs[j].ID()], bs[i])
			}
		}
	}

	s.collided = collided
}

func (s *space) getVertexes(c Collidable) []Vector {
	w, h := c.Bounds()
	return getVertexes(float64(w), float64(h), c, c.Parent())
}

func (s *space) AddBody(c Collidable) Space {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.bodies[c.ID()] = c
	return s
}

func (s *space) RemoveBody(c Collidable) Space {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.bodies, c.ID())
	return s
}

func (s *space) IsCollided(c Collidable) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.collided) != 0 && len(s.collided[c.ID()]) != 0
}

func (s *space) GetCollided(c Collidable) []Collidable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cs := s.collided[c.ID()]
	result := make([]Collidable, len(cs))

	copy(result, cs)

	return result
}
