package main

var operators = map[string]mathematician{
	"+":   sum{},
	"sum": sum{},
	"-":   sub{},
	"sub": sub{},
	"*":   mul{},
	"mul": mul{},
	"&":   raiz{},
	"rot": raiz{},
	"^":   pow{},
	"pow": pow{},
	"/":   div{},
	"div": div{},
}




