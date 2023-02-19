package conv

import (
	"reflect"
	"testing"
)

func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 0},
		{input: 212, want: 100},
		{input: -40, want: -40},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("FahrenheitToCelsius(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 32},
		{input: 100, want: 212},
		{input: -40, want: -40},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("CelsiusToFahrenheit(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: -459.67},
		{input: 273.15, want: 32},
		{input: 373.15, want: 212},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("KelvinToFahrenheit(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
		{input: 212, want: 373.15},
		{input: -40, want: 233.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("FahrenheitToKelvin(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
		{input: 100, want: 373.15},
		{input: -40, want: 233.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("CelsiusToKelvin(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: -273.15},
		{input: 273.15, want: 0},
		{input: 373.15, want: 100},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("KelvinToCelsius(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}
