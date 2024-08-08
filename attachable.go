package ebitenpkg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/pkg/logs"
)

type Attachable interface {
	coords

	ID() ID
	Draw(screen *ebiten.Image)
}

func attach(parent, child Attachable) {
	if parent == nil || child == nil {
		return
	}

	detach(child)

	logs.Debug("attach", child.ID())
	switch p := parent.(type) {
	case *eImage:
		previous, ok := p.children.FindAndSwap(func(c Attachable) bool {
			return c.ID() == child.ID()
		}, child)
		if !ok {
			p.children.Append(child)
		} else if previous != nil {
			detach(previous)
		}
	case *eText:
		previous, ok := p.children.FindAndSwap(func(c Attachable) bool {
			return c.ID() == child.ID()
		}, child)
		if !ok {
			p.children.Append(child)
		} else if previous != nil {
			detach(previous)
		}
	default:
		logs.Fatalf("invalid attachable type: %T", p)
	}

	logs.Debug("attached", child.ID())
	switch c := child.(type) {
	case *eImage:
		c.parent.Store(parent)
	case *eText:
		c.parent.Store(parent)
	default:
		logs.Fatalf("invalid attachable type: %T", c)
	}
}

func detach(child Attachable) {
	var cParent Attachable
	logs.Debug("detach", child.ID())
	switch c := child.(type) {
	case *eImage:
		cParent, _ = c.parent.Delete()
	case *eText:
		cParent, _ = c.parent.Delete()
	default:
		logs.Fatalf("invalid attachable type: %T", c)
	}

	logs.Debug("detached", child.ID())
	if cParent != nil {
		switch p := cParent.(type) {
		case *eImage:
			p.children.FindAndDelete(func(c Attachable) bool {
				return c.ID() == child.ID()
			})
		case *eText:
			p.children.FindAndDelete(func(c Attachable) bool {
				return c.ID() == child.ID()
			})
		default:
			logs.Fatalf("invalid attachable type: %T", cParent)
		}
	}
}
