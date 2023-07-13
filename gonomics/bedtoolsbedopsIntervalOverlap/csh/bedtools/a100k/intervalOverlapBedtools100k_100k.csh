#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100kb.bed -b testdata/test100ka.bed > out.bedtools.100k.4.bed
