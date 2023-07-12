#!/bin/bash
#SBATCH -J benchmarkBed
#SBATCH -o benchmarkBed.out
#SBATCH --mem=40G
#SBATCH --cpus-per-task=24
#SBATCH --mail-type=END,FAIL --mail-user=yanting.luo@duke.edu

go test -timeout 0 -bench . -count 3
