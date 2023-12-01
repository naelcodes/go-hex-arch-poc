package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// function that generate a random string
func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateCode(prefix string, number int) string {
	var code string
	UpperCasePrefix := strings.ToUpper(prefix)

	if number < 10 {
		code = UpperCasePrefix + "000" + strconv.Itoa(number)
	} else if number < 100 {
		code = UpperCasePrefix + "00" + strconv.Itoa(number)
	} else if number < 1000 {
		code = UpperCasePrefix + "0" + strconv.Itoa(number)
	} else {
		code = UpperCasePrefix + strconv.Itoa(number)
	}
	return code

}

func GetCurrentDate() string {
	currentDate := time.Now().Format("2006-01-02")
	return currentDate
}
