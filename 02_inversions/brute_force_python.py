def calcInversions(aList):
	numberInversions = 0;
	for i, iVal in enumerate(aList):
		for jVal in aList[i+1:]:
			if iVal > jVal:
				numberInversions += 1

	return numberInversions

testList = [6, 5, 4, 3, 2, 1]

print calcInversions(testList)