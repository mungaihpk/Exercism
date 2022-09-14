package main

type Planet string

const earth_period = 31557600.0

func Age(seconds float64, planet Planet) float64 {
	//panic("Please implement the Age function")

	if planet == "Sun" {
		return -1.00
	}

	periods := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	period := periods[planet]
	earth_years := seconds / earth_period

	return earth_years / period
}
