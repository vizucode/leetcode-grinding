package main

import "fmt"

func StringChallenge(str string) int {

	// str to array
	arrString := []string{}
	for _, value := range str {
		arrString = append(arrString, string(value))
	}

	// excetute
	similarCounter := 1
	for i := 1; i < len(arrString); i++ {
		if arrString[i] == arrString[i-1] {
			similarCounter++
		} else {
			// if different
			arrString[i] = arrString[i+1]
		}
	}

	if similarCounter == len(arrString) {
		return similarCounter
	}

	// code goes here
	return 0

}

func main() {
	fmt.Println(StringChallenge("cbc"))
}
