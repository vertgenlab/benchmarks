#!/bin/csh -ef
programs/intervalOverlap -threads 8 testdata/test100ka.bed testdata/test1kb.bed out.gonomics.100k.2.bed
