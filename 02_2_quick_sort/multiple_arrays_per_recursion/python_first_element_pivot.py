
count = 0

def quickSort(aList):
	if len(aList) <= 1:
		return aList

	else:
		global count
		less = []
		equal = []
		greater = []

		count += len(aList) - 1
		# Pivot is selected as first element of aList
		for x in aList:
			if x < aList[0]:
				less.append(x)
			elif x == aList[0]:
				equal.append(x)
			else:
				greater.append(x)

		return quickSort(less) + equal + quickSort(greater)


testList = [1,2,3,4]
print quickSort(testList)
print "Number of comparisons = " + str(count)

