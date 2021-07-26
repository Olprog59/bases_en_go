package main

import "fmt"

func main() {
	nums := make([]int, 2, 3)
	nums[0] = -17
	nums[1] = 42

	fmt.Printf("%v (len='%v', cap='%v')\n", nums, len(nums), cap(nums))
	// fmt.Println(nums + " (len='" + len(nums)  + "', cap('" + cap(nums) + "'))")

	nums = append(nums, 25)
	fmt.Printf("%v (len='%v', cap='%v')\n", nums, len(nums), cap(nums))

	nums = append(nums, 2021)
	fmt.Printf("%v (len='%v', cap='%v')\n", nums, len(nums), cap(nums))
	nums = append(nums, 2021)
	nums = append(nums, 2021)
	nums = append(nums, 2021)
	nums = append(nums, 2021)
	fmt.Printf("%v (len='%v', cap='%v')\n", nums, len(nums), cap(nums))
}
