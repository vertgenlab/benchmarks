#!/bin/bash
#SBATCH -J benchmarkBed
#SBATCH -o benchmarkBed.out
#SBATCH --mem=40G
#SBATCH --cpus-per-task=13

: << 'COMMENT'

# 1. Generate executables for all tools (programs/)

# Gonomics intervalOverlap
# Follow official instructions to install gonomics: https://github.com/vertgenlab/gonomics
# Generate executables for intervalOverlap and simulateBed and make sure they have execution permissions
#cd /path/to/gonomics
#cd cmd/intervalOverlap
#go build
#chmod +x intervalOverlap
# move intervalOverlap executable to desired path (As of now it is "programs/intervalOverlap")
#cd /path/to/gonomics
#cd cmd/simulateBed
#go build
#chmod +x simulateBed
# move simulateBed executable to desired path (As of now it is "programs/simulateBed")

# Bedtools
# On the Duke Computing Cluster (DCC):
#module load Bedtools
# If not on DCC:
# Follow official instructions to install bedtools: https://bedtools.readthedocs.io/en/latest/content/installation.html
# move bedtools executable to desired path (As of now it is a DCC module call for "bedtools", i.e. the executable lives in the PATH)

# Bedops
# Follow official instructions to install bedops: https://bedops.readthedocs.io/en/latest/content/installation.html
# move bedops executables to desired path (As of now they are "programs/sort-bed" and "programs/bedops")

# 2. Generate testdata (testdata/)
# change path to simulateBed in the code below as needed ($simulateBedPath)
# run the code below to generate testdata

COMMENT

# set path to the simulateBed executable
simulateBedPath='programs/simulateBed'
# set path to the hg38NoGapBed
hg38NoGapBedPath='testdata/hg38.noGap.bed'

# make test files
$simulateBedPath -N 100 $hg38NoGapBedPath testdata/100a.bed
$simulateBedPath -setSeed 1 -N 100 $hg38NoGapBedPath testdata/100b.bed

$simulateBedPath -N 1000 $hg38NoGapBedPath testdata/1000a.bed
$simulateBedPath -setSeed 1 -N 1000 $hg38NoGapBedPath testdata/1000b.bed

$simulateBedPath -N 10000 $hg38NoGapBedPath testdata/10000a.bed
$simulateBedPath -setSeed 1 -N 10000 $hg38NoGapBedPath testdata/10000b.bed

$simulateBedPath -N 100000 $hg38NoGapBedPath testdata/100000a.bed
$simulateBedPath -setSeed 1 -N 100000 $hg38NoGapBedPath testdata/100000b.bed

$simulateBedPath -N 1000000 $hg38NoGapBedPath testdata/1000000a.bed
$simulateBedPath -setSeed 1 -N 1000000 $hg38NoGapBedPath testdata/1000000b.bed

$simulateBedPath -N 10000000 $hg38NoGapBedPath testdata/10000000a.bed
$simulateBedPath -setSeed 1 -N 10000000 $hg38NoGapBedPath testdata/10000000b.bed

# make test files compatible with bedtools
for testdataBedtoolsIncompatible in testdata/1*.bed
do
  filenameBase="$(basename $testdataBedtoolsIncompatible)"
  cut -f 1-3 $testdataBedtoolsIncompatible > testdata/test$filenameBase
  rm $testdataBedtoolsIncompatible
done

: << 'COMMENT'

# 3. Benchmark
# for runtime:
# experiment: bedtoolsbedopsIntervalOverlap_test.go
# On DCC:
# directly run this script
#sbatch benchmarkBed.sh
# If not on DCC:
# change csh/*.csh executable paths
# ideally, have 8-13 cpus available
#go test -timeout 0 -bench . -count 10

COMMENT

go test -timeout 0 -bench . -count 10

: << 'COMMENT'

# for memory
# run not this script but memBench.csh
# adjust executable paths in memBench.csh as needed
# On DCC:
#sbatch memBench.csh
# If not on DCC:
#./memBench.csh

COMMENT

./memBench.csh

: << 'COMMENT'

# 4. Analysis (analysis/)
# Use files in analysis/
# Including Excel, csv, Rmd
# To generate analysis/figures/

COMMENT
