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
	Attachable

	CollisionID() ID
	CollisionGroup() int
	Bounds() (w, h int)
	Parent() Attachable
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
			if bs[i].CollisionGroup() == bs[j].CollisionGroup() {
				continue
			}

			iv := s.getVertexes(bs[i])
			jv := s.getVertexes(bs[j])
			if isOverlap(iv, jv) || gjk(iv, jv) {
				collided[bs[i].CollisionID()] = append(collided[bs[i].CollisionID()], bs[j])
				collided[bs[j].CollisionID()] = append(collided[bs[j].CollisionID()], bs[i])
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

	s.bodies[c.CollisionID()] = c
	return s
}

func (s *space) RemoveBody(c Collidable) Space {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.bodies, c.CollisionID())
	return s
}

func (s *space) IsCollided(c Collidable) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.collided) != 0 && len(s.collided[c.CollisionID()]) != 0
}

func (s *space) GetCollided(c Collidable) []Collidable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cs := s.collided[c.CollisionID()]
	result := make([]Collidable, len(cs))

	copy(result, cs)

	return result
}
