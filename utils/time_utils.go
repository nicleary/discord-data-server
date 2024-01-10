package utils

import (
	"fmt"
	"strings"
	"time"
)

func ConvertType(timeString string) (time.Time, error) {
	const layout = "2006-01-02 15:04:05"

	date, err := time.Parse(layout, timeString)

	if err != nil {
		// Clean the user input
		escapedString := strings.ReplaceAll(timeString, "\n", "")
		escapedString = strings.ReplaceAll(escapedString, "\r", "")
		fmt.Printf("Erorr occured while converting datetime: %s. Error: %s", escapedString, err)
		return time.Time{}, err
	}
	return date, nil
}
