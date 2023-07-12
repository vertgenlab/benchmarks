#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1kb.bed -b testdata/test10ka.bed > out.bedtools.10k.2.bed
