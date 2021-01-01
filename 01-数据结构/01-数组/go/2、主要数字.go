package main



func majorityElement(nums []int) int {

	dict := make(map[int]int)

	for _,v := range nums {
		dict[v]++
		if dict[v]*2 > len(nums) {
			return v
		}
	}

	return -1
}


func majorityElement2(nums []int) int {

	if len(nums) == 0 {
		 return -1
	}
	hero := nums[0]
	count := 1

	for _, v := range nums[1:] {
		if v == hero {
			count++
		}else {
			if count > 0 {
				count--
			}else {
				hero = v
				count++
			}
		}
	}

	if count > 0 {
		count = 0
		for _, v := range nums {
			if v == hero {
				count++
			}
		}
	}

	if count*2 > len(nums) {
		return hero
	}else {
		return -1
	}

}


