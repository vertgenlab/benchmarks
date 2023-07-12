package benchmarking_test

import (
	"log"
	"os/exec"
	"testing"
)

var (
	gonomicsIntervalOverlapProgram = "programs/intervalOverlap"
	testdata10mia = "testdata/test10mia.bed"
	testdata10mib = "testdata/test10mib.bed"
	outGonomics21 = "out1.21.bed"
)

// Benchmark calls
func BenchmarkGonomicsIntervalOverlap1k_100kLines_1threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("1", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_2threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("2", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_3threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("3", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_4threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("4", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_5threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("5", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_6threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("6", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_7threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("7", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_8threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("8", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_9threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("9", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_10threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("10", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_11threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("11", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_12threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("12", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_13threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("13", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_14threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("14", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_15threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("15", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_16threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("16", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_17threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("17", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_18threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("18", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_19threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("19", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_20threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("20", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_21threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("21", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_22threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("22", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_23threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("23", testdata10mia, testdata10mib, outGonomics21)
	}
}

func BenchmarkGonomicsIntervalOverlap1k_100kLines_24threads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsIntervalOverlap("24", testdata10mia, testdata10mib, outGonomics21)
	}
}

// function calls
func gonomicsIntervalOverlap(numThreads string, testdataA string, testdataB string, outGonomics string) {
	cmd := exec.Command(gonomicsIntervalOverlapProgram, "-threads", numThreads, testdataA, testdataB, outGonomics)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
