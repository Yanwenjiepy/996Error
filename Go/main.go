package main

import (
	"error_996/error_996"
	"errors"
	"log"
)

func main() {
	err := errors.New("996 work is evil")
	e := error_996.New("[I hate 996]")

	descStr := "test err: %w"
	err = e.Wrap(descStr, err)

	log.Println("wrapped err :", err)
	// wrapped err : [I hate 996]test err: 996 work is evil

	panic(err)
	/*
		panic: [I hate 996]test err: 996 work is evil

		goroutine 1 [running]:
		main.main()
		        /home/projects/996Error/Go/main.go:19 +0x18f

	*/
}
