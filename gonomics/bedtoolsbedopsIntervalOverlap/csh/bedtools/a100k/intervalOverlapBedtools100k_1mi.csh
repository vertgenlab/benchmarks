#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1mib.bed -b testdata/test100ka.bed > out.bedtools.100k.5.bed
