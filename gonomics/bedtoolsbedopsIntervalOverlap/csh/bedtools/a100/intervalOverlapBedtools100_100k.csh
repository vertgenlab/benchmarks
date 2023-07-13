#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100kb.bed -b testdata/test100a.bed > out.bedtools.100.4.bed
