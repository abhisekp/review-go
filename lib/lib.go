package lib

import "fmt"

type IBase interface {
	Setup() // Already implemented in base
	Print()
	GetA() int
	IsBuilt() bool
}

// IConcrete enforces both mandatory and implemented methods
type IConcrete interface {
	IBase
	Action() // Must be implemented by child struct
	GetB() int
}

// base struct provides an implementation for Setup
type base struct {
	_        struct{}
	internal bool
	a        int
	this     IBase
}

var _ IBase = (*base)(nil)

func NewBase() IBase {
	return &base{}
}

func (b *base) IsBuilt() bool {
	return b.internal
}

func (b *base) Setup() {
	fmt.Println("Setup method from base")
	b.a = 5
}

func (b *base) GetA() int {
	return b.a
}

func (b *base) Print() {
	if !b.IsBuilt() {
		panic("Base must be constructed internally")
	}
	a := b.GetA()
	if b.this != nil {
		a = b.this.GetA()
	}
	fmt.Println("Printing from base: a =", a)
}

// Concrete struct embedding base and implementing MandatoryMethod
// Use NewConcrete() to create a concrete
type Concrete struct {
	_     struct{}
	*base // Pointer embedding allows overriding if needed
	a     int
	b     int
}

var _ IConcrete = (*Concrete)(nil)

func NewConcrete() *Concrete {
	if newBase, ok := NewBase().(*base); ok && newBase != nil {
		newBase.internal = true
		this := &Concrete{base: newBase}
		newBase.this = this
		return this // Returning a pointer to Concrete struct with base embedded
	}
	return nil
}

// Action Implements the required MandatoryMethod from IConcrete
func (c *Concrete) Action() {
	fmt.Println("Action method implemented in Concrete")
	c.a = c.GetA() + 10
	c.b = c.GetB() + 28
}

func (c *Concrete) GetB() int {
	return c.b
}

func (c *Concrete) GetA() int {
	return c.a
}

// Setup overrides base's ImplementedMethod
func (c *Concrete) Setup() {
	c.base.Setup()
	// c.a = 20 // Overriding base's value
	c.b = 10
	fmt.Println("Setup method from Concrete")
}

func (c *Concrete) Print() {
	c.base.Print()                                        // Calling base's Print method
	fmt.Println("Printing using Concrete: b =", c.GetB()) // Printing Concrete's value, not base's
}
