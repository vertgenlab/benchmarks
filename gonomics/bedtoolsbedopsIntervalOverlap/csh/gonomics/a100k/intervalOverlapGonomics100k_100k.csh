#!/bin/csh -ef
programs/intervalOverlap -threads 8 testdata/test100ka.bed testdata/test100kb.bed out.gonomics.100k.4.bed
