package benchmarking_test

import (
	"github.com/vertgenlab/gonomics/numbers"
	"log"
	"os/exec"
	"testing"
	"syscall"
	"fmt"
)

// for benchmarking memory usage
var memUsage []float64
var cmd *exec.Cmd

// benchmarking calls

// 1 set of all tools
func BenchmarkGonomicsIntervalOverlapMem100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlapMem("testdata/test100a.bed", "testdata/test100b.bed")
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlapMem100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsSingleThreadIntervalOverlapMem("testdata/test100a.bed", "testdata/test100b.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlapMem100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlapMem("testdata/test100a.bed", "testdata/test100b.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortAMem100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortAMem("testdata/test100a.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortBMem100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortBMem("testdata/test100b.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlapMem100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlapMem("testdata/test100a.sorted.bed", "testdata/test100b.sorted.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of all tools
func BenchmarkGonomicsIntervalOverlapMem100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlapMem("testdata/test100a.bed", "testdata/test10mib.bed")
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlapMem100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsSingleThreadIntervalOverlapMem("testdata/test100a.bed", "testdata/test10mib.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlapMem100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlapMem("testdata/test100a.bed", "testdata/test10mib.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortAMem100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortAMem("testdata/test100a.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortBMem100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortBMem("testdata/test10mib.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlapMem100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlapMem("testdata/test100a.sorted.bed", "testdata/test10mib.sorted.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of all tools
func BenchmarkGonomicsIntervalOverlapMem10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlapMem("testdata/test10mia.bed", "testdata/test100b.bed")
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlapMem10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsSingleThreadIntervalOverlapMem("testdata/test10mia.bed", "testdata/test100b.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlapMem10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlapMem("testdata/test10mia.bed", "testdata/test100b.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortAMem10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortAMem("testdata/test10mia.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortBMem10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortBMem("testdata/test100b.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlapMem10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlapMem("testdata/test10mia.sorted.bed", "testdata/test100b.sorted.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of all tools
func BenchmarkGonomicsIntervalOverlapMem10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlapMem("testdata/test10mia.bed", "testdata/test10mib.bed")
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlapMem10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsSingleThreadIntervalOverlapMem("testdata/test10mia.bed", "testdata/test10mib.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlapMem10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlapMem("testdata/test10mia.bed", "testdata/test10mib.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortAMem10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortAMem("testdata/test10mia.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsSortBMem10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsSortBMem("testdata/test10mib.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlapMem10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlapMem("testdata/test10mia.sorted.bed", "testdata/test10mib.sorted.bed")
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// function calls

func gonomicsIntervalOverlapMem(fileA string, fileB string) *exec.Cmd{
	cmd := exec.Command("programs/intervalOverlap", "-threads", "8", fileA, fileB, "/dev/null")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func gonomicsSingleThreadIntervalOverlapMem(fileA string, fileB string) *exec.Cmd{
	cmd := exec.Command("programs/intervalOverlap", fileA, fileB, "/dev/null")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func bedtoolsIntervalOverlapMem(fileA string, fileB string) *exec.Cmd{
	cmd := exec.Command("bedtools", "intersect", "-u", "-wa", "-a", fileB, "-b", fileA)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func bedopsIntervalOverlapMem(fileA string, fileB string) *exec.Cmd{
	cmd := exec.Command("programs/bin/bedops", "--element-of", "1", fileB, fileA)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func bedopsSortAMem(fileA string) *exec.Cmd{
	cmd := exec.Command("programs/bin/sort-bed", fileA)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func bedopsSortBMem(fileB string) *exec.Cmd{
	cmd := exec.Command("programs/bin/sort-bed", fileB)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}
