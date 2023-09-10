package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	jim:=person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email:"jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimPointer:=&jim //memory address of jim
	jimPointer.updateName("jimmy") //can call with or without pointer jim.updateName("jimmy")
	jim.print()
}

func basicStruct(){
	// alex:=person{"Alex", "Anderson"} //one of the ways of defining the struct
	alex:=person{firstName: "Alex",lastName:"Anderson"}
	fmt.Println(alex)

	var ani person
	fmt.Println(ani) //prints empty ani
	ani.firstName="Anirudh"
	ani.lastName="Gugan"
	fmt.Printf("%+v",ani) //prints every variable and the value
}

func (p person) print(){
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateName(newFirstName string){ //here the pointer is showing that we are going to use a pointer as a type
	(*pointerToPerson).firstName=newFirstName  //*pointerToPerson take the memory address and give the direct access to the value 
}