package regular

import (
	"regexp"
	"strconv"
)

func RemoveNonNumericalInputs(input string) float32 {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}

	number, err := strconv.ParseFloat(reg.ReplaceAllLiteralString(input, ""), 64)
	return float32(number)
}
