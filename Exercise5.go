
package main
import (
	"fmt"
	"os"
	"strconv"
)
type geometry interface {
	area() float64
	perimeter() float64
}
func result(g geometry) {
	area := g.area()
	fmt.Println("Area:", area)
	perimeter := g.area()
	fmt.Println("Perimeter:", perimeter)
}

type Circle struct{
	radius float64
}
type Rectangle struct
{
	length float64
	breadth float64
}
var circle = Circle{}
var rectangle = Rectangle{}
func main(){
	shape := os.Args[2]
	if shape == "circle" {
		//var radius
		var radius, _  = strconv.ParseFloat(os.Args[3],64)
		circle=Circle{radius}
		result(circle)
	} else
	{
		var length,_=strconv.ParseFloat(os.Args[3],64)
		var breadth,_=strconv.ParseFloat(os.Args[4],64)
		rectangle=Rectangle{length,breadth}
		result(rectangle)
	}
}
func (rectangle Rectangle)area() float64{
	return rectangle.length*rectangle.breadth
}
func (rectangle Rectangle)perimeter() float64{
	return 2*(rectangle.length+rectangle.breadth)
}
func (circle Circle)area() float64{
	//fmt.Println("Area of Circle is: ",3.14*circle.radius*circle.radius)
	return circle.radius*3.14*3.14
}
func (circle Circle)perimeter() float64{
	return circle.radius*2*3.14
}