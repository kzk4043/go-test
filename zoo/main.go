package main

import (
	"fmt"
	"zoo/animals"
)

func main() {
	fmt.Println("Rabbit eats:", animals.RabbitFeed())
	fmt.Println("Monkey eats:", animals.MonkeyFeed())
	fmt.Println("Elephant eats:", animals.ElephantFeed())
}