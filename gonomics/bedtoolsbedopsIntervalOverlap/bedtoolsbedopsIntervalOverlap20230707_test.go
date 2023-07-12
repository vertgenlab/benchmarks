package benchmarking_test

import (
	"github.com/vertgenlab/gonomics/numbers"
	"log"
	"os/exec"
	"testing"
	"syscall"
	"fmt"
)

var (
	// gonomics
	// a100
	gonomics100_100 = "csh/gonomics/a100/intervalOverlapGonomics100_100.csh"
	gonomics100_1k = "csh/gonomics/a100/intervalOverlapGonomics100_1k.csh"
	gonomics100_10k = "csh/gonomics/a100/intervalOverlapGonomics100_10k.csh"
	gonomics100_100k = "csh/gonomics/a100/intervalOverlapGonomics100_100k.csh"
	gonomics100_1mi = "csh/gonomics/a100/intervalOverlapGonomics100_1mi.csh"
	gonomics100_10mi = "csh/gonomics/a100/intervalOverlapGonomics100_10mi.csh"
	// a1k
	gonomics1k_100 = "csh/gonomics/a1k/intervalOverlapGonomics1k_100.csh"
	gonomics1k_1k = "csh/gonomics/a1k/intervalOverlapGonomics1k_1k.csh"
	gonomics1k_10k = "csh/gonomics/a1k/intervalOverlapGonomics1k_10k.csh"
	gonomics1k_100k = "csh/gonomics/a1k/intervalOverlapGonomics1k_100k.csh"
	gonomics1k_1mi = "csh/gonomics/a1k/intervalOverlapGonomics1k_1mi.csh"
	gonomics1k_10mi = "csh/gonomics/a1k/intervalOverlapGonomics1k_10mi.csh"
	// a10k
	gonomics10k_100 = "csh/gonomics/a10k/intervalOverlapGonomics10k_100.csh"
	gonomics10k_1k = "csh/gonomics/a10k/intervalOverlapGonomics10k_1k.csh"
	gonomics10k_10k = "csh/gonomics/a10k/intervalOverlapGonomics10k_10k.csh"
	gonomics10k_100k = "csh/gonomics/a10k/intervalOverlapGonomics10k_100k.csh"
	gonomics10k_1mi = "csh/gonomics/a10k/intervalOverlapGonomics10k_1mi.csh"
	gonomics10k_10mi = "csh/gonomics/a10k/intervalOverlapGonomics10k_10mi.csh"
	// a100k
	gonomics100k_100 = "csh/gonomics/a100k/intervalOverlapGonomics100k_100.csh"
	gonomics100k_1k = "csh/gonomics/a100k/intervalOverlapGonomics100k_1k.csh"
	gonomics100k_10k = "csh/gonomics/a100k/intervalOverlapGonomics100k_10k.csh"
	gonomics100k_100k = "csh/gonomics/a100k/intervalOverlapGonomics100k_100k.csh"
	gonomics100k_1mi = "csh/gonomics/a100k/intervalOverlapGonomics100k_1mi.csh"
	gonomics100k_10mi = "csh/gonomics/a100k/intervalOverlapGonomics100k_10mi.csh"
	// a1mi
	gonomics1mi_100 = "csh/gonomics/a1mi/intervalOverlapGonomics1mi_100.csh"
	gonomics1mi_1k = "csh/gonomics/a1mi/intervalOverlapGonomics1mi_1k.csh"
	gonomics1mi_10k = "csh/gonomics/a1mi/intervalOverlapGonomics1mi_10k.csh"
	gonomics1mi_100k = "csh/gonomics/a1mi/intervalOverlapGonomics1mi_100k.csh"
	gonomics1mi_1mi = "csh/gonomics/a1mi/intervalOverlapGonomics1mi_1mi.csh"
	gonomics1mi_10mi = "csh/gonomics/a1mi/intervalOverlapGonomics1mi_10mi.csh"
	// a10mi
	gonomics10mi_100 = "csh/gonomics/a10mi/intervalOverlapGonomics10mi_100.csh"
	gonomics10mi_1k = "csh/gonomics/a10mi/intervalOverlapGonomics10mi_1k.csh"
	gonomics10mi_10k = "csh/gonomics/a10mi/intervalOverlapGonomics10mi_10k.csh"
	gonomics10mi_100k = "csh/gonomics/a10mi/intervalOverlapGonomics10mi_100k.csh"
	gonomics10mi_1mi = "csh/gonomics/a10mi/intervalOverlapGonomics10mi_1mi.csh"
	gonomics10mi_10mi = "csh/gonomics/a10mi/intervalOverlapGonomics10mi_10mi.csh"

	// bedtools
	// a100
	bedtools100_100 = "csh/bedtools/a100/intervalOverlapBedtools100_100.csh"
	bedtools100_1k = "csh/bedtools/a100/intervalOverlapBedtools100_1k.csh"
	bedtools100_10k = "csh/bedtools/a100/intervalOverlapBedtools100_10k.csh"
	bedtools100_100k = "csh/bedtools/a100/intervalOverlapBedtools100_100k.csh"
	bedtools100_1mi = "csh/bedtools/a100/intervalOverlapBedtools100_1mi.csh"
	bedtools100_10mi = "csh/bedtools/a100/intervalOverlapBedtools100_10mi.csh"
	// a1k
	bedtools1k_100 = "csh/bedtools/a1k/intervalOverlapBedtools1k_100.csh"
	bedtools1k_1k = "csh/bedtools/a1k/intervalOverlapBedtools1k_1k.csh"
	bedtools1k_10k = "csh/bedtools/a1k/intervalOverlapBedtools1k_10k.csh"
	bedtools1k_100k = "csh/bedtools/a1k/intervalOverlapBedtools1k_100k.csh"
	bedtools1k_1mi = "csh/bedtools/a1k/intervalOverlapBedtools1k_1mi.csh"
	bedtools1k_10mi = "csh/bedtools/a1k/intervalOverlapBedtools1k_10mi.csh"
	// a10k
	bedtools10k_100 = "csh/bedtools/a10k/intervalOverlapBedtools10k_100.csh"
	bedtools10k_1k = "csh/bedtools/a10k/intervalOverlapBedtools10k_1k.csh"
	bedtools10k_10k = "csh/bedtools/a10k/intervalOverlapBedtools10k_10k.csh"
	bedtools10k_100k = "csh/bedtools/a10k/intervalOverlapBedtools10k_100k.csh"
	bedtools10k_1mi = "csh/bedtools/a10k/intervalOverlapBedtools10k_1mi.csh"
	bedtools10k_10mi = "csh/bedtools/a10k/intervalOverlapBedtools10k_10mi.csh"
	// a100k
	bedtools100k_100 = "csh/bedtools/a100k/intervalOverlapBedtools100k_100.csh"
	bedtools100k_1k = "csh/bedtools/a100k/intervalOverlapBedtools100k_1k.csh"
	bedtools100k_10k = "csh/bedtools/a100k/intervalOverlapBedtools100k_10k.csh"
	bedtools100k_100k = "csh/bedtools/a100k/intervalOverlapBedtools100k_100k.csh"
	bedtools100k_1mi = "csh/bedtools/a100k/intervalOverlapBedtools100k_1mi.csh"
	bedtools100k_10mi = "csh/bedtools/a100k/intervalOverlapBedtools100k_10mi.csh"
	// a1mi
	bedtools1mi_100 = "csh/bedtools/a1mi/intervalOverlapBedtools1mi_100.csh"
	bedtools1mi_1k = "csh/bedtools/a1mi/intervalOverlapBedtools1mi_1k.csh"
	bedtools1mi_10k = "csh/bedtools/a1mi/intervalOverlapBedtools1mi_10k.csh"
	bedtools1mi_100k = "csh/bedtools/a1mi/intervalOverlapBedtools1mi_100k.csh"
	bedtools1mi_1mi = "csh/bedtools/a1mi/intervalOverlapBedtools1mi_1mi.csh"
	bedtools1mi_10mi = "csh/bedtools/a1mi/intervalOverlapBedtools1mi_10mi.csh"
	// a10mi
	bedtools10mi_100 = "csh/bedtools/a10mi/intervalOverlapBedtools10mi_100.csh"
	bedtools10mi_1k = "csh/bedtools/a10mi/intervalOverlapBedtools10mi_1k.csh"
	bedtools10mi_10k = "csh/bedtools/a10mi/intervalOverlapBedtools10mi_10k.csh"
	bedtools10mi_100k = "csh/bedtools/a10mi/intervalOverlapBedtools10mi_100k.csh"
	bedtools10mi_1mi = "csh/bedtools/a10mi/intervalOverlapBedtools10mi_1mi.csh"
	bedtools10mi_10mi = "csh/bedtools/a10mi/intervalOverlapBedtools10mi_10mi.csh"

	// bedops
	// a100
	bedops100_100 = "csh/bedops/a100/intervalOverlapBedops100_100.csh"
	bedops100_1k = "csh/bedops/a100/intervalOverlapBedops100_1k.csh"
	bedops100_10k = "csh/bedops/a100/intervalOverlapBedops100_10k.csh"
	bedops100_100k = "csh/bedops/a100/intervalOverlapBedops100_100k.csh"
	bedops100_1mi = "csh/bedops/a100/intervalOverlapBedops100_1mi.csh"
	bedops100_10mi = "csh/bedops/a100/intervalOverlapBedops100_10mi.csh"
	// a1k
	bedops1k_100 = "csh/bedops/a1k/intervalOverlapBedops1k_100.csh"
	bedops1k_1k = "csh/bedops/a1k/intervalOverlapBedops1k_1k.csh"
	bedops1k_10k = "csh/bedops/a1k/intervalOverlapBedops1k_10k.csh"
	bedops1k_100k = "csh/bedops/a1k/intervalOverlapBedops1k_100k.csh"
	bedops1k_1mi = "csh/bedops/a1k/intervalOverlapBedops1k_1mi.csh"
	bedops1k_10mi = "csh/bedops/a1k/intervalOverlapBedops1k_10mi.csh"
	// a10k
	bedops10k_100 = "csh/bedops/a10k/intervalOverlapBedops10k_100.csh"
	bedops10k_1k = "csh/bedops/a10k/intervalOverlapBedops10k_1k.csh"
	bedops10k_10k = "csh/bedops/a10k/intervalOverlapBedops10k_10k.csh"
	bedops10k_100k = "csh/bedops/a10k/intervalOverlapBedops10k_100k.csh"
	bedops10k_1mi = "csh/bedops/a10k/intervalOverlapBedops10k_1mi.csh"
	bedops10k_10mi = "csh/bedops/a10k/intervalOverlapBedops10k_10mi.csh"
	// a100k
	bedops100k_100 = "csh/bedops/a100k/intervalOverlapBedops100k_100.csh"
	bedops100k_1k = "csh/bedops/a100k/intervalOverlapBedops100k_1k.csh"
	bedops100k_10k = "csh/bedops/a100k/intervalOverlapBedops100k_10k.csh"
	bedops100k_100k = "csh/bedops/a100k/intervalOverlapBedops100k_100k.csh"
	bedops100k_1mi = "csh/bedops/a100k/intervalOverlapBedops100k_1mi.csh"
	bedops100k_10mi = "csh/bedops/a100k/intervalOverlapBedops100k_10mi.csh"
	// a1mi
	bedops1mi_100 = "csh/bedops/a1mi/intervalOverlapBedops1mi_100.csh"
	bedops1mi_1k = "csh/bedops/a1mi/intervalOverlapBedops1mi_1k.csh"
	bedops1mi_10k = "csh/bedops/a1mi/intervalOverlapBedops1mi_10k.csh"
	bedops1mi_100k = "csh/bedops/a1mi/intervalOverlapBedops1mi_100k.csh"
	bedops1mi_1mi = "csh/bedops/a1mi/intervalOverlapBedops1mi_1mi.csh"
	bedops1mi_10mi = "csh/bedops/a1mi/intervalOverlapBedops1mi_10mi.csh"
	// a10mi
	bedops10mi_100 = "csh/bedops/a10mi/intervalOverlapBedops10mi_100.csh"
	bedops10mi_1k = "csh/bedops/a10mi/intervalOverlapBedops10mi_1k.csh"
	bedops10mi_10k = "csh/bedops/a10mi/intervalOverlapBedops10mi_10k.csh"
	bedops10mi_100k = "csh/bedops/a10mi/intervalOverlapBedops10mi_100k.csh"
	bedops10mi_1mi = "csh/bedops/a10mi/intervalOverlapBedops10mi_1mi.csh"
	bedops10mi_10mi = "csh/bedops/a10mi/intervalOverlapBedops10mi_10mi.csh"

	// gonomicsSingleThread
	// a100
	gonomicsSingleThread100_100 = "csh/gonomicsSingleThread/a100/intervalOverlapSingleThread100_100.csh"
	gonomicsSingleThread100_1k = "csh/gonomicsSingleThread/a100/intervalOverlapSingleThread100_1k.csh"
	gonomicsSingleThread100_10k = "csh/gonomicsSingleThread/a100/intervalOverlapSingleThread100_10k.csh"
	gonomicsSingleThread100_100k = "csh/gonomicsSingleThread/a100/intervalOverlapSingleThread100_100k.csh"
	gonomicsSingleThread100_1mi = "csh/gonomicsSingleThread/a100/intervalOverlapSingleThread100_1mi.csh"
	gonomicsSingleThread100_10mi = "csh/gonomicsSingleThread/a100/intervalOverlapSingleThread100_10mi.csh"
	// a1k
	gonomicsSingleThread1k_100 = "csh/gonomicsSingleThread/a1k/intervalOverlapSingleThread1k_100.csh"
	gonomicsSingleThread1k_1k = "csh/gonomicsSingleThread/a1k/intervalOverlapSingleThread1k_1k.csh"
	gonomicsSingleThread1k_10k = "csh/gonomicsSingleThread/a1k/intervalOverlapSingleThread1k_10k.csh"
	gonomicsSingleThread1k_100k = "csh/gonomicsSingleThread/a1k/intervalOverlapSingleThread1k_100k.csh"
	gonomicsSingleThread1k_1mi = "csh/gonomicsSingleThread/a1k/intervalOverlapSingleThread1k_1mi.csh"
	gonomicsSingleThread1k_10mi = "csh/gonomicsSingleThread/a1k/intervalOverlapSingleThread1k_10mi.csh"
	// a10k
	gonomicsSingleThread10k_100 = "csh/gonomicsSingleThread/a10k/intervalOverlapSingleThread10k_100.csh"
	gonomicsSingleThread10k_1k = "csh/gonomicsSingleThread/a10k/intervalOverlapSingleThread10k_1k.csh"
	gonomicsSingleThread10k_10k = "csh/gonomicsSingleThread/a10k/intervalOverlapSingleThread10k_10k.csh"
	gonomicsSingleThread10k_100k = "csh/gonomicsSingleThread/a10k/intervalOverlapSingleThread10k_100k.csh"
	gonomicsSingleThread10k_1mi = "csh/gonomicsSingleThread/a10k/intervalOverlapSingleThread10k_1mi.csh"
	gonomicsSingleThread10k_10mi = "csh/gonomicsSingleThread/a10k/intervalOverlapSingleThread10k_10mi.csh"
	// a100k
	gonomicsSingleThread100k_100 = "csh/gonomicsSingleThread/a100k/intervalOverlapSingleThread100k_100.csh"
	gonomicsSingleThread100k_1k = "csh/gonomicsSingleThread/a100k/intervalOverlapSingleThread100k_1k.csh"
	gonomicsSingleThread100k_10k = "csh/gonomicsSingleThread/a100k/intervalOverlapSingleThread100k_10k.csh"
	gonomicsSingleThread100k_100k = "csh/gonomicsSingleThread/a100k/intervalOverlapSingleThread100k_100k.csh"
	gonomicsSingleThread100k_1mi = "csh/gonomicsSingleThread/a100k/intervalOverlapSingleThread100k_1mi.csh"
	gonomicsSingleThread100k_10mi = "csh/gonomicsSingleThread/a100k/intervalOverlapSingleThread100k_10mi.csh"
	// a1mi
	gonomicsSingleThread1mi_100 = "csh/gonomicsSingleThread/a1mi/intervalOverlapSingleThread1mi_100.csh"
	gonomicsSingleThread1mi_1k = "csh/gonomicsSingleThread/a1mi/intervalOverlapSingleThread1mi_1k.csh"
	gonomicsSingleThread1mi_10k = "csh/gonomicsSingleThread/a1mi/intervalOverlapSingleThread1mi_10k.csh"
	gonomicsSingleThread1mi_100k = "csh/gonomicsSingleThread/a1mi/intervalOverlapSingleThread1mi_100k.csh"
	gonomicsSingleThread1mi_1mi = "csh/gonomicsSingleThread/a1mi/intervalOverlapSingleThread1mi_1mi.csh"
	gonomicsSingleThread1mi_10mi = "csh/gonomicsSingleThread/a1mi/intervalOverlapSingleThread1mi_10mi.csh"
	// a10mi
	gonomicsSingleThread10mi_100 = "csh/gonomicsSingleThread/a10mi/intervalOverlapSingleThread10mi_100.csh"
	gonomicsSingleThread10mi_1k = "csh/gonomicsSingleThread/a10mi/intervalOverlapSingleThread10mi_1k.csh"
	gonomicsSingleThread10mi_10k = "csh/gonomicsSingleThread/a10mi/intervalOverlapSingleThread10mi_10k.csh"
	gonomicsSingleThread10mi_100k = "csh/gonomicsSingleThread/a10mi/intervalOverlapSingleThread10mi_100k.csh"
	gonomicsSingleThread10mi_1mi = "csh/gonomicsSingleThread/a10mi/intervalOverlapSingleThread10mi_1mi.csh"
	gonomicsSingleThread10mi_10mi = "csh/gonomicsSingleThread/a10mi/intervalOverlapSingleThread10mi_10mi.csh"

	// gonomicsA1kIntermediateThread
	gonomicsA1kIntermediateThread1k_100 = "csh/gonomicsA1kIntermediateThread/a1k/intervalOverlapGonomics1k_100.csh"
	gonomicsA1kIntermediateThread1k_1k = "csh/gonomicsA1kIntermediateThread/a1k/intervalOverlapGonomics1k_1k.csh"
	gonomicsA1kIntermediateThread1k_10k = "csh/gonomicsA1kIntermediateThread/a1k/intervalOverlapGonomics1k_10k.csh"
	gonomicsA1kIntermediateThread1k_100k = "csh/gonomicsA1kIntermediateThread/a1k/intervalOverlapGonomics1k_100k.csh"
	gonomicsA1kIntermediateThread1k_1mi = "csh/gonomicsA1kIntermediateThread/a1k/intervalOverlapGonomics1k_1mi.csh"
	gonomicsA1kIntermediateThread1k_10mi = "csh/gonomicsA1kIntermediateThread/a1k/intervalOverlapGonomics1k_10mi.csh"

	// gonomicsA1kOptimalThread
	gonomicsA1kOptimalThread1k_100 = "csh/gonomicsA1kOptimalThread/a1k/intervalOverlapGonomics1k_100.csh"
	gonomicsA1kOptimalThread1k_1k = "csh/gonomicsA1kOptimalThread/a1k/intervalOverlapGonomics1k_1k.csh"
	gonomicsA1kOptimalThread1k_10k = "csh/gonomicsA1kOptimalThread/a1k/intervalOverlapGonomics1k_10k.csh"
	gonomicsA1kOptimalThread1k_100k = "csh/gonomicsA1kOptimalThread/a1k/intervalOverlapGonomics1k_100k.csh"
	gonomicsA1kOptimalThread1k_1mi = "csh/gonomicsA1kOptimalThread/a1k/intervalOverlapGonomics1k_1mi.csh"
	gonomicsA1kOptimalThread1k_10mi = "csh/gonomicsA1kOptimalThread/a1k/intervalOverlapGonomics1k_10mi.csh"
)

