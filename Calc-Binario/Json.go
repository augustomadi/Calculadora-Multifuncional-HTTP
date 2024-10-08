package main

type ErrorMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type OperationError struct {
	Result    string `json:"result"`
	Operation string `json:"op"`
}

type OperationResult struct {
	Result    float64 `json:"result"`
	Operation string  `json:"op"`
}

