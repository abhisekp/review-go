package builder

import "fmt"

// base struct provides an implementation for Setup
type base struct {
	_         struct{}
	__builder bool
	__this    IBase
	a         int
}

var _ IBase = (*base)(nil)

func (b *base) IsBuilt() bool {
	return b.__builder
}

// Setup sets up the BaseBuilder
// Assigns 5 to a
func (b *base) Setup() {
	b.AssertBuild()
	fmt.Println("Setup method from base")
	b.a = 5
}

func (b *base) GetA() int {
	b.AssertBuild()
	if b.__this != nil {
		return b.__this.GetA()
	}
	return b.a
}

func (b *base) AssertBuild() {
	if !b.IsBuilt() {
		panic("Base must be constructed using NewBaseBuilder")
	}
}

func (b *base) Print() {
	b.AssertBuild()
	a := b.GetA()
	if b.__this != nil {
		a = b.__this.GetA()
	}
	fmt.Println("Printing from base: a =", a)
}
