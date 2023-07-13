#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10kb.bed -b testdata/test100ka.bed > out.bedtools.100k.3.bed
