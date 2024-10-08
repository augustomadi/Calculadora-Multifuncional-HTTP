package main

var operators = map[string]Mathematician{
	"+":   Add{},
	"sum": Add{},
	"-":   Sub{},
	"sub": Sub{},
	"*":   Mul{},
	"mul": Mul{},
	"/":   Div{},
	"div": Div{},
}


