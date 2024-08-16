package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	ipv4Counter "github.com/fr13n8/ipv4-counter/counter"
)

func main() {
	cpuprofile := flag.String("cf", "", "write cpu profile to `file`")
	memprofile := flag.String("mf", "", "write memory profile to `file`")
	tracefile := flag.String("tf", "", "write trace execution to `file`")
	input := flag.String("input", "", "path to the input file with ipv4 addresses")

	flag.Parse()

	if *input == "" {
		log.Fatal("you need to specify input file path")
	}

	if *tracefile != "" {
		f, err := os.Create("./profiles/" + *tracefile)
		if err != nil {
			log.Fatal("error when creating trace execution profile: ", err)
		}
		defer f.Close()
		trace.Start(f)
		defer trace.Stop()
	}

	if *cpuprofile != "" {
		f, err := os.Create("./profiles/" + *cpuprofile)
		if err != nil {
			log.Fatal("error when creating CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	count, err := ipv4Counter.IPV4CountFromFile(*input, runtime.NumCPU(), 1024)
	if err != nil {
		log.Fatalln("failed to count from file: ", err)
	}
	fmt.Printf("Count: %d\n", count)

	if *memprofile != "" {
		f, err := os.Create("./profiles/" + *memprofile)
		if err != nil {
			log.Fatal("error when creating memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("error when writing memory profile: ", err)
		}
	}
}
