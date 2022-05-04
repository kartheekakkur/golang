// You can edit this code!
// Click here and start typing.
package main

import "fmt"


type circle struct {
	radius int

}

type rectangle struct {
	length  int
	breadth int
}

type geometry interface {
	  area() int
	  peri() int

}

func main() {


	c1 := circle {radius:2,}
	r1 := rectangle { 
	   length :2,
	   breadth: 2,
	}
	
  c1.area()
  print(r1)
	
}

// Takes in interface to perform a function

func print(g geometry){

  fmt.Println("Area of the geomerty",g.area())
  fmt.Println("Peri of the geomerty",g.peri())
}

func (c circle) area() int {

	return c.radius * c.radius
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func (r rectangle) peri() int{
	return 2*(r.breadth+r.length)
}