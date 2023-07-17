package benchmarking_test

import (
	"log"
	"os/exec"
	"testing"
)

// benchmarking calls

// gonomicsA1kOptimalThread
func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4","1000","100")
	}
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4","1000","1000")
	}
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4","1000","10000")
	}
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4","1000","100000")
	}
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4","1000","1000000")
	}
}

func BenchmarkGonomicsA1kOptimalThreadIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4","1000","10000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100","100")
	}
}

func BenchmarkBedtoolsIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100","100")
	}
}

func BenchmarkBedopsIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100","100")
	}
}

func BenchmarkSingleThreadIntervalOverlap100_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100","100")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100","1000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100","1000")
	}
}

func BenchmarkBedopsIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100","1000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100","1000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100","10000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100","10000")
	}
}

func BenchmarkBedopsIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100","10000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100","10000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100","100000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100","100000")
	}
}

func BenchmarkBedopsIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100","100000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100","100000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100","1000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100","1000000")
	}
}

func BenchmarkBedopsIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100","1000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100","1000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100","10000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100","10000000")
	}
}

func BenchmarkBedopsIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100","10000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100","10000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000","100")
	}
}

func BenchmarkBedtoolsIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000","100")
	}
}

func BenchmarkBedopsIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000","100")
	}
}

func BenchmarkSingleThreadIntervalOverlap1k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000","100")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000","1000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000","1000")
	}
}

func BenchmarkBedopsIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000","1000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000","1000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000","10000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000","10000")
	}
}

func BenchmarkBedopsIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000","10000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000","10000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000","100000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000","100000")
	}
}

func BenchmarkBedopsIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000","100000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000","100000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000","1000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000","1000000")
	}
}

func BenchmarkBedopsIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000","1000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000","1000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000","10000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000","10000000")
	}
}

func BenchmarkBedopsIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000","10000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000","10000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000","100")
	}
}

func BenchmarkBedtoolsIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000","100")
	}
}

func BenchmarkBedopsIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000","100")
	}
}

func BenchmarkSingleThreadIntervalOverlap10k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000","100")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000","1000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000","1000")
	}
}

func BenchmarkBedopsIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000","1000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000","1000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000","10000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000","10000")
	}
}

func BenchmarkBedopsIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000","10000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000","10000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000","100000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000","100000")
	}
}

func BenchmarkBedopsIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000","100000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000","100000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000","1000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000","1000000")
	}
}

func BenchmarkBedopsIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000","1000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000","1000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000","10000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000","10000000")
	}
}

func BenchmarkBedopsIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000","10000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000","10000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100000","100")
	}
}

func BenchmarkBedtoolsIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100000","100")
	}
}

func BenchmarkBedopsIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100000","100")
	}
}

func BenchmarkSingleThreadIntervalOverlap100k_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100000","100")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100000","1000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100000","1000")
	}
}

func BenchmarkBedopsIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100000","1000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100k_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100000","1000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100000","10000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100000","10000")
	}
}

func BenchmarkBedopsIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100000","10000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100k_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100000","10000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100000","100000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100000","100000")
	}
}

func BenchmarkBedopsIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100000","100000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100k_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100000","100000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100000","1000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100000","1000000")
	}
}

func BenchmarkBedopsIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100000","1000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100k_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100000","1000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","100000","10000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("100000","10000000")
	}
}

func BenchmarkBedopsIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("100000","10000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap100k_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","100000","10000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000000","100")
	}
}

func BenchmarkBedtoolsIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000000","100")
	}
}

func BenchmarkBedopsIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000000","100")
	}
}

func BenchmarkSingleThreadIntervalOverlap1mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000000","100")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000000","1000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000000","1000")
	}
}

func BenchmarkBedopsIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000000","1000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000000","1000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000000","10000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000000","10000")
	}
}

func BenchmarkBedopsIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000000","10000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000000","10000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000000","100000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000000","100000")
	}
}

func BenchmarkBedopsIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000000","100000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000000","100000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000000","1000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000000","1000000")
	}
}

func BenchmarkBedopsIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000000","1000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000000","1000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","1000000","10000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("1000000","10000000")
	}
}

func BenchmarkBedopsIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("1000000","10000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap1mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","1000000","10000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000000","100")
	}
}

func BenchmarkBedtoolsIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000000","100")
	}
}

func BenchmarkBedopsIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000000","100")
	}
}

func BenchmarkSingleThreadIntervalOverlap10mi_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000000","100")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000000","1000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000000","1000")
	}
}

func BenchmarkBedopsIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000000","1000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10mi_1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000000","1000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000000","10000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000000","10000")
	}
}

func BenchmarkBedopsIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000000","10000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10mi_10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000000","10000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000000","100000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000000","100000")
	}
}

func BenchmarkBedopsIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000000","100000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10mi_100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000000","100000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000000","1000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000000","1000000")
	}
}

func BenchmarkBedopsIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000000","1000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10mi_1mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000000","1000000")
	}
}

// 1 set of gonomics, bedtools, bedops, gonomicsSingleThread
func BenchmarkGonomicsIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8","10000000","10000000")
	}
}

func BenchmarkBedtoolsIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedtoolsIntervalOverlap("10000000","10000000")
	}
}

func BenchmarkBedopsIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bedopsIntervalOverlap("10000000","10000000")
	}
}

func BenchmarkSingleThreadIntervalOverlap10mi_10mi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1","10000000","10000000")
	}
}

// function calls
func gonomicsIntervalOverlap(numWorkers string, linesInA string, linesInB string) {
	cmd := exec.Command("csh/intervalOverlapGonomics.csh", numWorkers, linesInA, linesInB)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func bedtoolsIntervalOverlap(linesInA string, linesInB string) {
	cmd := exec.Command("csh/intervalOverlapBedtools.csh", linesInA, linesInB)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func bedopsIntervalOverlap(linesInA string, linesInB string) {
	cmd := exec.Command("csh/intervalOverlapBedops.csh", linesInA, linesInB)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
