#!/bin/csh -ef
set linesInA = "$1"
set linesInB = "$2"

set filenameA = testdata/test"$linesInA"a
set filenameB = testdata/test"$linesInB"b

programs/sort-bed $filenameB.bed > $filenameB.sorted.bed
programs/sort-bed $filenameA.bed > $filenameA.sorted.bed
programs/bedops --element-of 1 $filenameB.sorted.bed $filenameA.sorted.bed > out.bedops.$linesInA.$linesInB.bed
