package builder

type Concrete2Builder struct {
	concrete *Concrete2
}

func NewConcrete2Builder() *Concrete2Builder {
	newConcrete := NewConcrete2()
	newConcrete.__builder = true
	newConcrete.Setup()
	return &Concrete2Builder{concrete: newConcrete}
}

func (c2b *Concrete2Builder) SetA(a int) *Concrete2Builder {
	c2b.concrete.a = a
	return c2b
}

func (c2b *Concrete2Builder) SetB(b int) *Concrete2Builder {
	c2b.concrete.b = b
	return c2b
}

func (c2b *Concrete2Builder) SetC(c int) *Concrete2Builder {
	c2b.concrete.c = c
	return c2b
}

func (c2b *Concrete2Builder) Build() *Concrete2 {
	return c2b.concrete
}

func GetDefaultConcrete2() *Concrete2 {
	return NewConcrete2Builder().Build()
}
