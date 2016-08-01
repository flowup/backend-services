package main

import "fmt"
import "time"

func main() {
	var event time.Time = time.Date(2013, 8, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println(calcAge(event))
}

//	Function that is calculating years from the given event to today
func calcAge(event time.Time) int {
	now := time.Now() 
	years := now.Year() - event.Year()

	//	Checks if is going to happen this year but next month or in next days in current month
	if now.Month() < event.Month() {
		years--
	} else if now.Month() == event.Month() && now.Day() < event.Day() {
		years --
	}
	
	return years
}
