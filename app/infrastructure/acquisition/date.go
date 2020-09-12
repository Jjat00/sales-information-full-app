package acquisition

import (
	"strconv"
	"time"
)

func getDate() string {
	t := time.Now()
	date := t.UTC().String()[:10]
	return date
}

func UnixTimeToDateString(unixTime string) string {
	var dateInput string
	if unixTime != "" {
		unix, _ := strconv.ParseInt(unixTime, 10, 64)
		date := time.Unix(unix, 0).UTC().String()
		t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", date)
		dateInput = t.UTC().String()[:10]
	} else {
		dateInput = getDate()
	}
	return dateInput
}
