package main

import (
	"flag"
	"fmt"
	elprepSam "github.com/exascience/elprep/sam"
	"github.com/vertgenlab/gonomics/exception"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	file := flag.String("i", "", "input file")
	cpuprof := flag.String("cpuprof", "", "cpu profile output")
	memprof := flag.String("memprof", "", "mem profile output")
	flag.Parse()

	if *cpuprof != "" {
		cpu, err := os.Create(*cpuprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(cpu)
		defer pprof.StopCPUProfile()
	}

	f, err := elprepSam.Open(*file, false)
	exception.PanicOnErr(err)
	defer f.Close()

	var numReads int
	f.SamReader().RunPipeline(elprepSam.NewSam(), []elprepSam.Filter{func(_ *elprepSam.Header) elprepSam.AlignmentFilter {
		return func(_ *elprepSam.Alignment) bool {
			numReads++
			return true
		}
	}}, elprepSam.Coordinate)
	fmt.Print(numReads)

	if *memprof != "" {
		mem, err := os.Create(*memprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(mem)
		mem.Close()
	}
}
