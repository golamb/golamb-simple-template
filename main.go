package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handle is lambda Handle function
func Handle(event json.RawMessage) (response interface{}, err error) {
	return string(event), nil
}

func main() {
	lambda.Start(Handle)
}
