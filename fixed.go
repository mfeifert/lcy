package main

// Events that have a fixed date

import (
	"time"
)

type eventFixed struct {
	name  string
	month int
	day   int
}

func (e eventFixed) getDate(year int) time.Time {
	return time.Date(year, time.Month(e.month), e.day, 0, 0, 0, 0, time.Local)
}

func (e eventFixed) makeEvent() string {
	format := "2006-01-02 Mon"
	return e.getDate(Year).Format(format) + " " + e.name
}

var (
	stNicholas = eventFixed{
		name:  "St. Nicholas",
		month: 12,
		day:   6,
	}

	christmas = eventFixed{
		name:  "Christmas Day",
		month: 12,
		day:   25,
	}

	newYearsDay = eventFixed{
		name:  "New Year's Day",
		month: 1,
		day:   1,
	}

	epiphany = eventFixed{
		name:  "Epiphany",
		month: 1,
		day:   6,
	}

	purification = eventFixed{
		name:  "Purification",
		month: 2,
		day:   2,
	}

	annunciation = eventFixed{
		name:  "Annunciation",
		month: 3,
		day:   25,
	}

	mayDay = eventFixed{
		name:  "May Day",
		month: 5,
		day:   1,
	}

	stJohn = eventFixed{
		name:  "St. John",
		month: 6,
		day:   24,
	}

	visitation = eventFixed{
		name:  "Visitation",
		month: 7,
		day:   2,
	}

	stMichael = eventFixed{
		name:  "St. Michael",
		month: 9,
		day:   29,
	}

	reformation = eventFixed{
		name:  "Reformation",
		month: 10,
		day:   31,
	}
)
