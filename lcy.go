package main

import (
	"fmt"
	"os"
	"strconv"
)

var Year int

func main() {

	yearStr := os.Args[1]
	year, _ := strconv.Atoi(yearStr)
	Year = year
	var events []string

	events = append(events, stNicholas.makeEvent())
	events = append(events, christmas.makeEvent())
	events = append(events, newYearsDay.makeEvent())
	events = append(events, epiphany.makeEvent())
	events = append(events, purification.makeEvent())
	events = append(events, annunciation.makeEvent())
	events = append(events, mayDay.makeEvent())
	events = append(events, stJohn.makeEvent())
	events = append(events, visitation.makeEvent())
	events = append(events, stMichael.makeEvent())
	events = append(events, reformation.makeEvent())

	for _, event := range events {
		fmt.Println(event)
	}
}
