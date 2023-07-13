#!/bin/csh -ef
programs/bin/sort-bed testdata/test1kb.bed > testdata/test1kb.sorted.bed
programs/bin/sort-bed testdata/test10mia.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test1kb.sorted.bed testdata/test100a.sorted.bed > out.bedops.10mi.2.bed
