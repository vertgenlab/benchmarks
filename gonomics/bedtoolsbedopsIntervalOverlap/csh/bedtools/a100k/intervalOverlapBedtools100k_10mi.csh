#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10mib.bed -b testdata/test100ka.bed > out.bedtools.100k.6.bed
