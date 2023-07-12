package benchmarking_test

import (
	biogoSam "github.com/biogo/hts/sam"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/fileio"
	"github.com/vertgenlab/gonomics/sam"
	"io"
	"os"
	"testing"
)

func gonomicsSamRead(file string) {
	f := fileio.EasyOpen(file)
	defer f.Close()
	var done bool
	for _, done = sam.ReadNext(f); !done; _, done = sam.ReadNext(f) {

	}
}

func biogoSamRead(file string) {
	f, err := os.Open(file)
	exception.PanicOnErr(err)
	defer f.Close()

	var br *biogoSam.Reader
	br, err = biogoSam.NewReader(f)
	exception.PanicOnErr(err)

	for _, err = br.Read(); err != io.EOF; _, err = br.Read() {
		exception.PanicOnErr(err)
	}
}

func gonomicsSamWrite(records []sam.Sam, header sam.Header) {
	sam.WriteHeaderToFileHandle(io.Discard, header)
	for i := range records {
		sam.WriteToFileHandle(io.Discard, records[i])
	}
}

func biogoSamWrite(records []*biogoSam.Record, header *biogoSam.Header) {
	w, err := biogoSam.NewWriter(io.Discard, header, biogoSam.FlagString)
	exception.PanicOnErr(err)
	for i := range records {
		err = w.Write(records[i])
		exception.PanicOnErr(err)
	}
}

func BenchmarkGonomicsCompSamRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gonomicsSamRead(sam10k)
	}
}

func BenchmarkBiogoCompSamRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biogoSamRead(sam10k)
	}
}

func BenchmarkGonomicsCompSamWrite(b *testing.B) {
	records, header := sam.Read(sam10k)
	for i := 0; i < b.N; i++ {
		gonomicsSamWrite(records, header)
	}
}

func BenchmarkBiogoCompSamWrite(b *testing.B) {
	var records []*biogoSam.Record
	f, err := os.Open(sam10k)
	exception.PanicOnErr(err)
	defer f.Close()
	r, err := biogoSam.NewReader(f)
	exception.PanicOnErr(err)
	var record *biogoSam.Record
	for record, err = r.Read(); err != io.EOF; record, err = r.Read() {
		records = append(records, record)
		exception.PanicOnErr(err)
	}

	for i := 0; i < b.N; i++ {
		biogoSamWrite(records, r.Header())
	}
}
