package main

// SETUP FUNCTIONS
/*
import (
	"github.com/vertgenlab/gonomics/sam"
	"math/rand"
	"testing"
)

func generateFasta() fasta.Fasta {
	return fasta.Fasta{Name: "chrDummy", Seq: generateSeq(1_000_000)}
}

func mutateFasta(f fasta.Fasta, numMut int) (fasta.Fasta, []vcf.Vcf) {
	var ans fasta.Fasta
	ans.Name = f.Name
	ans.Seq = slices.Clone(f.Seq)

	variants := make([]vcf.Vcf, numMut)
	var pos int
	for i := 0; i < numMut; i++ {
		pos = numbers.RandIntInRange(0, len(ans.Seq))
		var v vcf.Vcf
		v.Pos = pos
		v.Chr = f.Name
		v.Ref = dna.BaseToString(ans.Seq[pos])
		ans.Seq[pos] = getAlt(ans.Seq[pos])
		v.Alt = []string{dna.BaseToString(ans.Seq[pos])}
		v.Filter = "."
		v.Info = "."
		v.Qual = 0
		v.Id = "."
		variants[i] = v
	}
	return ans, variants
}

func generateSeq(length int) []dna.Base {
	ans := make([]dna.Base, length)
	for i := range ans {
		ans[i] = dna.Base(numbers.RandIntInRange(0, 4))
	}
	return ans
}

func getAlt(b dna.Base) dna.Base {
	var newBase int
	for {
		newBase = numbers.RandIntInRange(0, 4)
		if dna.Base(newBase) != b {
			return dna.Base(newBase)
		}
	}
}

func TestMakeFasta(t *testing.T) {
	originalFasta := generateFasta()
	mutantFasta, variants := mutateFasta(originalFasta, len(originalFasta.Seq)/1000)
	fasta.Write("testdata/1mb.fa", []fasta.Fasta{originalFasta})
	fasta.Write("testdata/1mb_mut.fa", []fasta.Fasta{mutantFasta})
	vcf.Write("testdata/truePos.vcf", variants)
}

func TestMakeDilutions(t *testing.T) {
	refReads, header := sam.Read("testdata/mut0%_1000x.bam")
	altReads, _ := sam.Read("testdata/mut100%_1000x.bam")

	for i := range refReads {
		refReads[i].QName = "ref:" + refReads[i].QName
	}
	for i := range altReads {
		altReads[i].QName = "alt:" + altReads[i].QName
	}

	dilutionReads := make([]sam.Sam, 0, len(refReads))

	//// make 50% dilution
	//dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.5)
	//sam.Write("testdata/mut50%_1000x.sam", dilutionReads, header)
	//
	//// make 25% dilution
	//dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.25)
	//sam.Write("testdata/mut25%_1000x.sam", dilutionReads, header)
	//
	//// make 10% dilution
	//dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.1)
	//sam.Write("testdata/mut10%_1000x.sam", dilutionReads, header)

	// make 7.5% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.075)
	sam.Write("testdata/mut7.5%_1000x.sam", dilutionReads, header)

	// make 5% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.05)
	sam.Write("testdata/mut5%_1000x.sam", dilutionReads, header)

	// make 2.5% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.025)
	sam.Write("testdata/mut2.5%_1000x.sam", dilutionReads, header)

	//// make 1% dilution
	//dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.01)
	//sam.Write("testdata/mut1%_1000x.sam", dilutionReads, header)
	//
	//// make 0.1% dilution
	//dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.001)
	//sam.Write("testdata/mut0.1%_1000x.sam", dilutionReads, header)
}

func getDilution(ref, alt, ans []sam.Sam, altFrac float64) []sam.Sam {
	ans = ans[:0]
	for i := range ref {
		if sam.IsUnmapped(ref[i]) {
			continue
		}
		if rand.Float64() > altFrac {
			ans = append(ans, ref[i])
		}
	}
	for i := range alt {
		if sam.IsUnmapped(alt[i]) {
			continue
		}
		if rand.Float64() <= altFrac {
			ans = append(ans, alt[i])
		}
	}
	return ans
}
*/
