#!/bin/bash

gunzip -c testdata/1mb.fa.gz > testdata/1mb.fa
gunzip -c testdata/1mb_mut.fa.gz > testdata/1mb_mut.fa

simulateSam -coverage 1000 -flatErrorRate 0.01 -setSeed 555 testdata/1mb_mut.fa stdout | samtools sort -@4 > testdata/mut100%_1000x.bam
simulateSam -coverage 1000 -flatErrorRate 0.01 -setSeed 666 testdata/1mb.fa stdout | samtools sort -@4 > testdata/mut0%_1000x.bam
simulateSam -coverage 1000 -flatErrorRate 0.01 -setSeed 777 testdata/1mb.fa stdout | samtools sort -@4 > testdata/normal_1000x.bam
go run setup/setup.go

samtools sort -@4 testdata/mut50%_1000x.sam > testdata/mut50%_1000x.bam
samtools sort -@4 testdata/mut25%_1000x.sam > testdata/mut25%_1000x.bam
samtools sort -@4 testdata/mut10%_1000x.sam > testdata/mut10%_1000x.bam
samtools sort -@4 testdata/mut7.5%_1000x.sam > testdata/mut7.5%_1000x.bam
samtools sort -@4 testdata/mut5%_1000x.sam > testdata/mut5%_1000x.bam
samtools sort -@4 testdata/mut2.5%_1000x.sam > testdata/mut2.5%_1000x.bam
samtools sort -@4 testdata/mut1%_1000x.sam > testdata/mut1%_1000x.bam

callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut100%_1000x.bam > testdata/mut100%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut50%_1000x.bam > testdata/mut50%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut25%_1000x.bam > testdata/mut25%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut10%_1000x.bam > testdata/mut10%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut7.5%_1000x.bam > testdata/mut7.5%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut5%_1000x.bam > testdata/mut5%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut2.5%_1000x.bam > testdata/mut2.5%_1000x.vcf
callVariants -p 2 -r 1mb.fa -minAF 0 -t 4 -maxStrandBias 1 -n testdata/normal_1000x.bam -i testdata/mut1%_1000x.bam > testdata/mut1%_1000x.vcf

rm testdata/*.sam
rm testdata/*.bam

go run vcfBenchmarking.go