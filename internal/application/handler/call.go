package handler

import (
	"fmt"
	"grpchello/internal/application/client"
	"log"
)

func ExecuteGRPCCall(args []string) {
	if len(args[1:]) <= 1 {
		log.Fatal("Missing arguments. Usage: 'make call <name-here>")
	}

	name := args[2]

	res, err := client.CallHelloGRPC("localhost:9000", name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res)
}
