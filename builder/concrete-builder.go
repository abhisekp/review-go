package builder

// ConcreteBuilder enforces creation via Builder
type ConcreteBuilder struct {
	concrete *Concrete
}

func NewConcreteBuilder() *ConcreteBuilder {
	newConcrete := NewConcrete()
	newConcrete.__builder = true
	newConcrete.Setup()
	return &ConcreteBuilder{concrete: newConcrete}
}

func (cb *ConcreteBuilder) SetA(a int) *ConcreteBuilder {
	cb.concrete.a = a
	return cb
}

func (cb *ConcreteBuilder) SetB(b int) *ConcreteBuilder {
	cb.concrete.b = b
	return cb
}

func (cb *ConcreteBuilder) Build() *Concrete {
	return cb.concrete
}

func GetDefaultConcrete() *Concrete {
	return NewConcreteBuilder().Build()
}
