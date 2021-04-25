package main

import "leetCodeExercise/exercise"

func main() {
	accounts := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	val := exercise.MaximumWealth(accounts)
	print(val)
}
