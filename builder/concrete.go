package builder

import "fmt"

// Concrete struct embedding base and implementing MandatoryMethod
type Concrete struct {
	_         struct{}
	*base     // Pointer embedding allows overriding if needed
	__builder bool
	__this    IConcrete
	// a         int
	b int
}

var _ IConcrete = (*Concrete)(nil)

func NewConcrete() *Concrete {
	if newBase, ok := NewBaseBuilder().Build().(*base); ok && newBase != nil {
		this := &Concrete{base: newBase}
		newBase.__this = this
		return this // Returning a pointer to Concrete struct with base embedded
	}
	return nil
}

// Action Implements the required MandatoryMethod from IConcrete
// Adds 10 to a
// Adds 28 to b
func (c *Concrete) Action() {
	fmt.Println("Action method implemented in Concrete [a + 10, b + 28]")
	c.a = c.GetA() + 10
	c.b = c.GetB() + 28
}

func (c *Concrete) IsBuilt() bool {
	return c.__builder && c.base != nil && c.base.IsBuilt()
}

func (c *Concrete) AssertBuild() {
	if !c.IsBuilt() {
		panic("Concrete must be constructed using NewConcreteBuilder")
	}
}

func (c *Concrete) GetB() int {
	c.AssertBuild()
	if c.__this != nil {
		return c.__this.GetB()
	}
	return c.b
}

func (c *Concrete) GetA() int {
	c.AssertBuild()
	if c.__this != nil {
		return c.__this.GetA()
	}
	return c.a
}

// Setup overrides base's ImplementedMethod
// Assigns 10 to b
func (c *Concrete) Setup() {
	c.AssertBuild()
	c.b = 10
	fmt.Println("Setup method from Concrete")
}

func (c *Concrete) Print() {
	c.AssertBuild()
	c.base.Print()
	fmt.Println("Printing using Concrete: b =", c.GetB())
}
