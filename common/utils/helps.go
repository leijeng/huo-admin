package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func MapKeyInIntSlice[T int | string](haystack []T, needle T) bool {

	set := make(map[T]struct{})
	for _, e := range haystack {
		set[e] = struct{}{}
	}

	_, ok := set[needle]

	return ok
}

func CalculateTwoDateBetweenDay(format, dateOne, dateTwo string) (int, error) {

	date1, err := time.ParseInLocation(format, dateOne, time.Local)

	if err != nil {
		return -1, err
	}

	date2, err := time.ParseInLocation(format, dateTwo, time.Local)

	if err != nil {
		return -1, err
	}

	var lessHour = date1.Sub(date2).Hours()

	if lessHour < 0 {
		return -1, nil
	}

	return int(date1.Sub(date2).Hours() / 24), nil
}

func StoreSixNum(data float64) float64 {

	if data == 0 {
		return 0
	}

	res, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", data), 64)

	return res
}

func StoreTwoNum(data float64) float64 {

	if data == 0 {
		return 0
	}

	res, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", data), 64)

	return res
}

func Decimal64(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func StructToJsonString(data interface{}) string {
	dataByte, err := json.Marshal(data)

	if err != nil {
		return ""
	}

	return string(dataByte)
}
