#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1mib.bed -b testdata/test1ka.bed > out.bedtools.1k.5.bed
