package benchmarking_test

import (
	"fmt"
	biogoBam "github.com/biogo/hts/bam"
	biogoSam "github.com/biogo/hts/sam"
	"github.com/vertgenlab/gonomics/dna"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fasta"
	"github.com/vertgenlab/gonomics/numbers"
	"github.com/vertgenlab/gonomics/sam"
	"io"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

var (
	sam1k   = "testdata/1k_reads.sam"
	sam10k  = "testdata/10k_reads.sam"
	sam100k = "testdata/100k_reads.sam"
	bam1k   = "testdata/1k_reads.bam"
	bam10k  = "testdata/10k_reads.bam"
	bam100k = "testdata/100k_reads.bam"
	bam1m   = "testdata/1m_reads.bam"
)

func gonomicsBamRead(file string) {
	var s sam.Sam
	var err error
	br, _ := sam.OpenBam(file)
	for {
		_, err = sam.DecodeBam(br, &s)
		if err == io.EOF {
			break
		}
	}
	err = br.Close()
	exception.PanicOnErr(err)
}

func biogoBamRead(file string) {
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	br, err := biogoBam.NewReader(f, 1)
	exception.PanicOnErr(err)

	for _, err = br.Read(); err != io.EOF; _, err = br.Read() {
		exception.PanicOnErr(err)
	}
}

func gonomicsBamWrite(records []sam.Sam, header sam.Header) {
	out := sam.NewBamWriter(io.Discard, header)
	for j := range records {
		sam.WriteToBamFileHandle(out, records[j], 0)
	}
	err := out.Close()
	exception.PanicOnErr(err)
}

func biogoBamWrite(records []*biogoSam.Record, header *biogoSam.Header) {
	writer, err := biogoBam.NewWriter(io.Discard, header, 1)
	exception.PanicOnErr(err)
	for j := range records {
		err = writer.Write(records[j])
		exception.PanicOnErr(err)
	}
	err = writer.Close()
	exception.PanicOnErr(err)
}

func generateFasta() []fasta.Fasta {
	var ans []fasta.Fasta
	ans = append(ans, fasta.Fasta{Name: "chrI", Seq: generateSeq(1000)})
	ans = append(ans, fasta.Fasta{Name: "chrII", Seq: generateSeq(777)})
	ans = append(ans, fasta.Fasta{Name: "chrIII", Seq: generateSeq(123)})
	return ans
}

func generateSeq(length int) []dna.Base {
	ans := make([]dna.Base, length)
	for i := range ans {
		ans[i] = dna.Base(numbers.RandIntInRange(0, 4))
	}
	return ans
}

func execPysamBamRead(file string) *exec.Cmd {
	cmd := exec.Command("python3", "pysamBamRead.py", file)
	err := cmd.Run()
	exception.PanicOnErr(err)
	return cmd
}

func execGonomicsBamRead(file string) *exec.Cmd {
	cmd := exec.Command("./gonomicsBamRead/gonomicsBamRead", "-i", file)
	err := cmd.Run()
	exception.PanicOnErr(err)
	return cmd
}

func execBiogoBamRead(file string) *exec.Cmd {
	cmd := exec.Command("./biogoBamRead/biogoBamRead", "-i", file)
	err := cmd.Run()
	exception.PanicOnErr(err)
	return cmd
}

func execElprepBamRead(file string) *exec.Cmd {
	cmd := exec.Command("./elprepBamRead/elprepBamRead", "-i", file)
	err := cmd.Run()
	exception.PanicOnErr(err)
	return cmd
}

func execSamtoolsBamRead(file string) *exec.Cmd {
	cmd := exec.Command("samtools", "view", file)
	err := cmd.Run()
	exception.PanicOnErr(err)
	return cmd
}

func BenchmarkExecSamtoolsBamRead1m(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execSamtoolsBamRead(bam1m)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecPysamBamRead1m(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execPysamBamRead(bam1m)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecElprepBamRead1m(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execElprepBamRead(bam1m)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecGonomicsBamRead1m(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execGonomicsBamRead(bam1m)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}
func BenchmarkExecBiogoBamRead1m(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execBiogoBamRead(bam1m)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecSamtoolsBamRead100k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execSamtoolsBamRead(bam100k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecPysamBamRead100k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execPysamBamRead(bam100k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecElprepBamRead100k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execElprepBamRead(bam100k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecGonomicsBamRead100k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execGonomicsBamRead(bam100k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}
func BenchmarkExecBiogoBamRead100k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execBiogoBamRead(bam100k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecSamtoolsBamRead10k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execSamtoolsBamRead(bam10k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecPysamBamRead10k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execPysamBamRead(bam10k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecElprepBamRead10k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execElprepBamRead(bam10k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecGonomicsBamRead10k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execGonomicsBamRead(bam10k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}
func BenchmarkExecBiogoBamRead10k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execBiogoBamRead(bam10k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecSamtoolsBamRead1k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execSamtoolsBamRead(bam1k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecPysamBamRead1k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execPysamBamRead(bam1k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecElprepBamRead1k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execElprepBamRead(bam1k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecGonomicsBamRead1k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execGonomicsBamRead(bam1k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkExecBiogoBamRead1k(b *testing.B) {
	var memUsage []float64
	var cmd *exec.Cmd
	for i := 0; i < b.N; i++ {
		cmd = execBiogoBamRead(bam1k)
		b.StopTimer()
		memUsage = append(memUsage, float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
		b.StartTimer()
	}
	fmt.Print(" ", int(numbers.AverageFloat64(memUsage)))
}

func BenchmarkGonomicsBamReadComp100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsBamRead(bam100k)
	}
}

func BenchmarkBiogoBamReadComp100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoBamRead(bam100k)
	}
}

func BenchmarkGonomicsBamReadComp10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsBamRead(bam10k)
	}
}

func BenchmarkBiogoBamReadComp10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoBamRead(bam10k)
	}
}

func BenchmarkGonomicsBamReadComp1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsBamRead(bam1k)
	}
}

func BenchmarkBiogoBamReadComp1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoBamRead(bam1k)
	}
}

func BenchmarkGonomicsBamWritingComp100k(b *testing.B) {
	samChan, header := sam.GoReadToChan(sam100k)
	var records []sam.Sam
	for record := range samChan {
		records = append(records, record)
	}
	for i := 0; i < b.N; i++ {
		gonomicsBamWrite(records, header)
	}
}

func BenchmarkBiogoBamWritingComp100k(b *testing.B) {
	var err error
	var record *biogoSam.Record
	var records []*biogoSam.Record
	infile, err := os.Open(sam100k)
	exception.PanicOnErr(err)
	reader, err := biogoSam.NewReader(infile)
	exception.PanicOnErr(err)
	for record, err = reader.Read(); err != io.EOF; record, err = reader.Read() {
		exception.PanicOnErr(err)
		records = append(records, record)
	}
	err = infile.Close()
	exception.PanicOnErr(err)

	for i := 0; i < b.N; i++ {
		biogoBamWrite(records, reader.Header())
	}
}

func BenchmarkGonomicsBamWritingComp10k(b *testing.B) {
	samChan, header := sam.GoReadToChan(sam10k)
	var records []sam.Sam
	for record := range samChan {
		records = append(records, record)
	}
	for i := 0; i < b.N; i++ {
		gonomicsBamWrite(records, header)
	}
}

func BenchmarkBiogoBamWritingComp10k(b *testing.B) {
	var err error
	var record *biogoSam.Record
	var records []*biogoSam.Record
	infile, err := os.Open(sam10k)
	exception.PanicOnErr(err)
	reader, err := biogoSam.NewReader(infile)
	exception.PanicOnErr(err)
	for record, err = reader.Read(); err != io.EOF; record, err = reader.Read() {
		exception.PanicOnErr(err)
		records = append(records, record)
	}
	err = infile.Close()
	exception.PanicOnErr(err)

	for i := 0; i < b.N; i++ {
		biogoBamWrite(records, reader.Header())
	}
}

func BenchmarkGonomicsBamWritingComp1k(b *testing.B) {
	samChan, header := sam.GoReadToChan(sam1k)
	var records []sam.Sam
	for record := range samChan {
		records = append(records, record)
	}
	for i := 0; i < b.N; i++ {
		gonomicsBamWrite(records, header)
	}
}

func BenchmarkBiogoBamWritingComp1k(b *testing.B) {
	var err error
	var record *biogoSam.Record
	var records []*biogoSam.Record
	infile, err := os.Open(sam1k)
	exception.PanicOnErr(err)
	reader, err := biogoSam.NewReader(infile)
	exception.PanicOnErr(err)
	for record, err = reader.Read(); err != io.EOF; record, err = reader.Read() {
		exception.PanicOnErr(err)
		records = append(records, record)
	}
	err = infile.Close()
	exception.PanicOnErr(err)

	for i := 0; i < b.N; i++ {
		biogoBamWrite(records, reader.Header())
	}
}
