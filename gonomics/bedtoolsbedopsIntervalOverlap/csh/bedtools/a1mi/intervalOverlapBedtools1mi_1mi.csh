#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1mib.bed -b testdata/test1mia.bed > out.bedtools.1mi.5.bed
