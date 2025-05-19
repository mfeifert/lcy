package main

import "fmt"

type event struct {
	name  string
	month int
	day   int
}

func main() {

	newYearsDay := event{
		name:  "New Year's Day",
		month: 1,
		day:   1,
	}

	fmt.Println(newYearsDay.name)

}
