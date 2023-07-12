#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100kb.bed -b testdata/test1ka.bed > out.bedtools.1k.4.bed
