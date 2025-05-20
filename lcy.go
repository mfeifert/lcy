package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	yearStr := os.Args[1]
	year, _ := strconv.Atoi(yearStr)
	var events []string

	events = append(events, stNicholas.makeEvent(year-1))
	events = append(events, christmas.makeEvent(year-1))
	events = append(events, newYearsDay.makeEvent(year))
	events = append(events, epiphany.makeEvent(year))
	events = append(events, purification.makeEvent(year))
	events = append(events, annunciation.makeEvent(year))
	events = append(events, mayDay.makeEvent(year))
	events = append(events, stJohn.makeEvent(year))
	events = append(events, visitation.makeEvent(year))
	events = append(events, stMichael.makeEvent(year))
	events = append(events, reformation.makeEvent(year))

	for _, event := range events {
		fmt.Println(event)
	}

}
