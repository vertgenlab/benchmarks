#!/bin/csh -ef
programs/intervalOverlap -threads 8 testdata/test100ka.bed testdata/test1mib.bed out.gonomics.100k.5.bed
