#!/bin/csh -ef
programs/bin/sort-bed testdata/test100kb.bed > testdata/test100kb.sorted.bed
programs/bin/sort-bed testdata/test100a.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test100kb.sorted.bed testdata/test100a.sorted.bed > out.bedops.100.4.bed
