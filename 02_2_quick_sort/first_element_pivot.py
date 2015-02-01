def quickSort(aList):
	if len(aList) == 1 or len(aList) == 0:
		return aList

	else:
		# Pivot is selected as first element
		i = 1

		# i is the sepearation between items less than the pivot and greater than
		# j is seperation between items that we have already place and not
		for j, value in enumerate(aList):
			if aList[j] < aList[0]:
				# Item is less than the pivot element
				# We therefore need to swap with the left most item that is 
				# bigger than the pivot
				aList[j], aList[i] = aList[i], aList[j]
				i += 1
		aList[0], aList[i-1] = aList[i-1], aList[0]

		leftList = aList[:i]
		rightList = aList[i:]
		newList1 = quickSort(leftList)
		newList2 = quickSort(rightList)

		sortedList = newList1 + newList2

		return sortedList

testList = [1,4,5,6,7,810, 2, 45, 78, 34, 1, 3]
print quickSort(testList)

