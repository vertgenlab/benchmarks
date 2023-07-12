#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10kb.bed -b testdata/test100a.bed > out.bedtools.100.3.bed
