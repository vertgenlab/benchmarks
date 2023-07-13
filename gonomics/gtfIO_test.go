package benchmarking_test

import (
	"compress/gzip"
	"github.com/biogo/biogo/feat"
	biogoGtf "github.com/biogo/biogo/io/featio/gff"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/gtf"
	"io"
	"os"
	"testing"
)

var (
	gtfTest = "testdata/balAcu1.ncbiRefSeq.sorted.gtf.gz"
)

func gonomicsGtfRead(file string) {
	gtf.Read(file)
}

func biogoGtfRead(file string) {
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	r := biogoGtf.NewReader(f)
	exception.PanicOnErr(err)

	// gonomics is reading to a map so we will do the same here
	m := make(map[string]feat.Feature)
	var rec feat.Feature
	for rec, err = r.Read(); err != io.EOF; rec, err = r.Read() {
		exception.PanicOnErr(err)
		m[rec.Name()] = rec
	}
}

func gonomicsGtfWrite(records []*gtf.Gene) {
	for i := range records {
		gtf.WriteToFileHandle(io.Discard, records[i])
	}
}

func biogoGtfWrite(records []feat.Feature) {
	var err error
	w := biogoGtf.NewWriter(io.Discard, 50, false)
	for i := range records {
		_, err = w.Write(records[i])
		exception.PanicOnErr(err)
	}
}

func unzip(file string) string {
	unzipFile, err := os.CreateTemp("", "gonomicsBenchmarking")
	exception.PanicOnErr(err)
	defer unzipFile.Close()

	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	r, err := gzip.NewReader(f)
	exception.PanicOnErr(err)
	defer r.Close()

	_, err = unzipFile.ReadFrom(r)
	if err != io.EOF {
		exception.PanicOnErr(err)
	}

	return unzipFile.Name()
}

func BenchmarkGonomicsGtfRead(b *testing.B) {
	file := unzip(gtfTest)
	for i := 0; i < b.N; i++ {
		gonomicsGtfRead(file)
	}
}

func BenchmarkBiogoGtfRead(b *testing.B) {
	file := unzip(gtfTest)
	for i := 0; i < b.N; i++ {
		biogoGtfRead(file)
	}
}

func BenchmarkGonomicsGtfWrite(b *testing.B) {
	file := unzip(gtfTest)
	m := gtf.Read(file)
	var records []*gtf.Gene
	for _, val := range m {
		records = append(records, val)
	}
	for i := 0; i < b.N; i++ {
		gonomicsGtfWrite(records)
	}
}

func BenchmarkBiogoGtfWrite(b *testing.B) {
	file := unzip(gtfTest)
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	r := biogoGtf.NewReader(f)
	exception.PanicOnErr(err)

	var rec feat.Feature
	var records []feat.Feature
	for rec, err = r.Read(); err != io.EOF; rec, err = r.Read() {
		exception.PanicOnErr(err)
		records = append(records, rec)
	}

	for i := 0; i < b.N; i++ {
		biogoGtfWrite(records)
	}
}
