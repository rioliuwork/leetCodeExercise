package main

import "leetCodeExercise/exercise"

func main() {
	candies := []int{1, 1, 1, 1}
	extraCandies := 2
	val := exercise.Shuffle(candies, extraCandies)
	print(val)
}
