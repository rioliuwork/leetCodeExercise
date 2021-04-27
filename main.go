package main

import "leetCodeExercise/exercise"

func main() {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}
	val := exercise.CountPoints(points, queries)
	print(val)
}
