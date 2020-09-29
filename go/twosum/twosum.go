package main

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i := 0 ; i < len(nums) ; i++ {
		compl := target - nums[i]
		if j, ok := mymap[compl]; ok && j != i {
			return []int{j, i}
		}
		mymap[nums[i]] = i
	}
	return nil
}
