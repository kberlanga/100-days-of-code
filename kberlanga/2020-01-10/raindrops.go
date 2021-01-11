package raindrops

import (
	"math"
	"strconv"
)

func Convert(drops int) string {
	result := ""
	if math.Mod(float64(drops), 3) == 0 {
		result += "Pling"
	}
	if math.Mod(float64(drops), 5) == 0 {
		result += "Plang"
	}
	if math.Mod(float64(drops), 7) == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(drops)
	}
	return result
}
