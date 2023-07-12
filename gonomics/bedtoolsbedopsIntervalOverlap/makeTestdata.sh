#!/bin/bash
#
#SBATCH --mem=10G
#SBATCH --cpus-per-task=4
#SBATCH --mail-type=END,FAIL --mail-user=yanting.luo@duke.edu

# make test files
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 2 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100c.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 3 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100d.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 4 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100e.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 5 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100f.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 6 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100g.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 7 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100h.bed

/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 2 -N 1000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/1kc.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 3 -N 1000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/1kd.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 4 -N 1000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/1ke.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 5 -N 1000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/1kf.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 6 -N 1000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/1kg.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 7 -N 1000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/1kh.bed

/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 2 -N 100000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100kc.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 3 -N 100000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100kd.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 4 -N 100000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100ke.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 5 -N 100000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100kf.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 6 -N 100000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100kg.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 7 -N 100000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100kh.bed

/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 2 -N 10000000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/10mic.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 3 -N 10000000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/10mid.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 4 -N 10000000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/10mie.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 5 -N 10000000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/10mif.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 6 -N 10000000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/10mig.bed
/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed -setSeed 7 -N 10000000 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/10mih.bed

# make test files compatible with bedtools
cut -f 1-3 testdata/100c.bed > testdata/test100c.bed
cut -f 1-3 testdata/100d.bed > testdata/test100d.bed
cut -f 1-3 testdata/100e.bed > testdata/test100e.bed
cut -f 1-3 testdata/100f.bed > testdata/test100f.bed
cut -f 1-3 testdata/100g.bed > testdata/test100g.bed
cut -f 1-3 testdata/100h.bed > testdata/test100h.bed

cut -f 1-3 testdata/1kc.bed > testdata/test1kc.bed
cut -f 1-3 testdata/1kd.bed > testdata/test1kd.bed
cut -f 1-3 testdata/1ke.bed > testdata/test1ke.bed
cut -f 1-3 testdata/1kf.bed > testdata/test1kf.bed
cut -f 1-3 testdata/1kg.bed > testdata/test1kg.bed
cut -f 1-3 testdata/1kh.bed > testdata/test1kh.bed

cut -f 1-3 testdata/100kc.bed > testdata/test100kc.bed
cut -f 1-3 testdata/100kd.bed > testdata/test100kd.bed
cut -f 1-3 testdata/100ke.bed > testdata/test100ke.bed
cut -f 1-3 testdata/100kf.bed > testdata/test100kf.bed
cut -f 1-3 testdata/100kg.bed > testdata/test100kg.bed
cut -f 1-3 testdata/100kh.bed > testdata/test100kh.bed

cut -f 1-3 testdata/10mic.bed > testdata/test10mic.bed
cut -f 1-3 testdata/10mid.bed > testdata/test10mid.bed
cut -f 1-3 testdata/10mie.bed > testdata/test10mie.bed
cut -f 1-3 testdata/10mif.bed > testdata/test10mif.bed
cut -f 1-3 testdata/10mig.bed > testdata/test10mig.bed
cut -f 1-3 testdata/10mih.bed > testdata/test10mih.bed

# rm bedtools-incompatible original files
rm testdata/100c.bed
rm testdata/100d.bed
rm testdata/100e.bed
rm testdata/100f.bed
rm testdata/100g.bed
rm testdata/100h.bed

rm testdata/1kc.bed
rm testdata/1kd.bed
rm testdata/1ke.bed
rm testdata/1kf.bed
rm testdata/1kg.bed
rm testdata/1kh.bed

rm testdata/100kc.bed
rm testdata/100kd.bed
rm testdata/100ke.bed
rm testdata/100kf.bed
rm testdata/100kg.bed
rm testdata/100kh.bed

rm testdata/10mic.bed
rm testdata/10mid.bed
rm testdata/10mie.bed
rm testdata/10mif.bed
rm testdata/10mig.bed
rm testdata/10mih.bed
