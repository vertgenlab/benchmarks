#!/bin/bash
#
#SBATCH --mem=10G
#SBATCH --cpus-per-task=4

# change paths as needed

# example: use seed 1 to generate a bed file with 100 lines from hg38

# set path to the simulateBed executable
simulateBedPath='/hpc/group/vertgenlab/raven/GOPATH/src/github.com/vertgenlab/gonomics/cmd/simulateBed/simulateBed'

# make test files
$simulateBedPath -setSeed 1 -N 100 /hpc/group/vertgenlab/RefGenomes/hg38beds/hg38.noGap.bed testdata/100b.bed

# make test files compatible with bedtools
cut -f 1-3 testdata/100b.bed > testdata/test100b.bed

# rm bedtools-incompatible original files
rm testdata/100b.bed
