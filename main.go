package main

import (
	"flag"
	"fmt"
	"github.com/Mic92/fcds-lab-2015/bucketsort"
	"github.com/Mic92/fcds-lab-2015/haar"
	"github.com/Mic92/fcds-lab-2015/threesat"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

const usage = "USAGE: %s bucketsort|threesat|haar\n"

func doBucketsort(in, out *os.File) time.Duration {
	duration, err := bucketsort.SortFile(in, out)
	if err != nil {
		log.Fatal(err)
	}
	return duration
}

func doThreesat(in, out *os.File) time.Duration {
	solver, err := threesat.New(in)
	if err != nil {
		log.Fatal(err)
	}
	duration, solution := solver.Solve()
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
	return duration
}

func doHaar(in, out *os.File) time.Duration {
	duration, err := haar.ProcessFile(in, out)
	if err != nil {
		log.Fatal(err)
	}
	return duration
}

func openInputOutput(args []string) (in, out *os.File, err error) {
	switch len(args) {
	case 0:
		in = os.Stdin
		out = os.Stdout
	case 1:
		in = os.Stdin
		out, err = os.OpenFile(args[0], os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0660)
		if err != nil {
			return nil, nil, fmt.Errorf("error opening output file '%s': %v", args[0], err)
		}
	default:
		in, err = os.Open(args[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error opening input file '%s': %v", args[0], err)
		}
		out, err = os.OpenFile(args[1], os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0660)
		if err != nil {
			return nil, nil, fmt.Errorf("error opening output file '%s': %v", args[1], err)
		}
	}
	return in, out, nil
}

func main() {
	log.SetOutput(os.Stderr)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}

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
		flag.Usage()
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	in, out, err := openInputOutput(args[1:])
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	defer out.Close()

	start := time.Now()
	var duration time.Duration
	switch args[0] {
	case "bucketsort":
		duration = doBucketsort(in, out)
	case "threesat":
		duration = doThreesat(in, out)
	case "haar":
		duration = doHaar(in, out)
	default:
		log.Fatalf("algorithm not implemented: %s", os.Args[1])
	}
	computation := float64(duration.Nanoseconds() / 1e6)
	overall := float64(time.Since(start).Nanoseconds() / 1e6)

	log.Printf("computation time: %fms, overall time: %fms", computation, overall)
}
