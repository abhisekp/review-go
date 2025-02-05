package builder

type IBase interface {
	Setup() // Already implemented in base
	Print()
	GetA() int
	IsBuilt() bool
	AssertBuild()
}

// IConcrete enforces both mandatory and implemented methods
type IConcrete interface {
	IBase
	Action() // Must be implemented by child struct
	GetB() int
}

type IConcrete2 interface {
	IConcrete
	GetC() int
}
