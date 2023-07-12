import pysam
import sys
samfile = pysam.AlignmentFile(sys.argv[1])
numAlign = 0
for read in samfile.fetch():
    numAlign += 1
print(numAlign)
samfile.close()