package main

import(
	. "./globalidentity"
	"fmt"
)

func main(){
	a := GlobalIdentity()
	fmt.Println(a)
}