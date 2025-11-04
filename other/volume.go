/*
We calculate and display the volumes of 3d solids: **Spheres**, **Cubes**, and **Pyramids**, using interfaces.

Defined structs for each solid (`Sphere`, `Cube`, `Pyramid`) and an interface called `Solid`. To the spec of the homework we added `Name() string`, to return the solid's name.

Assumed the file is named `./hw3data` as in the example given.

Not using the suggested input type struct because it was not needed.

Results are printed to stdout.
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Sphere struct {
	radius float64
}

type Cube struct {
	length float64
}

type Pyramid struct {
	base   float64
	height float64
}

type Solid interface {
	Volume() float64
	Name() string
}

func (c Cube) Volume() float64 {
	return math.Pow(c.length, 3)
}
func (Cube) Name() string {
	return "Cube"
}

func (s Sphere) Volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(s.radius, 3)
}
func (Sphere) Name() string {
	return "Sphere"
}

func (p Pyramid) Volume() float64 {
	return p.base * p.base * p.height / 3
}
func (Pyramid) Name() string {
	return "Pyramid"
}

func printAreas(solids map[string]Solid) {
	for _, shape := range solids {
		fmt.Println(shape.Name(), shape.Volume())
	}
}

func format(data []string) map[string]Solid {
	var solids = make(map[string]Solid)

	for i, line := range data {
		params := strings.Split(line, " ")
		p1 := params[0]
		p2, _ := strconv.ParseFloat(params[2], 64)
		index := p1 + fmt.Sprint(i)

		switch p1 {
		case "C":
			solids[index] = Cube{length: p2}
		case "S":
			solids[index] = Sphere{radius: p2}
		case "P":
			p3, _ := strconv.ParseFloat(params[3], 64)
			solids[index] = Pyramid{base: p2, height: p3}
		default:
			panic("Unknown type")
		}
	}

	return solids
}

func main() {
	data, _ := os.ReadFile("./hw3data")
	solids := format(strings.Split(string(data), "\n"))
	printAreas(solids)
}
