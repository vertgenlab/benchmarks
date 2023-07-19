package main

import (
	"github.com/vertgenlab/gonomics/dna"
	"github.com/vertgenlab/gonomics/numbers"
	"github.com/vertgenlab/gonomics/sam"
	"math/rand"
)

func getAlt(b dna.Base) dna.Base {
	var newBase int
	for {
		newBase = numbers.RandIntInRange(0, 4)
		if dna.Base(newBase) != b {
			return dna.Base(newBase)
		}
	}
}

func main() {
	refReads, header := sam.Read("testdata/mut0%_1000x.bam")
	altReads, _ := sam.Read("testdata/mut100%_1000x.bam")

	for i := range refReads {
		refReads[i].QName = "ref:" + refReads[i].QName
	}
	for i := range altReads {
		altReads[i].QName = "alt:" + altReads[i].QName
	}

	dilutionReads := make([]sam.Sam, 0, len(refReads))

	// make 50% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.5)
	sam.Write("testdata/mut50%_1000x.sam", dilutionReads, header)

	// make 25% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.25)
	sam.Write("testdata/mut25%_1000x.sam", dilutionReads, header)

	// make 10% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.1)
	sam.Write("testdata/mut10%_1000x.sam", dilutionReads, header)

	// make 7.5% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.075)
	sam.Write("testdata/mut7.5%_1000x.sam", dilutionReads, header)

	// make 5% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.05)
	sam.Write("testdata/mut5%_1000x.sam", dilutionReads, header)

	// make 2.5% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.025)
	sam.Write("testdata/mut2.5%_1000x.sam", dilutionReads, header)

	// make 1% dilution
	dilutionReads = getDilution(refReads, altReads, dilutionReads, 0.01)
	sam.Write("testdata/mut1%_1000x.sam", dilutionReads, header)

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
