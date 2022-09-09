package brtime

import "time"

type Month = time.Month

var monthNames = [...]string{
	time.January:   "Janeiro",
	time.February:  "Fevereiro",
	time.March:     "Março",
	time.April:     "Abril",
	time.May:       "Maio",
	time.June:      "Junho",
	time.July:      "Julho",
	time.August:    "Agosto",
	time.September: "Setembro",
	time.October:   "Outubro",
	time.November:  "Novembro",
	time.December:  "Dezembro",
}

func MonthName(month time.Month) string {
	if month >= time.January && month <= time.December {
		return monthNames[month]
	}

	return month.String()
}
