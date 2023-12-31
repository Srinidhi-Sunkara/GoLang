package main

import "fmt"
type bot interface{
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
func main(){
	eb:=englishBot{}
	sb:=spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}
// can not write same function twice as go doesnt allow  so implement interface
// func printGreeting(eb englishBot){
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }
func (eb englishBot) getGreeting() string{
	//very custom logic for english greeting
	return "Hi there"
}

func (sb spanishBot) getGreeting() string{
	return "Hola!"
}