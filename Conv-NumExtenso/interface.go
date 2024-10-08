package main 


type Operation struct{
	Number1 float64
	Number2 float64
}

type Mathematician interface{
	Calculate(numero1,numero2 int) int
}

type Add struct{
	Operation
}


func (m Add) Calculate(number1,number2 int)int{
	return number1 + number2
}

type Sub struct{
	Operation
}


func (m Sub) Calculate(number1,number2 int)int{
	return number1 - number2
}

type Mul struct{
	Operation
}


func (m Mul) Calculate(number1,number2 int)int{
	return number1 * number2
}

type Div struct{
	Operation
}


func (m Div) Calculate(number1,number2 int)int{	
	return number1 / number2
}

