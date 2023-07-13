package benchmarking_test

import (
	"fmt"
	biogoAlphabet "github.com/biogo/biogo/alphabet"
	biogoFastq "github.com/biogo/biogo/io/seqio/fastq"
	"github.com/biogo/biogo/seq"
	biogoSeq "github.com/biogo/biogo/seq/linear"
	"github.com/vertgenlab/gonomics/dna"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fastq"
	"github.com/vertgenlab/gonomics/fileio"
	"io"
	"os"
	"testing"
)

var (
	fastq10k = "testdata/10k_reads.fastq"
)

func gonomicsFastqRead(file string) {
	f := fileio.EasyOpen(file)
	defer f.Close()
	var done bool
	for _, done = fastq.NextFastq(f); !done; _, done = fastq.NextFastq(f) {

	}
}

func biogoFastqRead(file string) {
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	r := biogoFastq.NewReader(f, biogoSeq.NewSeq("", nil, biogoAlphabet.DNA))
	exception.PanicOnErr(err)

	for _, err = r.Read(); err != io.EOF; _, err = r.Read() {
		exception.PanicOnErr(err)
	}
}

func gonomicsFastqWrite(records []fastq.Fastq) {
	var err error
	for i := range records {
		// gonomics fastq writer has not been updated to use io.Writer so copying the function instead to send to discard
		_, err = fmt.Fprintf(io.Discard, "@%s\n%s\n+\n%s\n", records[i].Name, dna.BasesToString(records[i].Seq), fastq.QualString(records[i].Qual))
		exception.PanicOnErr(err)
	}
}

func biogoFastqWrite(records []seq.Sequence) {
	var err error
	w := biogoFastq.NewWriter(io.Discard)
	for i := range records {
		_, err = w.Write(records[i])
		exception.PanicOnErr(err)
	}
}

func BenchmarkGonomicsFastqRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsFastqRead(fastq10k)
	}
}

func BenchmarkBiogoFastqRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoFastqRead(fastq10k)
	}
}

func BenchmarkGonomicsFastqWrite(b *testing.B) {
	records := fastq.Read(fastq10k)
	for i := 0; i < b.N; i++ {
		gonomicsFastqWrite(records)
	}
}

func BenchmarkBiogoFastqWrite(b *testing.B) {
	var records []seq.Sequence
	f, err := os.Open(fastq10k)
	exception.PanicOnErr(err)
	r := biogoFastq.NewReader(f, biogoSeq.NewSeq("", nil, biogoAlphabet.DNA))
	exception.PanicOnErr(err)
	var record seq.Sequence
	for record, err = r.Read(); err != io.EOF; record, err = r.Read() {
		records = append(records, record)
		exception.PanicOnErr(err)
	}

	for i := 0; i < b.N; i++ {
		biogoFastqWrite(records)
	}
}
