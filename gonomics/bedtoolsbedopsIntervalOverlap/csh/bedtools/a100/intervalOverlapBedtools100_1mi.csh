#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1mib.bed -b testdata/test100a.bed > out.bedtools.100.5.bed
