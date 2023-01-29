package utils

import (
	"api/types"
	"time"
)

func BulkStrToDate(dates ...string) ([]time.Time, error) {
	result := []time.Time{}
	for _, date := range dates {
		parsed, err := time.Parse(types.TIME_FORMAT, date)
		if err != nil {
			return nil, err
		}
		result = append(result, parsed)
	}
	return result, nil
}
