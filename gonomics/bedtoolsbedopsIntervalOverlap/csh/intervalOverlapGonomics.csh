#!/bin/csh -ef
set numWorkers = "$1"
set linesInA = "$2"
set linesInB = "$3"

set filenameA = testdata/test"$linesInA"a
set filenameB = testdata/test"$linesInB"b

programs/intervalOverlap -threads $numWorkers $filenameA.bed $filenameB.bed out.gonomics.$linesInA.$linesInB.bed
