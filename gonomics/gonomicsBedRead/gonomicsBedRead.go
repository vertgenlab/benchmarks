package main

import (
	"flag"
	"fmt"
	"github.com/vertgenlab/gonomics/bed"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fileio"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func main() {
	file := flag.String("i", "", "input file")
	cpuprof := flag.String("cpuprof", "", "cpu profile output")
	memprof := flag.String("memprof", "", "mem profile output")
	exectrace := flag.String("exectrace", "", "DEBUG: execution trace output for use with go tool trace")
	flag.Parse()

	if *cpuprof != "" {
		cpu, err := os.Create(*cpuprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(cpu)
		defer pprof.StopCPUProfile()
	}

	var err error
	if *exectrace != "" {
		runtrace := fileio.EasyCreate(*exectrace)
		defer runtrace.Close()
		err = trace.Start(runtrace)
		exception.PanicOnErr(err)
	}

	var numReads int
	r := bed.GoReadToChan(*file)
	for _ = range r {
		time.Sleep(time.Millisecond)
		numReads++
	}
	fmt.Print(numReads)

	if *exectrace != "" {
		trace.Stop()
	}

	if *memprof != "" {
		mem, err := os.Create(*memprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(mem)
		mem.Close()
	}
}
