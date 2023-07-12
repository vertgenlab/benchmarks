#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10kb.bed -b testdata/test1ka.bed > out.bedtools.1k.3.bed
