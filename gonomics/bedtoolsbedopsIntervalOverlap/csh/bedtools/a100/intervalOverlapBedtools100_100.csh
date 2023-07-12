#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100b.bed -b testdata/test100a.bed > out.bedtools.100.1.bed
