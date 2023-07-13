#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100b.bed -b testdata/test10ka.bed > out.bedtools.10k.1.bed
