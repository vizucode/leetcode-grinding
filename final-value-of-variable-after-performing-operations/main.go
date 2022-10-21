package main

import "fmt"

/*
	There is a programming language with only four operations and one variable X:

	++X and X++ increments the value of the variable X by 1.
	--X and X-- decrements the value of the variable X by 1.
	Initially, the value of X is 0.

	Given an array of strings operations containing a list of operations, return the final value of X after performing all the operations.
*/

/*
	Input: operations = ["--X","X++","X++"]
	Output: 1
	Explanation: The operations are performed as follows:
	Initially, X = 0.
	--X: X is decremented by 1, X =  0 - 1 = -1.
	X++: X is incremented by 1, X = -1 + 1 =  0.
	X++: X is incremented by 1, X =  0 + 1 =  1.
*/

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, element := range operations {
		switch element {
		case "--X":
			x -= 1
			break
		case "X--":
			x -= 1
			break
		case "X++":
			x += 1
			break
		case "++X":
			x += 1
			break
		}
	}
	return x
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
