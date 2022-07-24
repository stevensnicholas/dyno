package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
<<<<<<< HEAD
	println("This is where the lambda for the issues part of the project will be")
=======
	lambda.Start(Handler)
}

func Handler() {
	fmt.Println("Function invoked!")
>>>>>>> 8e0ffd5 (Added terraform and begun to create lambda function:)
}
