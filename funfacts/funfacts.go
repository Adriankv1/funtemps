package funfacts

import (
	"math/rand"
	"time"
)

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	facts := FunFacts{
		Sun: []string{
			"Temperature in the Sun's core is 15,000,000 degrees Celsius",
			"Temperature of the outer layer of the Sun is 5778 Kelvin",
		},
		Luna: []string{
			"Temperature of the Moon's surface at night is -183 degrees Celsius",
			"Temperature of the Moon's surface during the day is 106 degrees Celsius",
		},
		Terra: []string{
			"Highest temperature measured on the Earth's surface is 134 degrees Fahrenheit, 56.7 degrees Celsius, 329.82 Kelvin",
			"Lowest temperature measured on the Earth's surface is -89.4 degrees Celsius",
			"Temperature in the Earth's inner core is 9392 Kelvin",
		},
	}

	var factsList []string

	switch about {
	case "sun":
		factsList = facts.Sun
	case "luna":
		factsList = facts.Luna
	case "terra":
		factsList = facts.Terra
	default:
		return factsList
	}

	// Sett opp en kilde for tilfeldige tall basert på nåværende tidspunkt
	rand.Seed(time.Now().UnixNano())

	// Velg en tilfeldig fun fact fra listen
	index := rand.Intn(len(factsList))
	fact := factsList[index]

	return []string{fact}
}
