#!/bin/csh -ef
programs/bin/sort-bed testdata/test1mib.bed > testdata/test1mib.sorted.bed
programs/bin/sort-bed testdata/test10ka.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test1mib.sorted.bed testdata/test100a.sorted.bed > out.bedops.10k.5.bed