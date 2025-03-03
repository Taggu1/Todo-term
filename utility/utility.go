package utility

import (
	"fmt"
	"time"
)

func RemoveSliceElement[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func FormatedNowTime() string {
	t := time.Now()

	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute())

	return formatted
}
