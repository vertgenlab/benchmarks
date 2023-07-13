#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test1kb.bed -b testdata/test10mia.bed > out.bedtools.10mi.2.bed
