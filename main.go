package main

import (
	"fmt"
	"GoRobots/NewRobots"
)

func main() {
	fmt.Println("\nHello world")
	token := NewRobots.GenerateToken()
	myRobot := token.CreateNewRobot("astro")
	myRobot.Start()
}