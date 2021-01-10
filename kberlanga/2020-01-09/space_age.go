package space

import "math"

type Planet string

// Age calcuate space age.
func Age(seconds float64, planet Planet) float64 {
	planets := map[Planet]float64{
		"Mercury": PeroidToSeconds(0.2408467),
		"Venus":   PeroidToSeconds(0.61519726),
		"Earth":   math.Pow(31557600, -1),
		"Mars":    PeroidToSeconds(1.8808158),
		"Jupiter": PeroidToSeconds(11.862615),
		"Saturn":  PeroidToSeconds(29.447498),
		"Uranus":  PeroidToSeconds(84.016846),
		"Neptune": PeroidToSeconds(164.79132),
	}
	return seconds * planets[planet]
}

func PeroidToSeconds(period float64) float64 {
	return math.Pow(31557600*period, -1)
}
