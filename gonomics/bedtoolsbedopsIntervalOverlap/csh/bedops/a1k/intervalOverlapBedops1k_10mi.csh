#!/bin/csh -ef
programs/bin/sort-bed testdata/test10mib.bed > testdata/test10mib.sorted.bed
programs/bin/sort-bed testdata/test1ka.bed > testdata/test100a.sorted.bed
programs/bin/bedops --element-of 1 testdata/test10mib.sorted.bed testdata/test100a.sorted.bed > out.bedops.1k.6.bed
