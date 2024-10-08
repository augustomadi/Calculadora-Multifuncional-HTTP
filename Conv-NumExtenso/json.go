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
	Result    int `json:"result"`
	Operation string  `json:"op"`
}

