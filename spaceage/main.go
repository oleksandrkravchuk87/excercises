package main

import "fmt"

type Planet string

//const (
//	Earth Planet = iota
//	Mercury
//	Venus
//	Mars
//	Jupiter
//	Saturn
//	Uranus
//	Neptune
//)

var spAge = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(sec float64, planet Planet) float64 {
	//return (sec / 60 / 60 / 24 / 365)
	return (sec / 60 / 60 / 24 / (365 * spAge[planet]))
}

func main() {
	fmt.Println(Age(2134835688, "Mercury"))
}

//Earth: orbital period 365.25 Earth days, or 31557600 seconds
//Mercury: orbital period 0.2408467 Earth years
//Venus: orbital period 0.61519726 Earth years
//Mars: orbital period 1.8808158 Earth years
//Jupiter: orbital period 11.862615 Earth years
//Saturn: orbital period 29.447498 Earth years
//Uranus: orbital period 84.016846 Earth years
//Neptune: orbital period 164.79132 Earth years
