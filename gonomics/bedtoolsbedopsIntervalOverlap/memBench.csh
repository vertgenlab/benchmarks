#!/bin/csh -ef
#
#SBATCH -J memBench
#SBATCH -o memBench.out
#SBATCH --mem=40G
#SBATCH --cpus-per-task=13

# set path to executables
set bedopsBinDir = "programs"
set intervalOverlapBinDir = "programs"


echo "100_100"

set fileA = "testdata/test100a.bed"
set fileB = "testdata/test100b.bed"

wc -l $fileA $fileB

echo "bedops (three commands)"
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileA > a.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileB > b.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/bedops --element-of 1 b.bed a.bed > out.bedops.bed

echo "bedtools"
# module load Bedtools
/usr/bin/time -f '%E %M' bedtools intersect -u -wa -a $fileB -b $fileA > out.bedtools.bed

echo "intervalOverlap (1 worker)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap $fileA $fileB out.gonomics.1worker.bed

echo "intervalOverlap (8 workers)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap -threads=8 $fileA $fileB out.gonomics.8workers.bed

echo "done"


echo "100_10mi"

set fileA = "testdata/test100a.bed"
set fileB = "testdata/test10000000b.bed"

wc -l $fileA $fileB

echo "bedops (three commands)"
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileA > a.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileB > b.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/bedops --element-of 1 b.bed a.bed > out.bedops.bed

echo "bedtools"
# module load Bedtools
/usr/bin/time -f '%E %M' bedtools intersect -u -wa -a $fileB -b $fileA > out.bedtools.bed

echo "intervalOverlap (1 worker)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap $fileA $fileB out.gonomics.1worker.bed

echo "intervalOverlap (8 workers)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap -threads=8 $fileA $fileB out.gonomics.8workers.bed

echo "done"


echo "10mi_100"

set fileA = "testdata/test10000000a.bed"
set fileB = "testdata/test100b.bed"

wc -l $fileA $fileB

echo "bedops (three commands)"
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileA > a.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileB > b.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/bedops --element-of 1 b.bed a.bed > out.bedops.bed

echo "bedtools"
# module load Bedtools
/usr/bin/time -f '%E %M' bedtools intersect -u -wa -a $fileB -b $fileA > out.bedtools.bed

echo "intervalOverlap (1 worker)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap $fileA $fileB out.gonomics.1worker.bed

echo "intervalOverlap (8 workers)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap -threads=8 $fileA $fileB out.gonomics.8workers.bed

echo "done"


echo "10mi_10mi"

set fileA = "testdata/test10000000a.bed"
set fileB = "testdata/test10000000b.bed"

wc -l $fileA $fileB

echo "bedops (three commands)"
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileA > a.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/sort-bed $fileB > b.bed
/usr/bin/time -f '%E %M' $bedopsBinDir/bedops --element-of 1 b.bed a.bed > out.bedops.bed

echo "bedtools"
# module load Bedtools
/usr/bin/time -f '%E %M' bedtools intersect -u -wa -a $fileB -b $fileA > out.bedtools.bed

echo "intervalOverlap (1 worker)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap $fileA $fileB out.gonomics.1worker.bed

echo "intervalOverlap (8 workers)"
/usr/bin/time -f '%E %M' $intervalOverlapBinDir/intervalOverlap -threads=8 $fileA $fileB out.gonomics.8workers.bed

echo "done"
