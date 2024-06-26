# About
[![codecov](https://codecov.io/gh/Hortau/unit/graph/badge.svg?token=7DEKCHwiZf)](https://codecov.io/gh/Hortau/unit)
[![GoDoc](https://godoc.org/github.com/hortau/unit?status.svg)](https://godoc.org/github.com/hortau/unit)

Conversion of unit library for golang


## Installation

```
go get -u github.com/hortau/unit
```


## General use

Basic usage:
```go
ft := 1 * unit.Foot
fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
```

To use your own data type, you need to convert to the base unit first (eg Length, Speed etc):
```go
type MyUnit int

n := MyUnit(2)
ft := Length(n) * Foot
fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
```


## Temperature

Cannot be used to scale directly like the other units.
Instead, use the From* functions to create a Temperature type:

```go
f := unit.FromFahrenheit(100)

fmt.Println("100 fahrenheit in celsius = ", f.Celsius())
```

Can also be converted directly from F to C or C to V for more precision

```go
f := unit.CelsiusToFahrenheit(22.3)
c := unit.FahrenheitToCelsius(80.0)
```


## Convert using unit abbreviation
```go
result, err := unit.NewConverter(24.4).From("m3").To("ft3")
```


## Notes

Please note the resulting precision is limited to the float64 type.


## License

Under [MIT](LICENSE)
