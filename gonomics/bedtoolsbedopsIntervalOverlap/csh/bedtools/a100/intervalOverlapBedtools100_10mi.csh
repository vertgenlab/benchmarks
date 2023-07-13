#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10mib.bed -b testdata/test100a.bed > out.bedtools.100.6.bed
