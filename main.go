package main

import (
	first "adventofcode/2024-go/1"
	tenth "adventofcode/2024-go/10"
	eleventh "adventofcode/2024-go/11"
	twelfth "adventofcode/2024-go/12"
	thirteenth "adventofcode/2024-go/13"
	fourteenth "adventofcode/2024-go/14"
	fifteenth "adventofcode/2024-go/15"
	sixteenth "adventofcode/2024-go/16"
	second "adventofcode/2024-go/2"
	third "adventofcode/2024-go/3"
	fourth "adventofcode/2024-go/4"
	fifth "adventofcode/2024-go/5"
	sixth "adventofcode/2024-go/6"
	seventh "adventofcode/2024-go/7"
	eighth "adventofcode/2024-go/8"
	ninth "adventofcode/2024-go/9"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

func init() {
	log.SetOutput(io.Discard)
}

const profiling = false

func main() {
	if profiling {
		f, err := os.Create("cpu.prof")
		if err != nil {

			fmt.Println(err)
			return

		}
		pprof.StartCPUProfile(f)
		pprof.Profiles()
		defer pprof.StopCPUProfile()
	}

	first.Run()
	second.Run()
	third.Run()
	fourth.Run()
	fifth.Run()
	sixth.Run()
	seventh.Run()
	eighth.Run()
	ninth.Run()
	tenth.Run()
	eleventh.Run()
	twelfth.Run()
	thirteenth.Run()
	fourteenth.Run()
	fifteenth.Run()
	sixteenth.Run()
}
