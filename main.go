package main

import (
	"flag"
	"fmt"
	"github.com/Mic92/fcds-lab-2015/bucketsort"
	"github.com/Mic92/fcds-lab-2015/haar"
	"github.com/Mic92/fcds-lab-2015/threesat"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func doBucketsort(in, out *os.File) {
	if err := bucketsort.SortFile(in, out); err != nil {
		log.Fatal(err)
	}
}

func doThreesat(in, out *os.File) {
	solver, err := threesat.New(in)
	if err != nil {
		log.Fatal(err)
	}
	solution := solver.Solve()
	if solution == nil {
		out.WriteString("Solution not found.\n")
	} else {
		fmt.Fprintf(out, "Solution found [%d]: ", *solution)
		for i := uint(0); i < solver.NVar; i++ {
			if (*solution)&(1<<i) == 0 {
				out.WriteString("0 ")
			} else {
				out.WriteString("1 ")
			}
		}
		out.WriteString("\n")
	}
}

func doHaar(in, out *os.File) {
	if err := haar.ProcessFile(in, out); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetOutput(os.Stderr)

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
		log.Fatalf("USAGE: %s bucketsort|threesat|haar", os.Args[0])
	}

	var in, out *os.File
	var err error

	switch len(args) {
	case 1:
		in = os.Stdin
		out = os.Stdout
	case 2:
		in = os.Stdin
		out, err = os.Create(os.Args[2])
		if err != nil {
			log.Fatalf("error opening output file '%s': %v", args[1], err)
		}
	default:
		in, err = os.Open(os.Args[2])
		if err != nil {
			log.Fatalf("error opening input file '%s': %v", err)
		}
		out, err = os.Create(os.Args[3])
		if err != nil {
			log.Fatalf("error opening output file '%s': %v", args[1], err)
		}
	}

	start := time.Now()
	switch args[0] {
	case "bucketsort":
		doBucketsort(in, out)
	case "threesat":
		doThreesat(in, out)
	case "haar":
		doHaar(in, out)
	default:
		log.Fatalf("algorithm not implemented: %s", os.Args[1])
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
