package main

type eventConstant struct {
	name  string
	month int
	day   int
}

var (
	stNicholas = eventConstant{
		name:  "St. Nicholas",
		month: 12,
		day:   6,
	}

	christmas = eventConstant{
		name:  "Christmas Day",
		month: 12,
		day:   25,
	}

	newYearsDay = eventConstant{
		name:  "New Year's Day",
		month: 1,
		day:   1,
	}

	epiphany = eventConstant{
		name:  "Epiphany",
		month: 1,
		day:   6,
	}

	purification = eventConstant{
		name:  "Purification",
		month: 2,
		day:   2,
	}

	annunciation = eventConstant{
		name:  "Annunciation",
		month: 3,
		day:   25,
	}

	mayDay = eventConstant{
		name:  "May Day",
		month: 5,
		day:   1,
	}

	stJohn = eventConstant{
		name:  "St. John",
		month: 6,
		day:   24,
	}

	visitation = eventConstant{
		name:  "Visitation",
		month: 7,
		day:   2,
	}

	stMichael = eventConstant{
		name:  "St. Michael",
		month: 9,
		day:   29,
	}

	reformation = eventConstant{
		name:  "Reformation",
		month: 10,
		day:   31,
	}
)
