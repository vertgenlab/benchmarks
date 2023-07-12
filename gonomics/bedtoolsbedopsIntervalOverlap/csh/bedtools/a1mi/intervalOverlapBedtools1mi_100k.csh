#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100kb.bed -b testdata/test1mia.bed > out.bedtools.1mi.4.bed
