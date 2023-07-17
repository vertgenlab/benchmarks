#!/bin/csh -ef
programs/bin/sort-bed testdata/test100b.bed > testdata/test100b.sorted.bed
programs/bin/sort-bed testdata/test100ka.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test100b.sorted.bed testdata/test100a.sorted.bed > out.bedops.100k.1.bed