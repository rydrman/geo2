# geo2

geo2 is a small library for doing 2d geometry calculations in go

## Installation
~~~~
go get github.com/rydrman/geo2
~~~~

## Simple Examples
```go
import "github.com/rydrman/geo2"

point := geo2.NewVector(5, 5)

path := &geo2.Path{
	geo2.NewVector(1, 0),
	geo2.NewVector(0, 1),
	geo2.NewVector(1, 1),
	geo2.NewVector(0, 0),
}

triangles = path.Triangulate()
```