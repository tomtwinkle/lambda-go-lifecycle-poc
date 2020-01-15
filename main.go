package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

var invokeCount = 0
var initCount = 0
var mainCount = 0

func init() {
	log.Print("init start")
	initCount++
	log.Print("init end")
}

func Handler() (string, error) {
	log.Print("handler start")
	invokeCount++
	log.Printf("invoke=%d, main=%d, init=%d", invokeCount, mainCount, initCount)
	log.Print("handler end")
	return fmt.Sprintf("invoke=%d, main=%d, init=%d", invokeCount, mainCount, initCount), nil
}

func main() {
	log.Print("main start")
	mainCount++
	defer log.Print("main defer")
	lambda.Start(Handler)
	log.Print("main end")
}
