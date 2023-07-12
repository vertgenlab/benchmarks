#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100b.bed -b testdata/test100ka.bed > out.bedtools.100k.1.bed
