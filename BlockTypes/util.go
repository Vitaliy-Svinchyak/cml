package BlockTypes

import (
	"regexp"
	"math/rand"
)

func nmPrKt(value string) bool {
	matched, err := regexp.MatchString("^[0-9]+[%]|(kt)?$", value)

	if err != nil {
		panic(err)
	}

	return matched
}

func ptPrKt(value string) bool {
	matched, err := regexp.MatchString("^[0-9]+(pt)|(%)|(kt)$", value)

	if err != nil {
		panic(err)
	}

	return matched
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
