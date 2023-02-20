package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	tests := []struct {
		about    string
		expected []string
	}{
		{
			about: "sun",
			expected: []string{
				"Temperature in the Sun's core is 15 000 000 degrees Celsius",
				"Temperature of the outer layer of the Sun is 5778 Kelvin",
			},
		},
		{
			about: "moon",
			expected: []string{
				"Temperature of the Moon's surface at night is -183 degrees Celsius",
				"Temperature of the Moon's surface during the day is 106 degrees Celsius",
			},
		},
		{
			about: "earth",
			expected: []string{
				"Highest temperature measured on the Earth's surface is 134 degrees Fahrenheit, 56.7 degrees Celsius, 329.82 Kelvin",
				"Lowest temperature measured on the Earth's surface is -89.4 degrees Celsius",
				"Temperature in the Earth's inner core is 9392 Kelvin",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.about, func(t *testing.T) {
			actual := GetFunFacts(tt.about)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("GetFunFacts(%s): expected %v, actual %v", tt.about, tt.expected, actual)
			}
		})
	}
}
