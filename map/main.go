package main
import "fmt"
func main(){
	colors:=map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
	}
	fmt.Println(colors)

	//other ways
	// var colors map[string]string
	colors2:=make(map[string]string)
	colors2["white"]="#ffffff"
	colors2["black"]="#000000"
	fmt.Println(colors2)
	delete(colors2,"black")
	fmt.Println(colors2)
	printMap(colors)
}

func printMap(c map[string]string){
	for color,hex:=range c{
		fmt.Println("Hex code for",color,"is",hex)
	}
}