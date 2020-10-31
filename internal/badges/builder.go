package badges

import "io"

type Builder struct {
	Badger Badger
	Logic  Logicer
}

func NewBuilder(b Badger, l Logicer) *Builder {
	return &Builder{
		Badger: b,
		Logic:  l,
	}
}

func (b *Builder) Build(o io.Writer) {
	svg := b.Badger.Build(b.Logic.GetParams())
	svg.Build(o)
}
