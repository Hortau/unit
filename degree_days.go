package unit

// DegreeDays represents accumulated temperature differences in Celsius-days.
// One Celsius-day equals 9/5 Fahrenheit-days; unlike absolute temperatures,
// degree-day conversions do not apply a 32-degree offset. NewConverter
// recognizes the symbols "C-day" and "F-day".
type DegreeDays Unit

const (
	// CelsiusDay is one Celsius degree-day.
	CelsiusDay DegreeDays = 1
	// FahrenheitDay is one Fahrenheit degree-day.
	FahrenheitDay DegreeDays = CelsiusDay * 5 / 9
)

// FromCelsiusDays returns a value in Celsius-days for conversion.
func FromCelsiusDays(val float64) Value {
	return Value{val * float64(CelsiusDay), degreeDays}
}

func toCelsiusDays(value Value) (float64, error) {
	if value.unit != degreeDays {
		return 0, ErrConversion
	}
	return DegreeDays(value.val).CelsiusDays(), nil
}

// CelsiusDays returns the accumulated temperature difference in Celsius-days.
func (d DegreeDays) CelsiusDays() float64 {
	return float64(d)
}

// FromFahrenheitDays returns a value in Fahrenheit-days for conversion.
func FromFahrenheitDays(val float64) Value {
	return Value{val * float64(FahrenheitDay), degreeDays}
}

func toFahrenheitDays(value Value) (float64, error) {
	if value.unit != degreeDays {
		return 0, ErrConversion
	}
	return DegreeDays(value.val).FahrenheitDays(), nil
}

// FahrenheitDays returns the accumulated temperature difference in Fahrenheit-days.
func (d DegreeDays) FahrenheitDays() float64 {
	return float64(d) * (9.0 / 5.0)
}
