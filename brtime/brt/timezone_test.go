package brt

import (
	"testing"
	"time"
)

func TestTimezoneBRT(t *testing.T) {
	testBRTLocation, err := time.LoadLocation(LocationName)
	if err != nil {
		t.Error("Failed to load the BRT location")
		t.FailNow()
	}

	const timeDifference = 3 * time.Hour

	timezoneFactory := func(location *time.Location) time.Time {
		return time.Date(2022, time.September, 1, 19, 11, 46, 0, location)
	}

	utc0Time := timezoneFactory(time.UTC).Add(timeDifference)
	brtTime := timezoneFactory(testBRTLocation)
	if !BatchConvertToBRT(&utc0Time, &brtTime) {
		t.Error("An error occurred during timezone conversion")
	}

	controlTime := timezoneFactory(testBRTLocation)
	utcControlTime := timezoneFactory(time.UTC).Add(timeDifference)
	for _, toCheck := range [...]time.Time{utc0Time, brtTime} {
		if !controlTime.Equal(toCheck) {
			t.Errorf("Time collision with time: `%v`", toCheck)
		}

		actualDifference := utcControlTime.Sub(toCheck)
		if actualDifference != 0 {
			t.Errorf("The given time difference was not 0s\nReceived: `%v`", actualDifference)
		}
	}
}

func TestFormat(t *testing.T) {
	testCases := [...]struct {
		timeElem time.Time
		expected string
	}{
		{time.Date(2002, time.September, 15, 8, 1, 0, 0, time.UTC), "15/09/2002 - 08:01:00"},
		{time.Date(2002, time.August, 20, 0, 0, 0, 0, time.UTC), "20/08/2002"},
		{time.Date(2002, time.February, 28, 0, 0, 37, 0, time.UTC), "28/02/2002 - 00:00:37"},
		{time.Date(2015, time.October, 8, 4, 49, 1, 0, time.UTC), "08/10/2015 - 04:49:01"},
	}

	for _, tCase := range testCases {
		if result := Format(tCase.timeElem); result != tCase.expected {
			t.Errorf("Formatted time is not the expected\nExpected: `%s`\nReceived: `%s`", tCase.expected, result)
		}
	}
}
