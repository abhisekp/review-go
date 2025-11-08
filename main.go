package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"review-go/add_asm"
	"review-go/bfs"
	"review-go/builder"
	. "review-go/dfs"
	"review-go/gcd"
	"review-go/hashtable"
	"review-go/lib"
	"review-go/linkedlist"
	"review-go/pipeline"
	"review-go/recursion"
	"review-go/signal"
	"review-go/ticker"
	. "review-go/utils"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/text/unicode/norm"
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

func main3() {
	Divider()

	concrete := builder.NewConcreteBuilder().SetA(10).SetB(28).Build()
	concrete.Print()
	concrete.Action()
	concrete.Print()

	Divider(DividerOption{Title: "default concrete"})

	concrete2 := builder.GetDefaultConcrete()
	concrete2.Print()
	concrete2.Action()
	concrete2.Print()

	Divider(DividerOption{Title: "base"})

	base := builder.NewBaseBuilder().SetA(10).Build()
	base.Print()

	Divider(DividerOption{Title: "concrete2 build"})

	concrete3 := builder.NewConcrete2Builder().SetB(15).Build()
	// concrete3.Action()
	concrete3.Print()

}

func main4() {
	fmt.Println("Sum =", recursion.Sum([]float64{1, 2.2, 3, 4, 5}))
	fmt.Println("Max =", recursion.Max([]float64{5, 4, 15, 45, 42, 14, 45}))

	sortedNums := recursion.QuickSort([]int{5, 4, 15, 45, 42, 14, 45})
	fmt.Println("QuickSort =", sortedNums)
	fmt.Println("BinarySearch =", recursion.BinarySearch(sortedNums, 15))
	fmt.Println("BinarySearch =", recursion.BinarySearch([]int{1}, 1))
	fmt.Println("Hash4 =", hashtable.Hash4("bag", 10))
	fmt.Println("Fibs =", slices.Collect(Take(50, Fibs())))
}

func main5() {
	bob := bfs.Person{Name: "Bob", ID: "bob"}
	alice := bfs.Person{Name: "Alice", ID: "alice"}
	claire := bfs.Person{Name: "Claire", ID: "claire"}
	you := bfs.Person{Name: "Abhisek", ID: "you"}
	_ = you
	anuj := bfs.Person{Name: "Anuj", ID: "anuj"}
	peggy := bfs.Person{Name: "Peggy", ID: "peggy"}
	jonny := bfs.Person{Name: "jonny", ID: "jonny"}
	thom := bfs.Person{Name: "Thom", ID: "thom", Seller: true}

	graph := map[bfs.Person][]bfs.Person{
		you:    {bob, alice, claire},
		bob:    {anuj, peggy},
		claire: {jonny, thom},
		alice:  {peggy},
	}

	if seller, level, found := bfs.ShortestPath(graph, you, func(person bfs.Person) bool {
		return person.Seller
	}); found {
		fmt.Println("Seller:", seller.Name)
		fmt.Println("Level", level)
	}
}

func main6() {
	graph := map[string][]string{
		"wakeup":        {},
		"shower":        {"exercise"},
		"brush-teeth":   {"wakeup"},
		"eat-breakfast": {"brush-teeth"},
		"exercise":      {"wakeup"},
		"get-dressed":   {"shower"},
		"pack-lunch":    {"wakeup"},
	}

	// start := "wakeup"
	end := "get-dressed"

	tasks := make([]string, 0, len(graph))

	_, level, _ := bfs.ShortestPath(graph, end, func(task string) bool {
		tasks = append(tasks, task)
		return false
	})

	slices.Reverse(tasks)
	fmt.Println(tasks)
	fmt.Println("Level:", level)
}

func main7() {
	graph := []FileNode{
		{
			IsFile: true,
			Name:   "A.jpg",
		}, {
			Name:   "B.jpg",
			IsFile: true,
		},
		{
			Name:        "apps",
			IsDirectory: true,
			Children: []FileNode{
				{
					Name:   "darth.exe",
					IsFile: true,
				},
				{
					IsDirectory: true,
					Name:        "secrets",
					Children: []FileNode{
						{
							IsFile: true,
							Name:   "secretA.exe",
						},
					},
				},
			},
		},
	}

	myDFS := NewDFS[FileNode]()
	myDFS.Run(graph, nil, func(node FileNode, ID string) (string, bool) {
		id := fmt.Sprintf("%s%s", ID, node)
		fmt.Println(id)
		return id, node.IsDirectory
	})
}

func main8() {
	// Alphabet based fb data store
	fbData := make([]*linkedlist.LinkedList[bfs.Person], 26)

	// Initialize FB data array
	for i := range len(fbData) {
		fbData[i] = linkedlist.NewLinkedList[bfs.Person]()
	}

}

func main9() {
	length := 1680
	breadth := 640

	fmt.Println("Smallest Square Size =", FindSmallestSquareSize(length, breadth))
}

func FindSmallestSquareSize(length, breadth int) int {
	return gcd.GCD(length, breadth)
}

func main10() {
	fmt.Println("1 + 2 =", add_asm.Add(1, 2))

}

