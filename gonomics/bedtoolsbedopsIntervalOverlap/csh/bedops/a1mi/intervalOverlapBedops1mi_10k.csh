#!/bin/csh -ef
programs/bin/sort-bed testdata/test10kb.bed > testdata/test10kb.sorted.bed
programs/bin/sort-bed testdata/test1mia.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test10kb.sorted.bed testdata/test100a.sorted.bed > out.bedops.1mi.3.bed
