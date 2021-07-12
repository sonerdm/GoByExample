package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

type Circle struct {
	x, y, r float64
}

// One thing to remember is that arguments are always copied in Go. If we attempted to modify one of the fields inside of the circleArea function,
//it would not modify the original variable. Because of this we would typically write the function like this:
func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

//In between the keyword func and the name of the function we've added a “receiver”.
//The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator:
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

//This would work, but we would rather say an Android is a Person, rather than an Android has a Person. Go supports relationships like this by using an embedded type.
//Also known as anonymous fields, embedded types look like this:
type Android struct {
	Person
	Model string
}

//You may have noticed that we were able to name the Rectangle's area method the same thing as the Circle's area method.
//This was no accident. In both real life and in programming, relationships like these are commonplace. Go has a way of making
//these accidental similarities explicit through a type known as an Interface. Here is an example of a Shape interface:

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

//Interfaces can also be used as fields:
type MultiShape struct {
	shapes []Shape
}

//Now a MultiShape can contain Circles, Rectangles or even other MultiShapes.
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cr float64 = 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cr))
	fmt.Println("-----Struct-----")
	//c := new(Circle) //This allocates memory for all the fields, sets each of them to their zero value and returns a pointer. (*Circle)
	c := Circle{x: 0, y: 0, r: 5}
	c.r = 3
	fmt.Println(c.x) //We can access fields using the . operator:
	fmt.Println(circleArea2(&c))

	fmt.Println("-----Methods-----")
	fmt.Println(c.area())
	//This is much easier to read, we no longer need the & operator (Go automatically knows to pass a pointer
	//to the circle for this method) and because this function can only be used with Circles we can rename the function to just area.
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println("-----Embedded Types-----")
	a := new(Android)
	a.Person.Talk()
	a.Talk()
	//The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk.

	fmt.Println("-----Interfaces-----")
	fmt.Println(totalArea(&c, &r))

	fmt.Println("-----Problem-----")
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}

// 1.Whats's the difference between a method and a function?
// Functions are sections of code that map zero or more input parameters to zero or more output parameters.
// Methods are a special type of functions that are more specific and bounded to a receiver. We can call methods using the . operator.

// 2. Why would you use an embedded anonymous field instead of a normal named field?
// In the case we want to represent inheritance and specificity. An object is a other object.
