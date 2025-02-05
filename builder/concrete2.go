package builder

import "fmt"

type concrete = Concrete

type Concrete2 struct {
	_         struct{}
	*concrete // Embed Concrete in Concrete2
	__builder bool
	__this    IConcrete2
	a         int
	b         int
	c         int
}

var _ IConcrete2 = (*Concrete2)(nil)

func NewConcrete2() *Concrete2 {
	newConcrete := GetDefaultConcrete()
	this := &Concrete2{concrete: newConcrete}
	newConcrete.__this = this
	return this
}

func (c *Concrete2) GetA() int {
	c.AssertBuild()
	if c.__this != nil {
		return c.__this.GetA()
	}
	return c.a
}

func (c *Concrete2) GetC() int {
	c.AssertBuild()
	if c.__this != nil {
		return c.__this.GetC()
	}
	return c.c
}

func (c *Concrete2) GetB() int {
	c.AssertBuild()
	if c.__this != nil {
		return c.__this.GetB()
	}
	return c.b
}

func (c *Concrete2) Setup() {
	c.AssertBuild()
	c.c = 10
	fmt.Println("Setup method from Concrete2")
}

func (c *Concrete2) IsBuilt() bool {
	return c.__builder && c.concrete != nil && c.concrete.IsBuilt()
}

func (c *Concrete2) AssertBuild() {
	if !c.IsBuilt() {
		panic("Concrete2 must be constructed using NewConcrete2Builder")
	}
}

// Action Implements the required MandatoryMethod from IConcrete
// Adds 10 to a
// Adds 28 to b
func (c *Concrete2) Action() {
	fmt.Println("Action method implemented in Concrete2 [a + 12, b + 15, c + 13]")
	c.a = c.GetA() + 12
	c.b = c.GetB() + 15
	c.c = c.GetC() + 13
}

func (c *Concrete2) Print() {
	c.AssertBuild()
	c.concrete.Print()                                    // Calling base's Print method
	fmt.Println("Printing using Concrete: c =", c.GetC()) // Printing Concrete's value, not base's
}
