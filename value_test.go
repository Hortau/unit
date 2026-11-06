package unit

import "testing"

func TestConverterWithUpstreamUnits(t *testing.T) {
	tests := []struct {
		name  string
		from  string
		to    string
		value float64
		want  float64
	}{
		{"miles to meters", "mi", "m", 1, 1609.344},
		{"meters to miles", "m", "mi", 1609.344, 1},
		{"acres to hectares", "ac", "ha", 1, 0.40468564224},
		{"hectares to acres", "ha", "ac", 0.40468564224, 1},
		{"cubic feet to cubic meters", "ft3", "m3", 1, 0.028316846592},
		{"cubic meters to cubic feet", "m3", "ft3", 0.028316846592, 1},
		{"gallons to liters", "gal", "l", 1, 3.785411784},
		{"liters to gallons", "l", "gal", 3.785411784, 1},
		{"inches of mercury to kilopascals", "in-hg", "kPa", 1, 3.386389},
		{"kilopascals to inches of mercury", "kPa", "in-hg", 3.386389, 1},
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
