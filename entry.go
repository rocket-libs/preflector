package main

import (
	"log"
	"rocketlibs/preflector/delegate"
)

func main() {
	err := delegate.Work("")
	if err != nil {
		log.Fatal("Error occured")
		log.Fatal(err)
	}
}
