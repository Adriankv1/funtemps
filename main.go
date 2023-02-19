package main

import (
	"flag"
	"fmt"

	"funtemps/v2/conv"
)

var (
	f   float64
	c   float64
	k   float64
	out string
	t   string
)

func init() {
	flag.Float64Var(&f, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&c, "C", 0.0, "temperature in degrees Celsius")
	flag.Float64Var(&k, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&out, "out", "C", "output temperature scale: C - Celsius, F - Fahrenheit, K - Kelvin")
	flag.StringVar(&t, "t", "C", "temperature scale for fun facts: C - Celsius, F - Fahrenheit, K - Kelvin")
}

func main() {
	flag.Parse()

	numTempFlags := 0
	if f != 0 {
		numTempFlags++
	}
	if c != 0 {
		numTempFlags++
	}
	if k != 0 {
		numTempFlags++
	}
	if numTempFlags != 1 {
		fmt.Println("Please specify exactly one of the temperature flags (-F, -C, -K)")
		return
	}

	var converted float64
	switch {
	case f != 0:
		switch out {
		case "C":
			converted = conv.FahrenheitToCelsius(f)
		case "K":
			converted = conv.FahrenheitToKelvin(f)
		case "F":
			converted = f
		}
	case c != 0:
		switch out {
		case "F":
			converted = conv.CelsiusToFahrenheit(c)
		case "K":
			converted = conv.CelsiusToKelvin(c)
		case "C":
			converted = c
		}
	case k != 0:
		switch out {
		case "F":
			converted = conv.KelvinToFahrenheit(k)
		case "C":
			converted = conv.KelvinToCelsius(k)
		case "K":
			converted = k
		}
	}

	switch out {
	case "C":
		fmt.Printf("Output: %.2fK er %.2f°C\n", converted, c)
	case "F":
		fmt.Printf("Output: %.2fK er %.2f°F\n", converted, f)
	case "K":
		fmt.Printf("Output: %.2fK er %.2fK\n", converted, k)
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
