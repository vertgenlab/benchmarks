#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1mib.bed -b testdata/test10ka.bed > out.bedtools.10k.5.bed
