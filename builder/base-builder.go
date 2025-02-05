package builder

// BaseBuilder enforces creation via Builder
type BaseBuilder struct {
	base *base
}

func NewBaseBuilder() *BaseBuilder {
	newBase := &base{__builder: true}
	newBase.Setup()
	return &BaseBuilder{base: newBase}
}

func (bb *BaseBuilder) SetA(a int) *BaseBuilder {
	bb.base.a = a
	return bb
}

func (bb *BaseBuilder) Build() IBase {
	return bb.base
}

func GetDefaultBase() IBase {
	return NewBaseBuilder().Build()
}
