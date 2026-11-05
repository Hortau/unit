package unit

import (
	"math"
	"testing"
)

func TestDegreeDaysUnits(t *testing.T) {
	tests := []struct {
		name       string
		days       DegreeDays
		celsius    float64
		fahrenheit float64
	}{
		{"one Celsius-day", CelsiusDay, 1, 1.8},
		{"one Fahrenheit-day", FahrenheitDay, 5.0 / 9.0, 1},
		{"negative Celsius-days", -2.5 * CelsiusDay, -2.5, -4.5},
		{"Fahrenheit-days", 18 * FahrenheitDay, 10, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloatEqual(t, tt.celsius, tt.days.CelsiusDays())
			assertFloatEqual(t, tt.fahrenheit, tt.days.FahrenheitDays())
		})
	}
}

func TestDegreeDaysConverter(t *testing.T) {
	tests := []struct {
		name  string
		from  string
		to    string
		value float64
		want  float64
	}{
		{"zero Celsius to Fahrenheit", "C-day", "F-day", 0, 0},
		{"one Celsius to Fahrenheit", "C-day", "F-day", 1, 1.8},
		{"positive Celsius to Fahrenheit", "C-day", "F-day", 10, 18},
		{"fractional Celsius to Fahrenheit", "C-day", "F-day", 0.5, 0.9},
		{"negative Celsius to Fahrenheit", "C-day", "F-day", -10, -18},
		{"large Celsius to Fahrenheit", "C-day", "F-day", math.MaxFloat64 / 2, (math.MaxFloat64 / 2) * (9.0 / 5.0)},
		{"zero Fahrenheit to Celsius", "F-day", "C-day", 0, 0},
		{"one Fahrenheit to Celsius", "F-day", "C-day", 1, 5.0 / 9.0},
		{"positive Fahrenheit to Celsius", "F-day", "C-day", 18, 10},
		{"fractional Fahrenheit to Celsius", "F-day", "C-day", 0.9, 0.5},
		{"negative Fahrenheit to Celsius", "F-day", "C-day", -18, -10},
		{"large Fahrenheit to Celsius", "F-day", "C-day", math.MaxFloat64, math.MaxFloat64 * (5.0 / 9.0)},
		{"same Celsius units", "C-day", "C-day", 12.5, 12.5},
		{"same Fahrenheit units", "F-day", "F-day", 25.5, 25.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConverter(tt.value).From(tt.from).To(tt.to)
			if err != nil {
				t.Fatalf("converting %s to %s: %v", tt.from, tt.to, err)
			}
			assertFloatEqual(t, tt.want, got)
		})
	}
}

func TestDegreeDaysConstructors(t *testing.T) {
	tests := []struct {
		name  string
		value Value
		to    string
		want  float64
	}{
		{"from Celsius-days", FromCelsiusDays(10), "F-day", 18},
		{"from Fahrenheit-days", FromFahrenheitDays(18), "C-day", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.value.To(tt.to)
			if err != nil {
				t.Fatal(err)
			}
			assertFloatEqual(t, tt.want, got)
		})
	}
}

func TestDegreeDaysRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		from  string
		to    string
		value float64
	}{
		{"Celsius-days", "C-day", "F-day", 23.45},
		{"Fahrenheit-days", "F-day", "C-day", -23.45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converted, err := NewConverter(tt.value).From(tt.from).To(tt.to)
			if err != nil {
				t.Fatal(err)
			}
			got, err := NewConverter(converted).From(tt.to).To(tt.from)
			if err != nil {
				t.Fatal(err)
			}
			assertFloatEqual(t, tt.value, got)
		})
	}
}

func TestDegreeDaysRejectIncompatibleUnits(t *testing.T) {
	tests := []struct {
		name  string
		value Value
		to    string
	}{
		{"Celsius-days to length", NewConverter(1).From("C-day"), "m"},
		{"Fahrenheit-days to length", NewConverter(1).From("F-day"), "m"},
		{"length to Celsius-days", NewConverter(1).From("m"), "C-day"},
		{"length to Fahrenheit-days", NewConverter(1).From("m"), "F-day"},
		{"temperature to Celsius-days", Value{273.15, temperature}, "C-day"},
		{"temperature to Fahrenheit-days", Value{273.15, temperature}, "F-day"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.value.To(tt.to)
			if err != ErrConversion {
				t.Fatalf("got error %v, want %v", err, ErrConversion)
			}
		})
	}
}
