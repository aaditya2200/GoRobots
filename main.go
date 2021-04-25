package main

import (
	"GoRobots/NewRobots"
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nHello world")
	token := NewRobots.GenerateToken()
	myRobot := token.CreateNewRobotWithBlinkFunc("astro", "13", 1 * time.Second)
	myRobot.Start()
}