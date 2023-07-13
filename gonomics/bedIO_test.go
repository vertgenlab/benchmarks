package benchmarking_test

import (
	"compress/gzip"
	"github.com/biogo/biogo/feat"
	biogoBed "github.com/biogo/biogo/io/featio/bed"
	"github.com/vertgenlab/gonomics/bed"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fileio"
	"io"
	"os"
	"strings"
	"testing"
)

var (
	bed10k = "testdata/10k.bed"
)

func gonomicsBedRead(file string) {
	f := fileio.EasyOpen(file)
	defer f.Close()
	var done bool
	for _, done = bed.NextBed(f); !done; _, done = bed.NextBed(f) {

	}
}

func biogoBedRead(file string) {
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	var br *biogoBed.Reader
	if strings.HasSuffix(file, ".gz") {
		gr, err := gzip.NewReader(f)
		exception.PanicOnErr(err)
		defer gr.Close()
		br, err = biogoBed.NewReader(gr, 3)
		exception.PanicOnErr(err)
	} else {
		br, err = biogoBed.NewReader(f, 3)
		exception.PanicOnErr(err)
	}

	for _, err = br.Read(); err != io.EOF; _, err = br.Read() {
		exception.PanicOnErr(err)
	}
}

func gonomicsBedWrite(records []bed.Bed) {
	for i := range records {
		bed.WriteBed(io.Discard, records[i])
	}
}

func biogoBedWrite(records []feat.Feature) {
	w, err := biogoBed.NewWriter(io.Discard, 3)
	exception.PanicOnErr(err)
	for i := range records {
		w.Write(records[i])
	}
}

func BenchmarkGonomicsBedRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsBedRead(bed10k)
	}
}

func BenchmarkBiogoBedRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoBedRead(bed10k)
	}
}

func BenchmarkGonomicsBedWrite(b *testing.B) {
	records := bed.Read(bed10k)
	for i := 0; i < b.N; i++ {
		gonomicsBedWrite(records)
	}
}

func BenchmarkBiogoBedWrite(b *testing.B) {
	var records []feat.Feature
	f, err := os.Open(bed10k)
	exception.PanicOnErr(err)
	br, err := biogoBed.NewReader(f, 3)
	exception.PanicOnErr(err)
	var record feat.Feature
	for record, err = br.Read(); err != io.EOF; record, err = br.Read() {
		records = append(records, record)
		exception.PanicOnErr(err)
	}

	for i := 0; i < b.N; i++ {
		biogoBedWrite(records)
	}
}
