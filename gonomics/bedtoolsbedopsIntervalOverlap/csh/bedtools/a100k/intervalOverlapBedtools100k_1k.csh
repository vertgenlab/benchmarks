#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1kb.bed -b testdata/test100ka.bed > out.bedtools.100k.2.bed
