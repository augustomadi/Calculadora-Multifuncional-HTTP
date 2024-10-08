package main 

type Operation struct{
	Valor_inicial float64
}

type MoneyChanger interface{
	Calculate(value float64) float64
}

type RealDolar struct{
	Operation
}

func (r RealDolar) Calculate(value float64) float64{
	return value * 5
}

type DolarReal struct{
	Operation
}

func (r DolarReal) Calculate(value float64) float64{
	return value / 5
}