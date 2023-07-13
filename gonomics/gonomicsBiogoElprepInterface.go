package benchmarking

import (
	"fmt"
	elprepBed "github.com/exascience/elprep/bed"
	"github.com/exascience/elprep/utils"
	gonomicsBed "github.com/vertgenlab/gonomics/bed"
	"github.com/vertgenlab/gonomics/exception"
	"strconv"
)

func GonomicsToElprepBed(b gonomicsBed.Bed) elprepBed.Region {
	var ans elprepBed.Region
	ans.Chrom = &b.Chrom
	ans.Start = int32(b.ChromStart)
	ans.End = int32(b.ChromEnd)
	ans.OptionalFields = make([]interface{}, b.FieldsInitialized-3)
	for i := range ans.OptionalFields {
		switch i {
		case 0: // Name
			ans.OptionalFields[i] = b.Name

		case 1: // Score
			ans.OptionalFields[i] = b.Score

		case 2: // Strand
			strand := string(b.Strand)
			ans.OptionalFields[i] = &strand

		case 3: // ThickStart
			thickStart, err := strconv.Atoi(b.Annotation[0])
			exception.PanicOnErr(err)
			ans.OptionalFields[i] = thickStart

		case 4: // ThickEnd
			thickEnd, err := strconv.Atoi(b.Annotation[1])
			exception.PanicOnErr(err)
			ans.OptionalFields[i] = thickEnd

		case 5: // ItemRgb
			ans.OptionalFields[i] = b.Annotation[2] == "on"

		case 6: // BlockCount
			blockCount, err := strconv.Atoi(b.Annotation[3])
			exception.PanicOnErr(err)
			ans.OptionalFields[i] = blockCount

		case 7: // BlockSizes
			blockSizes, err := strconv.Atoi(b.Annotation[4])
			exception.PanicOnErr(err)
			ans.OptionalFields[i] = blockSizes

		case 8: // BlockStarts
			blockStarts, err := strconv.Atoi(b.Annotation[5])
			exception.PanicOnErr(err)
			ans.OptionalFields[i] = blockStarts
		}
	}
	return ans
}

func ElprepToGonomicsBed(b elprepBed.Region) gonomicsBed.Bed {
	var ans gonomicsBed.Bed
	ans.FieldsInitialized = 3 + len(b.OptionalFields)

	ans.Chrom = *b.Chrom
	ans.ChromStart = int(b.Start)
	ans.ChromEnd = int(b.End)
	ans.Annotation = make([]string, ans.FieldsInitialized-6)
	for i := range b.OptionalFields {
		switch i {
		case 0: // Name
			ans.Name = b.OptionalFields[i].(string)

		case 1: // Score
			ans.Score = b.OptionalFields[i].(int)

		case 2: // Strand
			ans.Strand = gonomicsBed.Strand((*b.OptionalFields[i].(utils.Symbol))[0])

		case 3: // ThickStart
			ans.Annotation[0] = fmt.Sprintf("%d", b.OptionalFields[i].(int))

		case 4: // ThickEnd
			ans.Annotation[1] = fmt.Sprintf("%d", b.OptionalFields[i].(int))

		case 5: // ItemRgb
			if b.OptionalFields[i].(bool) {
				ans.Annotation[2] = "on"
			} else {
				ans.Annotation[2] = "."
			}

		case 6: // BlockCount
			ans.Annotation[3] = fmt.Sprintf("%d", b.OptionalFields[i].(int))

		case 7: // BlockSizes
			ans.Annotation[4] = fmt.Sprintf("%d", b.OptionalFields[i].(int))

		case 8: // BlockStarts
			ans.Annotation[5] = fmt.Sprintf("%d", b.OptionalFields[i].(int))
		}
	}
	return ans
}
