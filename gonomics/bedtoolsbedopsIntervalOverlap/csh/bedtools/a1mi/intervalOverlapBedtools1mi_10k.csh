#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10kb.bed -b testdata/test1mia.bed > out.bedtools.1mi.3.bed
