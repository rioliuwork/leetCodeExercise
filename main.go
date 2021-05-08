package main

import "leetCodeExercise/exercise"

func main() {
	points := [][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}
	val := exercise.MinTimeToVisitAllPoints(points)
	print(val)
}
