
count = 0

def quickSort(aList):
	if len(aList) <= 1:
		return aList

	else:
		global count
		less = []
		greater = []

		# number of comparisons is one less than the length of the list
		count += len(aList) - 1
		# Pivot is selected as first element of aList
		pivotIndex = len(aList) - 1

		#Swap so pivot is first element of list
		aList[0], aList[pivotIndex] = aList[pivotIndex], aList[0]
		print "printing the list"
		print aList
		for x in aList[1:]:
			if x < aList[0]:
				less.append(x)
			else:
				greater.append(x)

		return quickSort(less) + [aList[0]] + quickSort(greater)


testList = [1,2,3,4, 7, 1, 3]
print quickSort(testList)
print "Number of comparisons = " + str(count)

