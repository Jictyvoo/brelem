package brt

import "time"

var brtLocation *time.Location

func init() {
	_ = ReloadLocation()
}

// ReloadLocation Tries to reload the BRT location used to create and parse new dates
func ReloadLocation() (err error) {
	brtLocation, err = time.LoadLocation(LocationName)
	return
}

// Date Calls the time.Date function to generate a new time.Time using the BRT-0 location as standard
func Date(day int, month time.Month, year, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, brtLocation)
}
