package benchmarking_test

import (
	"compress/gzip"
	biogoAlphabet "github.com/biogo/biogo/alphabet"
	biogoFasta "github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq"
	biogoSeq "github.com/biogo/biogo/seq/linear"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fasta"
	"io"
	"os"
	"strings"
	"testing"
)

var (
	fa      = "testdata/test.fasta"
	faBig   = "testdata/test_big.fasta"
	faGz    = "testdata/test.fasta.gz"
	faBigGz = "testdata/test_big.fasta.gz"
)

func gonomicsFastaRead(file string) {
	fasta.Read(file)
}

func biogoFastaRead(file string) {
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	var r *biogoFasta.Reader
	if strings.HasSuffix(file, ".gz") {
		gr, err := gzip.NewReader(f)
		exception.PanicOnErr(err)
		defer gr.Close()
		r = biogoFasta.NewReader(gr, biogoSeq.NewSeq("", nil, biogoAlphabet.DNA))
	} else {
		r = biogoFasta.NewReader(f, biogoSeq.NewSeq("", nil, biogoAlphabet.DNA))
	}

	for _, err = r.Read(); err != io.EOF; _, err = r.Read() {
		exception.PanicOnErr(err)
	}
}

func gonomicsFastaWrite(records []fasta.Fasta) {
	fasta.WriteToFileHandle(io.Discard, records, 50)
}

func biogoFastaWrite(records []seq.Sequence) {
	w := biogoFasta.NewWriter(io.Discard, 50)
	var err error
	for i := range records {
		_, err = w.Write(records[i])
		exception.PanicOnErr(err)
	}
}

func BenchmarkGonomicsFastaRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsFastaRead(fa)
	}
}

func BenchmarkBiogoFastaRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoFastaRead(fa)
	}
}

func BenchmarkGonomicsBigFastaRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsFastaRead(faBig)
	}
}

func BenchmarkBiogoBigFastaRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoFastaRead(faBig)
	}
}

func BenchmarkGonomicsFastaReadGz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsFastaRead(faGz)
	}
}

func BenchmarkBiogoFastaReadGz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoFastaRead(faGz)
	}
}

func BenchmarkGonomicsBigFastaReadGz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsFastaRead(faBigGz)
	}
}

func BenchmarkBiogoBigFastaReadGz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoFastaRead(faBigGz)
	}
}

func BenchmarkGonomicsFastaWrite(b *testing.B) {
	records := fasta.Read(fa)
	for i := 0; i < b.N; i++ {
		gonomicsFastaWrite(records)
	}
}

func BenchmarkBiogoFastaWrite(b *testing.B) {
	f, err := os.Open(fa)
	exception.PanicOnErr(err)
	defer f.Close()
	r := biogoFasta.NewReader(f, biogoSeq.NewSeq("", nil, biogoAlphabet.DNA))
	var records []seq.Sequence
	rec, err := r.Read() // fa is a single entry
	exception.PanicOnErr(err)
	records = append(records, rec)
	for i := 0; i < b.N; i++ {
		biogoFastaWrite(records)
	}
}

func BenchmarkGonomicsBigFastaWrite(b *testing.B) {
	records := fasta.Read(faBig)
	for i := 0; i < b.N; i++ {
		gonomicsFastaWrite(records)
	}
}

func BenchmarkBiogoBigFastaWrite(b *testing.B) {
	f, err := os.Open(faBig)
	exception.PanicOnErr(err)
	defer f.Close()
	r := biogoFasta.NewReader(f, biogoSeq.NewSeq("", nil, biogoAlphabet.DNA))
	var records []seq.Sequence
	rec, err := r.Read() // faBig is a single entry
	exception.PanicOnErr(err)
	records = append(records, rec)
	for i := 0; i < b.N; i++ {
		biogoFastaWrite(records)
	}
}
