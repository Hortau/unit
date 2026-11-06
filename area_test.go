package unit

import (
	"testing"
)

func TestArea(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e6, (1 * SquareZeptometer).SquareYoctometers())
	assertFloatEqual(t, 1e6, (1 * SquareAttometer).SquareZeptometers())
	assertFloatEqual(t, 1e6, (1 * SquareFemtometer).SquareAttometers())
	assertFloatEqual(t, 1e6, (1 * SquarePicometer).SquareFemtometers())
	assertFloatEqual(t, 1e6, (1 * SquareNanometer).SquarePicometers())
	assertFloatEqual(t, 1e6, (1 * SquareMicrometer).SquareNanometers())
	assertFloatEqual(t, 1e6, (1 * SquareMillimeter).SquareMicrometers())
	assertFloatEqual(t, 1e6, (1 * SquareMeter).SquareMillimeters())

	assertFloatEqual(t, 1e4, (1 * SquareMeter).SquareCentimeters())
	assertFloatEqual(t, 1e2, (1 * SquareMeter).SquareDecimeters())
	assertFloatEqual(t, 1e0, (1 * SquareMeter).SquareMeters())
	assertFloatEqual(t, 1e-2, (1 * SquareMeter).SquareDecameter())
	assertFloatEqual(t, 1e-4, (1 * SquareMeter).SquareHectometer())

	assertFloatEqual(t, 1e-6, (1 * SquareMeter).SquareKilometers())
	assertFloatEqual(t, 1e-6, (1 * SquareKilometer).SquareMegameters())
	assertFloatEqual(t, 1e-6, (1 * SquareMegameter).SquareGigameters())
	assertFloatEqual(t, 1e-6, (1 * SquareGigameter).SquareTerameters())
	assertFloatEqual(t, 1e-6, (1 * SquareTerameter).SquarePetameters())
	assertFloatEqual(t, 1e-6, (1 * SquarePetameter).SquareExameters())
	assertFloatEqual(t, 1e-6, (1 * SquareExameter).SquareZettameters())
	assertFloatEqual(t, 1e-6, (1 * SquareZettameter).SquareYottameters())

	// US
	assertFloatEqual(t, 0.00064516, (1 * SquareInch).SquareMeters())
	assertFloatEqual(t, 10.763910416709724, (1 * SquareMeter).SquareFeet())
	assertFloatEqual(t, 144.0, (1 * SquareFoot).SquareInches())
	assertFloatEqual(t, 9.0, (1 * SquareYard).SquareFeet())
	assertFloatEqual(t, 4840.0, (1 * Acre).SquareYards())
	assertFloatEqual(t, 640.0, (1 * SquareMile).Acres())
	assertFloatEqual(t, 0.0015625, (1 * Acre).SquareMiles())

	// imperial
	assertFloatEqual(t, 25.29285264, (1 * SquareRod).SquareMeters())
	assertFloatEqual(t, 1011.7141055999998, (1 * Rood).SquareMeters())
	assertFloatEqual(t, 0.025, (1 * SquareRod).Roods())

	// aliases
	assertFloatEqual(t, 1.0, (1 * SquareMeter).Centiares())
	assertFloatEqual(t, 1.0, (1 * SquareDecameter).Ares())
	assertFloatEqual(t, 1.0, (1 * SquareHectometer).Hectares())
	assertFloatEqual(t, 1.0, (1 * SquarePerch).SquareRods())
}

func TestTaiwaneseAreas(t *testing.T) {
	tests := []struct {
		name           string
		unit           Area
		squareMeters   float64
		perSquareMeter float64
	}{
		{"gong qing", TWGongQing, 10000, 1.0 / 10000.0},
		{"ping", TWPing, 400.0 / 121.0, 121.0 / 400.0},
		{"fen", TWFen, 117360.0 / 121.0, 121.0 / 117360.0},
		{"jia", TWJia, 1173600.0 / 121.0, 121.0 / 1173600.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloatEqual(t, tt.squareMeters, tt.unit.SquareMeters())
			assertFloatEqual(t, tt.perSquareMeter, float64(SquareMeter/tt.unit))
			assertFloatEqual(t, 2.5*tt.squareMeters, (2.5 * tt.unit).SquareMeters())
		})
	}

	assertFloatEqual(t, 293.4, float64(TWFen/TWPing))
	assertFloatEqual(t, 10, float64(TWJia/TWFen))
	assertFloatEqual(t, 1173600.0/121.0/10000.0, (1 * TWJia).Hectares())
}
