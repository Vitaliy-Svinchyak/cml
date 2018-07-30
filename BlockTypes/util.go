package BlockTypes

import "regexp"

func numberOrPercent(value string) bool {
	matched, err := regexp.MatchString("[0-9]+(pt)|(%)", value)

	if err != nil {
		panic(err)
	}

	return matched
}
