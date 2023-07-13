#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100b.bed -b testdata/test10mia.bed > out.bedtools.10mi.1.bed
