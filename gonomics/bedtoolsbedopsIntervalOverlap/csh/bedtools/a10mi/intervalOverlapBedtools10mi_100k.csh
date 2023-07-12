#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test100kb.bed -b testdata/test10mia.bed > out.bedtools.10mi.4.bed
