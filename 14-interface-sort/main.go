package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type earthMass float64
type au float64

type Planet struct {
	name     string
	mass     earthMass
	distance au
}

type By func(p1, p2 *Planet) bool

func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by,
	}
	sort.Sort(ps)
}

type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool
}

func (s *planetSorter) Len() int {
	return len(s.planets)
}

func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

type Change struct {
	user     string
	language string
	lines    int
}

type lessFunc func(p1, p2 *Change) bool

type multiSorter struct {
	changes []Change
	less    []lessFunc
}

func (ms *multiSorter) Sort(changes []Change) {
	ms.changes = changes
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)

}

func main() {

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	var planets = []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}

	sortByName := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	sortByMess := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	sortByDistance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	sortByDistanceDesc := func(p1, p2 *Planet) bool {
		return !sortByDistance(p1, p2)
	}

	By(sortByName).Sort(planets)
	fmt.Println("name", planets)

	By(sortByMess).Sort(planets)
	fmt.Println("mess", planets)

	By(sortByDistance).Sort(planets)
	fmt.Println("distance", planets)

	By(sortByDistanceDesc).Sort(planets)
	fmt.Println("distance???", planets)

	ints := []int{5, 4, 3, 2, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	var changes = []Change{
		{"gri", "Go", 100},
		{"ken", "C", 150},
		{"glenda", "Go", 200},
		{"rsc", "Go", 200},
		{"r", "Go", 100},
		{"ken", "Go", 200},
		{"dmr", "C", 100},
		{"r", "C", 150},
		{"gri", "Smalltalk", 80},
	}
	user := func(c1, c2 *Change) bool {
		return c1.user < c2.user
	}
	language := func(c1, c2 *Change) bool {
		return c1.language < c2.language
	}
	increasingLines := func(c1, c2 *Change) bool {
		return c1.lines < c2.lines
	}
	decreasingLines := func(c1, c2 *Change) bool {
		return c1.lines > c2.lines // Note: > orders downwards.
	}
	OrderedBy(user).Sort(changes)
	fmt.Println("By user:", changes)

	OrderedBy(user, increasingLines).Sort(changes)
	fmt.Println("By user,<lines:", changes)
	OrderedBy(user, decreasingLines).Sort(changes)
	fmt.Println("By user,>lines:", changes)
	OrderedBy(language, increasingLines).Sort(changes)
	fmt.Println("By language,<lines:", changes)
	OrderedBy(language, increasingLines, user).Sort(changes)
	fmt.Println("By language,<lines,user:", changes)

}
