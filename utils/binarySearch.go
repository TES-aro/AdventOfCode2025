package utilsAOC


func BinarySearch(orderedArray []int, target int) int{
	middle := len(orderedArray) / 2
	if len(orderedArray) == 0{
		return -1
	}
	if target < orderedArray[middle]{
		return BinarySearch(orderedArray[:middle],target)
	}
	if target > orderedArray[middle]{
		result := BinarySearch(orderedArray[middle+1:], target)
		if result != -1 {
			return result + middle+1
		}
	}
	return middle
}
