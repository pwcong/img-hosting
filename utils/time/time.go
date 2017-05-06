package time

import "time"

func FormatTime(format string) string {
	return time.Now().Format(format)
}

func Now() time.Time {
	return time.Now()
}

// GetDate return a date. eg 2006-01-02
func GetDate() string {

	return FormatTime("2006-01-02")
}

// GetTime return a time. eg 15:04:05.000000
func GetTime() string {
	return FormatTime("15:04:05.000000")
}

// GetDateTime return a datetime. eg 2006-01-02 15:04:05.000000
func GetDateTime() string {
	return FormatTime("2006-01-02 15:04:05.000000")
}

func GetYear() string {
	return FormatTime("2006")
}

func GetMonth() string {
	return FormatTime("01")
}

func GetDay() string {
	return FormatTime("02")
}
