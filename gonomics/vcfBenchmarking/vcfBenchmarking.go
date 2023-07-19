package main

import (
	"fmt"
	"github.com/vertgenlab/gonomics/exception"
	"github.com/vertgenlab/gonomics/interval"
	"github.com/vertgenlab/gonomics/vcf"
	"path"
	"strconv"
	"strings"
)

var (
	truePosFile = "testdata/truePos.vcf"
	mut100File  = "testdata/mut100%_1000x.vcf"
	mut50File   = "testdata/mut50%_1000x.vcf"
	mut25File   = "testdata/mut25%_1000x.vcf"
	mut10File   = "testdata/mut10%_1000x.vcf"
	mut7_5File  = "testdata/mut7.5%_1000x.vcf"
	mut5File    = "testdata/mut5%_1000x.vcf"
	mut2_5File  = "testdata/mut2.5%_1000x.vcf"
	mut1File    = "testdata/mut1%_1000x.vcf"
)

func main() {
	truePos, _ := vcf.Read(truePosFile)
	var intervals []interval.Interval
	for i := range truePos {
		truePos[i].Pos += 1
		intervals = append(intervals, truePos[i])
	}
	tree := interval.BuildTree(intervals)

	fmt.Printf("File\tCutoff\tTPR\tFPR\n")
	longTestCutoffs(tree, truePos, mut100File)
	longTestCutoffs(tree, truePos, mut50File)
	longTestCutoffs(tree, truePos, mut25File)
	longTestCutoffs(tree, truePos, mut10File)
	longTestCutoffs(tree, truePos, mut7_5File)
	longTestCutoffs(tree, truePos, mut5File)
	longTestCutoffs(tree, truePos, mut2_5File)
	longTestCutoffs(tree, truePos, mut1File)
}

func longTestCutoffs(tree map[string]*interval.IntervalNode, truePos []vcf.Vcf, file string) {
	vars, _ := vcf.Read(file)

	var TPR, FPR, pCutoff float64
	fmt.Printf("%s\t%f\t%0.5f\t%0.5f\n", path.Base(file), 1.0, 1.0, 1.0)
	for pCutoff = 1; pCutoff >= 0.0001; pCutoff *= 0.9 {
		TPR, FPR = getRates(tree, truePos, vars, pCutoff)
		fmt.Printf("%s\t%f\t%0.5f\t%0.5f\n", path.Base(file), pCutoff, TPR, FPR)
	}
	fmt.Printf("%s\t%f\t%0.5f\t%0.5f\n", path.Base(file), 0.0, 0.0, 0.0)
}

func getRates(tree map[string]*interval.IntervalNode, truePos, vars []vcf.Vcf, cutoff float64) (TPR, FPR float64) {
	positives, negatives := classify(vars, cutoff)
	var truePositives, falsePositives, trueNegatives, falseNegatives int
	var overlaps []interval.Interval
	for i := range positives {
		overlaps = interval.Query(tree, positives[i], "any")
		if len(overlaps) == 0 {
			continue
		}
		if overlaps[0].(vcf.Vcf).Alt[0] == getMinPvAlt(positives[i]) {
			truePositives++
			continue
		}
	}
	falsePositives = len(positives) - truePositives
	falseNegatives = len(truePos) - truePositives
	trueNegatives = len(negatives) - falseNegatives
	TPR = float64(truePositives) / float64(truePositives+falseNegatives)
	FPR = 1 - (float64(trueNegatives) / float64(trueNegatives+falsePositives))
	return
}

func classify(vars []vcf.Vcf, cutoff float64) (positive []vcf.Vcf, negative []vcf.Vcf) {
	positive = make([]vcf.Vcf, 0, len(vars))
	negative = make([]vcf.Vcf, 0, len(vars))
	for i := range vars {
		if getPV(vars[i]) <= cutoff {
			positive = append(positive, vars[i])
		} else {
			negative = append(negative, vars[i])
		}
	}
	return
}

func getPV(v vcf.Vcf) float64 {
	words := strings.Split(v.Samples[0].FormatData[3], ",")
	var minPv, pv float64
	minPv = 1
	var err error
	for i := range words {
		pv, err = strconv.ParseFloat(words[i], 64)
		exception.PanicOnErr(err)
		if pv < minPv {
			minPv = pv
		}
	}
	return minPv
}

func getMinPvAlt(v vcf.Vcf) string {
	words := strings.Split(v.Samples[0].FormatData[3], ",")
	var minPv, pv float64
	minPv = 1
	var minIdx int
	var err error
	for i := range words {
		pv, err = strconv.ParseFloat(words[i], 64)
		exception.PanicOnErr(err)
		if pv < minPv {
			minPv = pv
			minIdx = i
		}
	}
	return v.Alt[minIdx]
}
