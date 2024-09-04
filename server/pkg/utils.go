package utils

import (
	"strconv"
	"time"
)

func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

// generateID(req.Username, time.Now()),
func GenerateID(username string, now time.Time) int {
	aa, err := strconv.Atoi(username)
	err = err
	return (aa%1000)*10000 + int(now.Unix())%10000
}
