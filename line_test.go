package geo2

import "fmt"

func ExampleGetPerc() {
	line := &Line{
		&Vector{1, 1},
		&Vector{2, 2},
	}
	//this point is 75% of the way from line.A to line.B
	fmt.Println(line.GetPerc(&Vector{1.75, 1.75}))
	// Output:
	// 0.75
}

func ExampleIntersection() {
	line1 := &Line{
		&Vector{0, 0},
		&Vector{2, 2},
	}
	Line := &Line{
		&Vector{0, 2},
		&Vector{2, 0},
	}
	//these lines intersect exacly at (1, 1)
	fmt.Println(line1.Intersection(Line, false))
	// Output:
	// {x: 1.0000, y: 1.0000}
}
