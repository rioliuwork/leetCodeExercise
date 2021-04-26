package main

import "leetCodeExercise/exercise"

func main() {
	s := "8*((1*(5+6))*(8/6))"
	val := exercise.MaxDepth(s)
	print(val)
}
