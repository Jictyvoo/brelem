package brt

import "time"

const LocationName = "America/Sao_Paulo"

func Format(element time.Time) string {
	// Check if it has hour/minute/second
	if element.Hour() > 0 || element.Minute() > 0 || element.Second() > 0 {
		return element.Format("02/01/2006 - 15:04:05")
	}
	return element.Format("02/01/2006")
}
