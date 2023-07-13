#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10mib.bed -b testdata/test1ka.bed > out.bedtools.1k.6.bed
