#!/bin/csh -ef
set linesInA = "$1"
set linesInB = "$2"

set filenameA = testdata/test"$linesInA"a
set filenameB = testdata/test"$linesInB"b

programs/bin/sort-bed $filenameB.bed > $filenameB.sorted.bed
programs/bin/sort-bed $filenameA.bed > $filenameA.sorted.bed
programs/bin/bedops --element-of 1 $filenameB.sorted.bed $filenameA.sorted.bed > out.bedops.$linesInA.$linesInB.bed
