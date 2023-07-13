#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10kb.bed -b testdata/test10ka.bed > out.bedtools.10k.3.bed
