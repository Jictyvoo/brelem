package brt

import "time"

const LocationName = "America/Sao_Paulo"

const (
	DateFormat     = "02/01/2006"
	TimeFormat     = "15:04:05"
	DateTimeFormat = DateFormat + " - " + TimeFormat
)

// BatchConvertToBRT receives a list of addresses and sets the timezone to BRT-0 (which is UTC-3) in every element
func BatchConvertToBRT(timestamps ...*time.Time) bool {
	if utc3Timezone, err := time.LoadLocation(LocationName); err == nil {
		for _, element := range timestamps {
			// It's allowed to receive nil pointers, so check it
			if element != nil {
				*element = element.In(utc3Timezone)
			}
		}
		return true
	}
	return false
}

// AsBRT receives a timestamp and then return a new time.Time in UTC-3 (BRT-0)
func AsBRT(timestamp time.Time) time.Time {
	return timestamp.In(brtLocation)
}

func Format(element time.Time) string {
	// Check if it has hour/minute/second
	if element.Hour() > 0 || element.Minute() > 0 || element.Second() > 0 {
		return element.Format(DateTimeFormat)
	}
	return element.Format(DateFormat)
}
