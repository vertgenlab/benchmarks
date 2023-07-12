package main

import (
	"flag"
	"fmt"
	biogoBam "github.com/biogo/hts/bam"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fileio"
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

	if *cpuprof != "" {
		cpu, err := os.Create(*cpuprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(cpu)
		defer pprof.StopCPUProfile()
	}

	f, err := os.Open(*file)
	exception.PanicOnErr(err)
	br, err := biogoBam.NewReader(f, 1)
	exception.PanicOnErr(err)

	if *exectrace != "" {
		runtrace := fileio.EasyCreate(*exectrace)
		err = trace.Start(runtrace)
		exception.PanicOnErr(err)
	}

	var numReads int
	for _, err = br.Read(); err != io.EOF; _, err = br.Read() {
		exception.PanicOnErr(err)
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
