package main

//interface are the collection of method signatures.
// a type implements an interface by implementing the methods.
// there is no explicit declaration of intent, no "implements" keyword.
import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func interfaceShape(s shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}
