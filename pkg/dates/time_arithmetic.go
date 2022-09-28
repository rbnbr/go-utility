package dates

import "time"

// RoundToBeginningOfDay
// Given a time object, this function returns the same time object but drops all information except
// Year, Month, Day, and Location.
// Returns the same Day at 00h:00m:00s:0000ns
func RoundToBeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// RoundToEndOfDay
// Given a time object, this function returns the same time object but changes all information except
// Year, Month, Day, and Location to the very last measurable point in time for this day same day.
// Returns the same Day at 23h:59m:59s:999999999ns
func RoundToEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 1e9-1, t.Location())
}

// SetTimeZone
// Returns a new time object which has the same values as t except its location has changed.
// I.e., providing a time at 07:00 in Berlin and calling this function with this time and a location in New York
// 	will return a time at 07:00 in New York.
func SetTimeZone(t time.Time, loc *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}
