package utils

import (
	"fmt"
	"time"
)

func ConvertType(timeString string) (time.Time, error) {
	const layout = "2006-01-02 15:04:05"

	date, err := time.Parse(layout, timeString)

	if err != nil {
		fmt.Sprintf("Erorr occured while converting datetime: %s. Error: %s", timeString, err)
		fmt.Print(fmt.Sprintf("Erorr occured while converting datetime: %s. Error: %s", timeString, err))
		return time.Time{}, err
	}
	return date, nil
}
