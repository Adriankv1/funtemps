package conv

func FahrenheitToCelsius(value float64) float64 { // Fahrenheit to Celsius
	return (value - 32) * 5 / 9
}

func CelsiusToFahrenheit(value float64) float64 { // Celsius to Fahrenheit
	return (value * 9 / 5) + 32
}

func KelvinToFahrenheit(value float64) float64 { //Kelvin to Fahrenheit
	return (value * 9 / 5) - 459.67
}

func FahrenheitToKelvin(value float64) float64 { // Fahrenheit to Kelvin
	return (value + 459.67) * 5 / 9
}

func CelsiusToKelvin(value float64) float64 { // Celsius to Kelvin
	return value + 273.15
}

func KelvinToCelsius(value float64) float64 { // Kelvin to Celsius
	return value - 273.15
}