func main11() {
	fmt.Println("Ticker Timer")
	for i, t := range ticker.Tick(1*time.Second, 15) {
		timeStr := t.Format(time.RFC3339)
		fmt.Println("Tick:", i+1, ".", timeStr)
	}
}

func main12() {
	Panic()
}

func Panic() {
	BusinessError := fmt.Errorf("buinessError: %w", errors.New("something went wrong"))
	_ = BusinessError

	num := 50

	defer func(num int) {
		fmt.Println(num)

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered from panic 2:", err)
			}
		}()

		if err := recover(); err != nil {
			if _err, ok := err.(error); ok && errors.Is(_err, BusinessError) {
				fmt.Println(err)
				return
			}
			fmt.Println("Recovered from panic 1:", err)
			panic(err)
		}
	}(num)

	num = rand.Intn(10)
	fmt.Println(num)
	if num >= 4 {
		panic(BusinessError)
	} else {
		panic("This is non business panic!")
	}
}

type Point struct {
	X    int
	Y    int
	Next *Point
}

func (p *Point) Add(other Point) Point {
	p.X = other.X
	p.Y = other.Y
	p.Next = other.Next
	return Point{
		X:    p.X + other.X,
		Y:    p.Y + other.Y,
		Next: p.Next,
	}
}

func main13() {
	p := Point{1, 2, &Point{}}
	fmt.Println(p.Add(Point{10, 20, nil}))
	fmt.Println(p)

	fmt.Println((*Point).Add(&p, Point{10, 20, nil}))
	fmt.Println(p)
}

func SelectBreak() {
	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		done <- struct{}{}
		close(done)
	}()

locked:
	for {
		if _, err := os.Stat("/tmp/lockfile"); !os.IsNotExist(err) {
			break locked
		}

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Lock file not found, retrying...")
		case <-done:
			fmt.Println("Process is done, exiting...")
			break locked
		}
	}
}

type cursedZST = [0]map[struct{}]struct{} // don't do this.

func main14() {
	// SelectBreak()

	arr := []struct {
		a, b int
		c    []string
	}{
		{},
		{1, 2, []string{"test"}},
		{},
	}
	_ = arr

	// fmt.Println(arr)
	var x cursedZST
	fmt.Println(x)
}

func main15() { // https://go.dev/play/p/4H7V_kKDw5m
	type Point struct{ X, Y, Z uint16 }
	type PaddedPoint struct {
		_       struct{}
		X, Y, Z uint16
		_       struct{}
	}
	const format = "%12v\t%3v\t%6v\n"
	fmt.Printf(format, "type", "size", "align")
	fmt.Printf(format, strings.Repeat("-", len("type")), strings.Repeat("-", len("size")), strings.Repeat("-", len("align")))
	fmt.Printf(format, "Point", unsafe.Sizeof(Point{}), unsafe.Alignof(Point{}))
	fmt.Printf(format, "PaddedPoint", unsafe.Sizeof(PaddedPoint{}), unsafe.Alignof(PaddedPoint{}))
}

type CharCount struct {
	char  string
	count int
}

func main16() {
	input := "👨‍👨‍👦👨‍👨‍👦👨‍👨‍👦aaabbbcca👫👫"

	fmt.Println(input)
	s := norm.NFC.String(input)

	var res []CharCount

	for _, r := range s {
		fmt.Println(strconv.QuoteRune(r))
		if len(res) == 0 {
			res = append(res, CharCount{
				char:  string(r),
				count: 1,
			})
		} else {
			// Get the first character of the string as rune
			// and compare it with the current rune

			if bytes.Runes([]byte(res[len(res)-1].char))[0] == r {
				res[len(res)-1].count++
			} else {
				res = append(res, CharCount{
					char:  string(r),
					count: 1,
				})
			}
		}
	}

	fmt.Println(res)
}

func main17() {
	graph := map[string][]string{
		"cab": {"car", "cat"},
		"cat": {"mat", "bat"},
		"car": {"bar", "cat"},
		"bat": {},
		"mat": {"bat"},
		"bar": {"bat"},
	}
	if _, level, ok := bfs.ShortestPath(graph, "cab", func(s string) bool {
		return s == "bat"
	}); ok {
		fmt.Println("Level:", level)
	}
}

func main18() {
	signal.SignalCtx()
}

func main19() {
	i := 0
	cond := sync.NewCond(&sync.Mutex{})
	cond2 := sync.NewCond(&sync.Mutex{})
	t := time.NewTicker(time.Second)

	// Consumer
	go func() {

		for {
			cond.L.Lock()
			cond.Wait()
			if i > 5 {
				fmt.Println("Caught", i)
			}

			if i == 9 {
				t.Stop()
				cond2.Signal()
			}

			i = 0
			cond.L.Unlock()
		}
	}()

	// Producer
	go func() {
		// Every 1s, a random number appears
		for range t.C {
			cond.L.Lock()
			i = rand.Intn(10)
			cond.L.Unlock()

			fmt.Println("Emit", i)
			cond.Signal()
		}
	}()

	cond2.L.Lock()
	defer cond2.L.Unlock()
	cond2.Wait()
}

func main() {
	pipeline.Run()
}
