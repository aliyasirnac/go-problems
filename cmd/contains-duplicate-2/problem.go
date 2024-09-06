package containsduplicate2

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i, n := range nums {
		if prevIndex, found := indexMap[n]; found && i-prevIndex <= k {
			return true
		}
		indexMap[n] = i
	}
	return false
}
