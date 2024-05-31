package day_27

import "strings"

type Planet string

const earthYear = 31557600

var factors = map[Planet]float64{
	"mercury": 0.2408467,
	"venus":   0.61519726,
	"earth":   1.0,
	"mars":    1.8808158,
	"jupiter": 11.862615,
	"saturn":  29.447498,
	"uranus":  84.016846,
	"neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	planet = Planet(strings.ToLower(string(planet)))
	_, ok := factors[planet]
	if !ok {
		return -1.0
	}
	age := seconds / earthYear * (1 / factors[planet])
	return age
}
