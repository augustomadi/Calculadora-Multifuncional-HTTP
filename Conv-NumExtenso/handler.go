package main

import ("encoding/json"
	"log"
	"net/http"
	"strings"
	
)

type IsaacHandler struct{}

func (IsaacHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

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
	x := StringToInt(numbers[0])
	y := StringToInt(numbers[1])
	
	
	result := mathe.Calculate(x, y)

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


