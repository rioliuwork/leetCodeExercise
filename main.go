package main

import "leetCodeExercise/exercise"

func main() {
	points := []int{-2, 0, 3, -5, 2, -1}
	val := exercise.Constructor(points)
	v := val.SumRange(0, 2)
	print(v)
}
