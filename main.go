package main

import (
	"fmt"
	"time"
)

func main() {
	var start = time.Now()
	timeExample()
	time.Sleep(200 * time.Millisecond)
	var time = time.Now()

	fmt.Println("Duration", time.Sub(start)) //difference
}

func timeExample() {
	var now = time.Now() //current local time
	fmt.Println("Now", now)

	var then = time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	//year int, month Month, day, hour, min, sec, nsec int, loc *Location
	fmt.Println("Then", then)

	fmt.Println("Year", then.Year())
	fmt.Println("Month", then.Month())
	fmt.Println("Day", then.Day())

	//True - False
	fmt.Println("Before", then.Before(now)) //report previous time
	fmt.Println("After", then.After(now))   //report later time
	fmt.Println("Equal", then.Equal(now))   //report equivalent time
}
