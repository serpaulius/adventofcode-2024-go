package main

import (
	first "adventofcode/2024-go/1"
	second "adventofcode/2024-go/2"
	third "adventofcode/2024-go/3"
	fourth "adventofcode/2024-go/4"
	fifth "adventofcode/2024-go/5"
	sixth "adventofcode/2024-go/6"
	seventh "adventofcode/2024-go/7"
	"io"
	"log"
)

func init() {
	log.SetOutput(io.Discard)
}

func main() {
	first.Run()
	second.Run()
	third.Run()
	fourth.Run()
	fifth.Run()
	sixth.Run()
	seventh.Run()
}
