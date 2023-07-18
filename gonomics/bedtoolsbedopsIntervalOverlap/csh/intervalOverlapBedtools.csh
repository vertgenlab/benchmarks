#!/bin/csh -ef
set linesInA = "$1"
set linesInB = "$2"

set filenameA = testdata/test"$linesInA"a
set filenameB = testdata/test"$linesInB"b

bedtools intersect -u -wa -a $filenameB.bed -b $filenameA.bed > out.bedtools.$linesInA.$linesInB.bed
