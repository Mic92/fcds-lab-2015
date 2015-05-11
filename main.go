package main

import (
	"flag"
	"fmt"
	"github.com/Mic92/fcds-lab-2015/bucketsort"
	"github.com/Mic92/fcds-lab-2015/threesat"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func doBucketsort() {
	in, err := os.Open("bucketsort/input/medium.in")
	if err != nil {
		log.Fatal(err)
	}

	if err := bucketsort.SortFile(in, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func doThreesat() {
	solver, err := threesat.New(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	solution := solver.Solve()
	if solution == nil {
		fmt.Println("Solution not found.")
	} else {
		fmt.Printf("Solution found [%d]: ", *solution)
		for i := uint(0); i < solver.NVar; i++ {
			if (*solution)&(1<<i) == 0 {
				fmt.Print("0 ")
			} else {
				fmt.Print("1 ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("USAGE: %s bucketsort|threesat", os.Args[0])
	}

	start := time.Now()
	switch args[0] {
	case "bucketsort":
		doBucketsort()
	case "threesat":
		doThreesat()
	default:
		log.Fatalf("algorithm not implemented: %s", os.Args[1])
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
