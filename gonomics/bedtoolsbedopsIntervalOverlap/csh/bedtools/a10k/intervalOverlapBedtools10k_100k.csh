#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100kb.bed -b testdata/test10ka.bed > out.bedtools.10k.4.bed
