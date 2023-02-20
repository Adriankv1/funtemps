package main

import (
	"flag"
	"fmt"
	"funtemps/v2/conv"
)

var from string
var to string
var temp float64

func init() {
	flag.StringVar(&from, "from", "C", "source temperature scale (C/F/K)")
	flag.StringVar(&to, "to", "F", "target temperature scale (C/F/K)")
	flag.Float64Var(&temp, "temp", 0, "temperature to convert")
}

func main() {
	flag.Parse()

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

	if result == 0 {
		fmt.Println("Invalid input. Please check your source and target temperature scales.")
	} else {
		if result == float64(int(result)) {
			fmt.Printf("%.0g degrees %s is %g degrees %s\n", temp, from, result, to)
		} else {
			fmt.Printf("%.3g degrees %s is %g degrees %s\n", temp, from, result, to)
		}
	}
}