// for benchmarking memory usage
var memUsage []float64
var cmd *exec.Cmd

// benchmarking calls

// gonomicsA1kOptimalThread
func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kOptimalThread1k_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kOptimalThread1k_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kOptimalThread1k_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kOptimalThread1k_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kOptimalThread1k_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kOptimalThread1k_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// gonomicsA1kIntermediateThread
func BenchmarkGonomicsA1kIntermediateThreadIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kIntermediateThread1k_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kIntermediateThreadIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kIntermediateThread1k_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kIntermediateThreadIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kIntermediateThread1k_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kIntermediateThreadIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kIntermediateThread1k_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kIntermediateThreadIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kIntermediateThread1k_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsA1kIntermediateThreadIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsA1kIntermediateThread1k_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1k_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1k_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1k_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1k_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1k_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1k_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10k_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10k_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10k_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10k_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10k_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10k_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100k_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100k_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100k_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100k_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100k_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100k_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100k_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100k_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100k_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100k_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics100k_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools100k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops100k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread100k_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1mi_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1mi_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1mi_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1mi_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1mi_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1mi_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1mi_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1mi_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1mi_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1mi_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1mi_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1mi_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1mi_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1mi_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1mi_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1mi_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1mi_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1mi_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1mi_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1mi_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics1mi_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools1mi_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops1mi_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread1mi_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10mi_100)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10mi_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10mi_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10mi_100)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10mi_1k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10mi_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10mi_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10mi_1k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10mi_10k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10mi_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10mi_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10mi_10k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10mi_100k)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10mi_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10mi_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10mi_100k)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10mi_1mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10mi_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10mi_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10mi_1mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomics10mi_10mi)
		b.StopTimer() // prevent memUsage append from counting in benchmarking time
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedtoolsIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedtoolsIntervalOverlap(bedtools10mi_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkBedopsIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = bedopsIntervalOverlap(bedops10mi_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkSingleThreadIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd = gonomicsIntervalOverlap(gonomicsSingleThread10mi_10mi)
		b.StartTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

// function calls
func gonomicsIntervalOverlap(csh string) *exec.Cmd{
	cmd := exec.Command(csh)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func bedtoolsIntervalOverlap(csh string) *exec.Cmd{
	cmd := exec.Command(csh)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func bedopsIntervalOverlap(csh string) *exec.Cmd{
	cmd := exec.Command(csh)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}
