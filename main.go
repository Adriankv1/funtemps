package main

import (
	"flag"
	"fmt"
	"github.com/Adriankv1/funtemps/conv"
	"github.com/Adriankv1/funtemps/funfacts"
	"os"
	"strings"
)

var from string
var to string
var temp float64

func init() {
	flag.StringVar(&from, "from", "c", "source temperature scale (c/f/k)")
	flag.StringVar(&to, "to", "f", "target temperature scale (c/f/k)")
	flag.Float64Var(&temp, "temp", 0, "temperature to convert")
}

func main() {

	flag.Parse()

	args := os.Args

	// Sett opp en liste over gyldige verdier for "about"
	aboutValues := []string{"sun", "luna", "terra"}

	// Sjekk om det er to argumenter og det andre argumentet er en gyldig verdi for "about"
	if len(args) == 2 && contains(aboutValues, args[1]) {
		// Hent ut en tilfeldig fun fact om valgt tema
		facts := funfacts.GetFunFacts(args[1])

		// Skriv ut fun facten
		if len(facts) > 0 {
			fmt.Println("Did you know? ", facts[0])
		} else {
			fmt.Println("No fun facts found for", args[1])
		}
	} else {
		from = strings.ToUpper(from)
		to = strings.ToUpper(to)
		var result float64

		switch from {
		case "C":
			if to == "F" {
				result = conv.CelsiusToFahrenheit(temp)
			} else if to == "K" {
				result = conv.CelsiusToKelvin(temp)
			}
		case "F":
			if to == "C" {
				result = conv.FahrenheitToCelsius(temp)
			} else if to == "K" {
				result = conv.FahrenheitToKelvin(temp)
			}
		case "K":
			if to == "C" {
				result = conv.KelvinToCelsius(temp)
			} else if to == "F" {
				result = conv.KelvinToFahrenheit(temp)
			}
		}

		if result == float64(int(result)) {
			fmt.Printf("%g degrees %s is %g degrees %s\n", temp, from, result, to)
		} else {
			fmt.Printf("%.3g degrees %s is %g degrees %s\n", temp, from, result, to)
		}
	}
}

// Hjelpefunksjon som sjekker om en gitt streng finnes i en gitt liste
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
