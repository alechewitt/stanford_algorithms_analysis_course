
def quickSort(aList):
	if len(aList) <= 1:
		return aList

	else:
		less = []
		equal = []
		greater = []

		# Pivot is selected as first element of aList
		for x in aList:
			if x < aList[0]:
				less.append(x)
			elif x == aList[0]:
				equal.append(x)
			else:
				greater.append(x)

		return quickSort(less) + equal + quickSort(greater)


testList = [6, 4, 2, 3, 9, 1, 8, 12, 34]
print quickSort(testList)
