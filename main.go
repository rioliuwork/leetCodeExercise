package main

import "leetCodeExercise/exercise"

func main() {
	var nums = []int{3, 2, 3, 3, 3}
	len := exercise.RemoveElement(nums, 3)
	for i := 0; i < len; i++ {
		print(nums[i])
	}
}
