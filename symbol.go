package unit

var fromMap = map[string]func(val float64) Value{
	"A":       FromAmperes,
	"ac":      FromAcres,
	"ac-ft":   FromAcreFeet,
	"ac-in":   FromAcreInches,
	"cm":      FromCentimeters,
	"deg":     FromDegrees,
	"ft3":     FromCubicFeet,
	"ft3/s":   FromCubicFeetPerSecond,
	"gal":     FromUSLiquidGallons,
	"gal/h":   FromUSLiquidGallonsPerHour,
	"gal/min": FromUSLiquidGallonsPerMinute,
	"ha":      FromHectares,
	"Hz":      FromHertz,
	"in":      FromInches,
	"in-hg":   FromInchesOfMercury,
	"kA":      FromKiloamperes,
	"Kb":      FromKilobits,
	"kb":      FromKilobits,
	"kHz":     FromKilohertz,
	"km":      FromKilometers,
	"kPa":     FromKilopascals,
	"kV":      FromKilovolts,
	"kW":      FromKilowatts,
	"l":       FromLiters,
	"l/h":     FromLitersPerHour,
	"l/min":   FromLitersPerMinute,
	"m3":      FromCubicMeters,
	"m3/s":    FromCubicMetersPerSecond,
	"Mb":      FromMegabits,
	"mb":      FromMegabits,
	"mA":      FromMilliamperes,
	"mi":      FromMiles,
	"mm":      FromMillimeters,
	"mV":      FromMillivolts,
	"psi":     FromPoundsPerSquareInch,
	"rpm":     FromRevolutionsPerMinute,
	"V":       FromVolts,
	"W":       FromWatts,
}

var toMap = map[string]func(value Value) (float64, error){
	"A":       toAmperes,
	"ac":      toAcres,
	"ac-ft":   toAcreFeet,
	"ac-in":   toAcreInches,
	"cm":      toCentimeters,
	"deg":     toDegrees,
	"ft3":     toCubicFeet,
	"ft3/s":   toCubicFeetPerSecond,
	"gal":     toUSLiquidGallons,
	"gal/h":   toUSLiquidGallonsPerHour,
	"gal/min": toUSLiquidGallonsPerMinute,
	"ha":      toHectares,
	"Hz":      toHertz,
	"in":      toInches,
	"in-hg":   toInchesOfMercury,
	"kA":      toKiloamperes,
	"kb":      toKilobits,
	"Kb":      toKilobits,
	"kHz":     toKilohertz,
	"km":      toKilometers,
	"kPa":     toKilopascals,
	"kV":      toKilovolts,
	"kW":      toKilowatts,
	"l":       toLiters,
	"l/h":     toLitersPerHour,
	"l/min":   toLitersPerMinute,
	"m3":      toCubicMeters,
	"m3/s":    toCubicMetersPerSecond,
	"Mb":      toMegabits,
	"mb":      toMegabits,
	"mA":      toMilliamperes,
	"mi":      toMiles,
	"mm":      toMillimeters,
	"mV":      toMillivolts,
	"psi":     toPoundsPerSquareInch,
	"rpm":     toRevolutionsPerMinute,
	"V":       toVolts,
	"W":       toWatts,
}