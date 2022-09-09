package brtime

import "time"

type Weekday = time.Weekday

var weekdayNames = [...]string{
	time.Sunday:    "Domingo",
	time.Monday:    "Segunda-feira",
	time.Tuesday:   "Terça-feira",
	time.Wednesday: "Quarta-feira",
	time.Thursday:  "Quinta-feira",
	time.Friday:    "Sexta-feira",
	time.Saturday:  "Sábado",
}

func WeekdayName(weekday time.Weekday) string {
	if weekday >= time.Sunday && weekday <= time.Saturday {
		return weekdayNames[weekday]
	}

	return weekday.String()
}
