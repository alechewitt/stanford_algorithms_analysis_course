# Method returns a sorted list and the number of inversions
# in that list
def countInversions(aList): 
	if len(aList) == 1:
		return 0, aList;

	halfLength = len(aList) / 2
	aNumInv, listA = countInversions(aList[0:halfLength])
	bNumInv, listB = countInversions(aList[halfLength:])

	cNumInv = 0
	listC = []
	while len(listA) > 0 or len(listB) > 0:
		if len(listA)> 0 and len(listB) >0:
			if listA[0] <= listB[0]:
				# No inversio
				newItem = listA.pop(0)
				listC.append(newItem)
			else:
				# The current lengt of listA is the number
				# of inversions we have
				cNumInv += len(listA)
				newItem = listB.pop(0)
				listC.append(newItem)
		elif len(listA) > 0:
			# List B has no elements left but A is 
			# still full of elements. We have counted all inversions
			# but still need to append the rest of listA to listC
			listC += listA
			listA = []
		else:
			# listA is empty but listB has still some elements
			# we have counted all inversions but still need to append 
			# the rest of listB to listC
			listC += listB
			listB = []


	totalInversions = aNumInv + bNumInv + cNumInv

	return totalInversions, listC

testList = [6,5,4,3,2,1]
result = countInversions(testList)
print 'Original List = ' + str(testList)
print 'Number of inversions = ' + str(result[0])
print 'Sorted list = ' + str(result[1])

