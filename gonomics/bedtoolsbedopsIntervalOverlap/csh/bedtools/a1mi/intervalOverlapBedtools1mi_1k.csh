#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1kb.bed -b testdata/test1mia.bed > out.bedtools.1mi.2.bed
