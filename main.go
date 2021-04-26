package main

import "leetCodeExercise/exercise"

func main() {
	candies := []int{1, 1, 1, 1}
	extraCandies := 3
	val := exercise.KidsWithCandies(candies, extraCandies)
	print(val)
}
