package utils

import "time"

func Today() string {
	return time.Now().Format("20060102")
}

func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}
