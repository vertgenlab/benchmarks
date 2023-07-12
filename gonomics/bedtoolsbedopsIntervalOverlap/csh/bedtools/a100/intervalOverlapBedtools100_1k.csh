#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1kb.bed -b testdata/test100a.bed > out.bedtools.100.2.bed
