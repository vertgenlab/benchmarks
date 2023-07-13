package main

import (
	"flag"
	"fmt"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fileio"
	"github.com/vertgenlab/gonomics/sam"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

func main() {
	file := flag.String("i", "", "input file")
	cpuprof := flag.String("cpuprof", "", "cpu profile output")
	memprof := flag.String("memprof", "", "mem profile output")
	exectrace := flag.String("exectrace", "", "DEBUG: execution trace output for use with go tool trace")
	flag.Parse()

	var err error
	if *cpuprof != "" {
		cpu, err := os.Create(*cpuprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(cpu)
		defer pprof.StopCPUProfile()
	}

	if *exectrace != "" {
		runtrace := fileio.EasyCreate(*exectrace)
		err = trace.Start(runtrace)
		exception.PanicOnErr(err)
	}

	var numReads int
	var s sam.Sam
	br, _ := sam.OpenBam(*file)
	for {
		_, err = sam.DecodeBam(br, &s)
		if err == io.EOF {
			break
		}
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
