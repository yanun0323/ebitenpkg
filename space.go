package ebitenpkg

import "sync"

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

func (s *space) Update() {
	s.mu.Lock()
	defer s.mu.Unlock()

	collided := make(map[ID][]Collidable, len(s.bodies))
	bs := make([]Collidable, 0, len(s.bodies))

	for _, b := range s.bodies {
		bs = append(bs, b)
	}

	for i := range bs {
		for j := i; j < len(bs); j++ {
			if bs[i].Type() == bs[j].Type() {
				continue
			}

			ivs := bs[i].controller().vertexes()
			jvs := bs[j].controller().vertexes()
			collide := false

			for _, p := range ivs {
				if isInside(jvs, p) {
					collide = true
					collided[bs[i].ID()] = append(collided[bs[i].ID()], bs[j])
					collided[bs[j].ID()] = append(collided[bs[j].ID()], bs[i])
					break
				}
			}

			if collide {
				continue
			}

			for _, p := range jvs {
				if isInside(ivs, p) {
					collide = true
					collided[bs[i].ID()] = append(collided[bs[i].ID()], bs[j])
					collided[bs[j].ID()] = append(collided[bs[j].ID()], bs[i])
					break
				}
			}

		}
	}

	s.collided = collided
}

func (s *space) AddBody(b Collidable) Space {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.bodies[b.ID()] = b
	return s
}

func (s *space) RemoveBody(id ID) Space {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.bodies, id)
	return s
}

func (s *space) IsCollided(id ID) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.collided) != 0 && len(s.collided[id]) != 0
}

func (s *space) GetCollided(id ID) []Collidable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c := s.collided[id]
	result := make([]Collidable, len(c))

	copy(result, c)

	return result
}
