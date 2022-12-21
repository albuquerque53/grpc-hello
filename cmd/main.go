package main

import (
	"flag"
	"grpchello/internal/application/handler"
	"log"
	"os"
)

func main() {
	var operation string

	flag.StringVar(&operation, "op", "serve", "The operation type, allowed values are: serve, call (the default is 'serve')")
	flag.Parse()

	if operation == "serve" {
		handler.RunServer()
		return
	}

	if operation == "call" {
		handler.ExecuteGRPCCall(os.Args)
		return
	}

	log.Fatalf("unknown operation %v, allowed values are: serve, call", operation)
}
