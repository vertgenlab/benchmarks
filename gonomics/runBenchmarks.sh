#!/bin/bash

# uncomment below to install pysam
# pip3 install --user pysam

# uncomment below to install samtools (via homebrew)
# brew install samtools

for DIR in *BamRead; do
	cd $DIR
	go build
	cd ..
done

go test -bench=Bed -benchmem
go test -bench=BigFasta -benchmem
go test -bench=Fastq -benchmem
go test -bench=Comp10k -benchmem
go test -bench=CompSam -benchmem
go test -bench=Gtf -benchmem

# cant use -benchmem for the exec tests so we print the average Maxrss for the process though the formatting can be messy due to how go runs benchmarks
# elprep automatically uses all procs so limit all benchmarks to 1 proc for comparison
# 1m bam files not added to repo due to size. Must be rederived with simulateSam before uncommenting next line
#GOMAXPROCS=1 go test -bench=BamRead1m
