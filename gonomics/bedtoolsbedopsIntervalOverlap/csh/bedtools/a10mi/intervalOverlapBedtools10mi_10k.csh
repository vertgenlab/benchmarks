#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10kb.bed -b testdata/test10mia.bed > out.bedtools.10mi.3.bed