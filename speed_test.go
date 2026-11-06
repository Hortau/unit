package unit

import (
	"testing"
)

func TestSpeed(t *testing.T) {

	// SI
	assertFloatEqual(t, 10.0, (10 * MetersPerSecond).MetersPerSecond())
	assertFloatEqual(t, 36.0, (10 * MetersPerSecond).KilometersPerHour())
	assertFloatEqual(t, -9.0, (-2.5 * MetersPerSecond).KilometersPerHour())
	assertFloatEqual(t, 1000.0/3600.0, (1 * KilometersPerHour).MetersPerSecond())
	assertFloatEqual(t, 1000.0, (3600 * KilometersPerHour).MetersPerSecond())

	// US
	assertFloatEqual(t, 32.808398950131235, (10 * MetersPerSecond).FeetPerSecond())
	assertFloatEqual(t, 22.369362920544024, (10 * MetersPerSecond).MilesPerHour())

	// misc
	assertFloatEqual(t, 36000.0/1852.0, (10 * MetersPerSecond).Knots())
	assertFloatEqual(t, 1852.0/3600.0, (1 * Knot).MetersPerSecond())
	assertFloatEqual(t, 1852.0, (3600 * Knot).MetersPerSecond())

	// space
	assertFloatEqual(t, 0.03335640951981521, (10000000 * MetersPerSecond).SpeedOfLight())
}
