
package main

import (
"fmt"
"os"
"strconv"
)

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
		fmt.Println("Area is", circle.area())
		fmt.Println("Perimeter is",circle.perimeter())

	} else
	{
		var length,_=strconv.ParseFloat(os.Args[3],64)
		var breadth,_=strconv.ParseFloat(os.Args[4],64)
		rectangle=Rectangle{length,breadth}
		fmt.Println("Area is", rectangle.area())
		fmt.Println("Perimeter is",rectangle.perimeter())
	}
}
func (rectangle Rectangle)area() float64{
	//fmt.Println("Area of Circle is: ",3.14*circle.radius*circle.radius)
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