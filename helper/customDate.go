package helper

import (
	"fmt"
	"time"
)

// CustomDate for parsing "YYYY-MM-DD" format
type CustomDate struct {
	time.Time
}

// UnmarshalJSON to parse "YYYY-MM-DD" format
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == `""` || str == "null" { // Handle empty or null values
		cd.Time = time.Time{}
		return nil
	}

	// Remove double quotes from JSON string
	str = str[1 : len(str)-1]
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return fmt.Errorf("invalid date format, use YYYY-MM-DD")
	}
	cd.Time = parsedTime
	return nil
}
