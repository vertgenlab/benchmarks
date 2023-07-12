#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1mib.bed -b testdata/test10mia.bed > out.bedtools.10mi.5.bed
