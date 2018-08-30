package space

var (
	earth     = "Earth"
	earthYear = float64(31557600)
	planets   = map[string]float64{
		"Mercury": float64(0.2408467),
		"Venus":   float64(0.61519726),
		"Mars":    float64(1.8808158),
		"Jupiter": float64(11.862615),
		"Saturn":  float64(29.447498),
		"Uranus":  float64(84.016846),
		"Neptune": float64(164.79132),
	}
)

func Age(seconds float64, planet string) float64 {
	if planet == earth {
		return seconds / earthYear
	}
	return (seconds / earthYear) / planets[planet]
}
