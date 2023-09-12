package main

import "fmt"
type shape interface{
	getArea() float64
}
type triangle struct{
	base float64
	height float64
}
type square struct{
	side float64
}

func (t triangle) getArea() float64{
	return 0.5*t.base*t.height
}
func (s square) getArea() float64{
	return s.side*s.side
}

func printArea(sh shape){
	fmt.Println(sh.getArea())
}

func main(){
	tr:=triangle{base:43.0,height: 67.98}
	sq:=square{side:4}
	printArea(tr)
	printArea(sq)
}