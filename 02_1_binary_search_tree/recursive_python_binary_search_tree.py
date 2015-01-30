def findItem(item, aList):
	if len(aList) == 1:
		if aList[0] == item :
			return True
		else: 
			return False

	else:
		middleIndex = len(aList) / 2
		if item == aList[middleIndex] :
			return True
		elif item < aList[middleIndex]:
			newList = aList[:middleIndex]
			return findItem(item, newList)
		else:
			newList = aList[middleIndex:]
			return findItem(item, newList)		

aList = [0, 1, 2, 3, 4, 5, 7, 9, 233, 400, 435]

print findItem(9, aList)
