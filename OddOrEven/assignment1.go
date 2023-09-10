package main

import "fmt"

func main(){
	intslice:=[]int{0,1,2,3,4,5,6,7,8,9,10};
	for _,val:=range intslice{
		if(val%2==0){
			fmt.Println("Even")
		}else{
			fmt.Println("Odd")
		}
	}
}