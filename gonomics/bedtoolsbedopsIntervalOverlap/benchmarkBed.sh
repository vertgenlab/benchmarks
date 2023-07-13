#!/bin/bash
#SBATCH -J benchmarkBed
#SBATCH -o benchmarkBed.out
#SBATCH --mem=40G
#SBATCH --cpus-per-task=13

: << 'COMMENT'
# 1. Generate executables for all tools

# Gonomics intervalOverlap
# Follow official instructions to install gonomics: https://github.com/vertgenlab/gonomics
#cd /path/to/gonomics
#cd cmd/intervalOverlap
#go build
# move intervalOverlap executable to desired path ($gonomicsIntervalOverlap. As of now it is "programs/intervalOverlap")

# Bedtools
# On the Duke Computing Cluster (DCC):
#module load Bedtools
# If not on DCC:
# Follow official instructions to install bedtools: https://bedtools.readthedocs.io/en/latest/content/installation.html
# move bedtools executable to desired path ($bedtools. As of now it is a direct DCC module call "bedtools")

# Bedops
# Follow official instructions to install bedops: https://bedops.readthedocs.io/en/latest/content/installation.html
# move bedops executables to desired path ($bedopsSortBed, $bedops. As of now they are "programs/bin/sort-bed" and "programs/bin/bedops")

# 2. Generate testdata
# Modify code in makeTestdata.sh to generate files needed
# On DCC:
#sbatch makeTestdata.sh
# If not on DCC:
# run lines of code in makeTestdata.sh as needed

# 3. Benchmark
# Only have 1 _test.go file in the directory where benchmarks will be run
# for runtime: bedtoolsbedopsIntervalOverlap20230707_test.go
# for memory: bedtoolsbedopsIntervalOverlapMem_test.go (requires that testdata/*.sorted.bed already exist)
# to optimize the number of workers: bedtoolsbedopsIntervalOverlapOptimizeThreads_test.go
# On DCC:
# run this script
#sbatch benchmarkBed.sh
# If not on DCC:
# change csh/*/*.csh executable paths ($gonomicsIntervalOverlap, $bedtools, $bedopsSortBed, $bedops)
# ideally, have 8-13 cpus available
#go test -timeout 0 -bench . -count 3

# 4. Analysis
# Use files in analysis/
# Including Excel, csv, Rmd
# To generate analysis/figures/
# Code to extract memUsage data:
#sed -i -e '1,4d' benchmarkBed.out;
#sed -i "$(( $(wc -l <benchmarkBed.out)-2+1 )),$ d" benchmarkBed.out;
#expand benchmarkBed.out | sed 's/ /,/g' | sed 's/,\{1,\}/,/g' | awk -F , '{print $(NF-3)}' > mem.out;
#cat mem.out | sed 's/B/,B/g' > mem2field.out;

COMMENT

go test -timeout 0 -bench . -count 3
