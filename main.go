package main

import "fmt"


func main(){
	todos := Todos{}

	todos.add("buy milk")
	todos.add("buy bread")

	fmt.Printf("%+v\n\n",todos)

	todos.delete(0)

	fmt.Printf("%+v",todos)
}