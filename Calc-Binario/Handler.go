package main

import ("encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings")

type AndreiHandler struct{}

func (AndreiHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	log.Print("[INFO]" + req.Method + " " + req.URL.Path)

	if req.Method != "GET" || req.URL.Path != "/result" {
		e := ErrorMessage{Code: 404, Error: "Not Found"}
		r, _ := json.Marshal(e)
		resp.WriteHeader(404)
		resp.Write(r)
		return
	}

	operation := req.URL.Query().Get("op")

	if !isOperationValid(operation) {
		c := OperationError{Result: "invalid expression", Operation: operation}
		r, _ := json.Marshal(c)
		resp.WriteHeader(400)
		resp.Write(r)
		return
	}

	operator := getOperatorFromString(operation)
	numbers := strings.Split(operation, operator)

	mathe := operators[operator]
	x := numbers[0]
	y := numbers[1]
	
	number1,_ := binaryToDecimal(x)
	number2,_ := binaryToDecimal(y)
	
	numbertofloat := float64(number1)
	number2tofloat := float64(number2)

	result := mathe.calculate(numbertofloat, number2tofloat)

	c := OperationResult{Result: result, Operation: operation}
	r, _ := json.Marshal(c)
	resp.WriteHeader(200)
	resp.Write(r)

}

func isOperationValid(op string) bool {
	for k := range operators {
		if strings.Contains(op, k) {
			return true
		}
	}

	return false
}

func getOperatorFromString(op string) string {

	for k := range operators {
		if strings.Contains(op, k) {
			return k
		}
	}

	return ""
}


func binaryToDecimal(binary string) (int64, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

