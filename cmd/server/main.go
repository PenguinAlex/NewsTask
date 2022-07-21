package main

import "awesomeProject/pkg/adder"

func main() {
	go adder.RunRest()
	adder.RunGrpc()
}
