#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100b.bed -b testdata/test1mia.bed > out.bedtools.1mi.1.bed
