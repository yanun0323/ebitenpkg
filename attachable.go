package ebitenpkg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/pkg/logs"
)

type attachable interface {
	coords

	ID() ID
	Draw(screen *ebiten.Image)
}

func attach(parent, child attachable) {
	if parent == nil || child == nil {
		return
	}

	detach(child)

	switch p := parent.(type) {
	case *eImage:
		previous, ok := p.children.Swap(child.ID(), child)
		if ok && previous != nil {
			detach(previous)
		}
	case *eText:
		previous, ok := p.children.Swap(child.ID(), child)
		if ok && previous != nil {
			detach(previous)
		}
	default:
		logs.Fatalf("invalid attachable type: %T", p)
	}

	switch c := child.(type) {
	case *eImage:
		c.parent.Store(parent)
	case *eText:
		c.parent.Store(parent)
	default:
		logs.Fatalf("invalid attachable type: %T", c)
	}
}

func detach(child attachable) {
	var cParent attachable
	switch c := child.(type) {
	case *eImage:
		cParent, _ = c.parent.Delete()
	case *eText:
		cParent, _ = c.parent.Delete()
	default:
		logs.Fatalf("invalid attachable type: %T", c)
	}

	if cParent != nil {
		switch p := cParent.(type) {
		case *eImage:
			p.children.Delete(child.ID())
		case *eText:
			p.children.Delete(child.ID())
		default:
			logs.Fatalf("invalid attachable type: %T", cParent)
		}
	}
}
