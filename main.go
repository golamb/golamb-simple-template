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
	/* Run in locally my way */
	// res, err := Handle([]byte("Hello"))
	// log.Println(res, err)
	lambda.Start(Handle)
}
