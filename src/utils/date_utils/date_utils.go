package date_utils

import "time"

const (
	API_DATE_LAYOUT = "2006-01-02T15:04:05Z"
	API_DB_LAYOUT   = "2006-01-02 15:04:05"
)

// GetNow get date now
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString convert date to String
func GetNowString() string {
	return GetNow().Format(API_DATE_LAYOUT)
}

//GetNowDataBaseFormat convert date to String
func GetNowDataBaseFormat() string {
	return GetNow().Format(API_DB_LAYOUT)
}
