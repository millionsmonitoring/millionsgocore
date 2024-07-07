package timing

import (
	"log/slog"
	"time"
)

const (
	Layout           = "01/02 03:04:05PM '06 -0700"
	ANSIC            = "Mon Jan _2 15:04:05 2006"
	UnixDate         = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate         = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822           = "02 Jan 06 15:04 MST"
	RFC822Z          = "02 Jan 06 15:04 -0700"
	RFC850           = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123          = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z         = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339          = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano      = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen          = "3:04PM"
	DateUSA          = "2006-02-01"
	DateIndian       = "02/01/2006"
	Date             = "02 Jan 2006"
	DateIndianDashed = "2006-01-02"

	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

func TimeZone(zone string) *time.Location {
	if zone == "" {
		zone = time.UTC.String()
	}
	location, err := time.LoadLocation(zone)
	if err != nil {
		slog.Error("TimeZone:", err)
	}
	return location
}

func IndiaTimeZone() *time.Location {
	return TimeZone("Asia/Kolkata")
}

func NowIST() time.Time {
	return time.Now().In(IndiaTimeZone())
}

func StartOfDayIST(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, IndiaTimeZone())
}

func EndOfDayIST(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, IndiaTimeZone())
}
