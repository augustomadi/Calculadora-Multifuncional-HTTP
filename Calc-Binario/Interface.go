package main

type operation struct {
	numx     float64
	numy     float64
	operator string
}

type mathematician interface {
	calculate(x, y float64) float64
}

