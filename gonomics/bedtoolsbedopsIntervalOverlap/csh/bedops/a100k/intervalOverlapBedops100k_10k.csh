#!/bin/csh -ef
programs/bin/sort-bed testdata/test10kb.bed > testdata/test10kb.sorted.bed
programs/bin/sort-bed testdata/test100ka.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test10kb.sorted.bed testdata/test100a.sorted.bed > out.bedops.100k.3.bed