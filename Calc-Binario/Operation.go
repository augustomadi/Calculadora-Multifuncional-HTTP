package main

import (
	"math"
)


type sum struct{}

func (s sum) calculate(x, y float64) float64 {
	return x + y
}

type sub struct{}

func (s sub) calculate(x, y float64) float64 {
	return x - y
}

type mul struct{}

func (m mul) calculate(x, y float64) float64 {
	return x * y
}

type div struct{}

func (d div) calculate(x, y float64) float64 {
	return x / y
}

type raiz struct{}

func (r raiz) calculate(x, y float64) float64 {
	return math.Pow(x, 1/y)
}

type pow struct{}

func (p pow) calculate(x, y float64) float64 {
	return math.Pow(x, y)
}

