package main

import (
	"fmt"
	"review-go/builder"
	"review-go/lib"
	. "review-go/utils"
)

func main2() {
	// Enforcing interface compliance
	var obj lib.IConcrete = lib.NewConcrete()

	if obj != nil {
		// Calling methods
		obj.Setup()  // Output: Setup method from base and Concrete
		obj.Action() // Output: Action method implemented in Concrete
		obj.Print()  // Output: Printing using Concrete: 15 38
		fmt.Println("Is Internal?", obj.IsBuilt())
	}

	var obj2 = lib.NewBase()
	fmt.Println("Is Internal?", obj2.IsBuilt())

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	obj2.Print() // panic
}

func main() {
	/*Divider()

	concrete := builder.NewConcreteBuilder().SetA(10).SetB(28).Build()
	concrete.Print()
	concrete.Action()
	concrete.Print()*/

	Divider(DividerOption{Title: "default concrete"})

	concrete2 := builder.GetDefaultConcrete()
	concrete2.Print()
	concrete2.Action()
	concrete2.Print()

	Divider(DividerOption{Title: "base"})

	base := builder.NewBaseBuilder().SetA(10).Build()
	base.Print()

	Divider(DividerOption{Title: "concrete2 build"})

	concrete3 := builder.NewConcrete2Builder().SetA(8).SetB(15).SetC(17).Build()
	concrete3.Action()
	concrete3.Print()

}
