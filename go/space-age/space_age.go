package space

type Planet string

const earthRevolution = float64(31557600)

func Age(seconds float64, planet Planet) float64 {

	// 1.0 Earth-year compared to the others
	PlanetYears := map[Planet]float64{
		"Earth":   1.0 * earthRevolution,
		"Mercury": 0.2408467 * earthRevolution,
		"Venus":   0.61519726 * earthRevolution,
		"Mars":    1.8808158 * earthRevolution,
		"Jupiter": 11.862615 * earthRevolution,
		"Saturn":  29.447498 * earthRevolution,
		"Uranus":  84.016846 * earthRevolution,
		"Neptune": 164.79132 * earthRevolution,
	}

	return calculateAge(seconds, PlanetYears[planet])
}

func calculateAge(seconds float64, planetRevolution float64) float64 {
	return seconds / planetRevolution
}
