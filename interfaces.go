package main

import (
	"fmt"
	"math"
)

//Geometricer can evaluate geometric operations (composed interface with Stringer)
type Geometricer interface {
	Area() float64
	fmt.Stringer
}

//--------------------

//Square is a geometric form
type Square struct {
	Width, Height float64
}

//Area calculated from Square
func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

//--------------------

//Circle is a geometric form
type Circle struct {
	Radius float64
}

//Area calculated from Circle
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

//--------------------

func main() {

	s := Square{Width: 10, Height: 5}

	c := Circle{Radius: 2}

	//Square operations
	fmt.Println(s.String())
	printArea(s)

	//Circle operations
	fmt.Println(c.String())
	printArea(c)
}

func printArea(g Geometricer) {
	fmt.Printf("Area of %s = %.2f\n", g.String(), g.Area())
}
