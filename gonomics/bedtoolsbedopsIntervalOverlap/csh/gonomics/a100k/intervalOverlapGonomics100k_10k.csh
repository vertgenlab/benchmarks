#!/bin/csh -ef
programs/intervalOverlap -threads 8 testdata/test100ka.bed testdata/test10kb.bed out.gonomics.100k.3.bed
