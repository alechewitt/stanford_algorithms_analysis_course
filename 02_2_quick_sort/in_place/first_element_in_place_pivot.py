# This implementation only requires one list. per recursion
# It is therefore much more efficient use of memory
count = 0

def quickSort(aList):
	if len(aList) <= 1:
		return aList

	else:
		#Count number of comparisions needed
		global count
		count += len(aList) - 1
		# i is the sepearation between items less than the pivot and greater than
		# j is seperation between items that we have already place and not
		i = 1
		for j, value in enumerate(aList):
			if aList[j] < aList[0]:
				# Item is less than the pivot element
				# We therefore need to swap with the left most item that is 
				# bigger than the pivot
				aList[j], aList[i] = aList[i], aList[j]
				i += 1

		aList[0], aList[i-1] = aList[i-1], aList[0]

		return quickSort(aList[:i-1]) + [aList[i-1]] + quickSort(aList[i:])


testList = [1,2,3,4]
print quickSort(testList)
print "number of comparisions done: " + str(count)

