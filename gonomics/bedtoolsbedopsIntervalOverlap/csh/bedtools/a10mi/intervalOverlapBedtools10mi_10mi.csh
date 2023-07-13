#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10mib.bed -b testdata/test10mia.bed > out.bedtools.10mi.6.bed
