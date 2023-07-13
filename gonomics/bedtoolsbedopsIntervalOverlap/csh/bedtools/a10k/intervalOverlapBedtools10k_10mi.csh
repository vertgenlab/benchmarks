#!/bin/csh -ef
bedtools intersect -u -wa -a testdata/test10mib.bed -b testdata/test10ka.bed > out.bedtools.10k.6.bed
