#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100b.bed -b testdata/test1ka.bed > out.bedtools.1k.1.bed
