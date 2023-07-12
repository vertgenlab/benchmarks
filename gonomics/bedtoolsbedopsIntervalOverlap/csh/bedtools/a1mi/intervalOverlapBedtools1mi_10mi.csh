#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10mib.bed -b testdata/test1mia.bed > out.bedtools.1mi.6.bed
