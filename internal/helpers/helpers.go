package helpers

import "strconv"

func ParseToFloat(str string) float64 {
	if s, err := strconv.ParseFloat(str, 64); err == nil {
		return s
	}

	return 0
}
