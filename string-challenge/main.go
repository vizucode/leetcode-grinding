package main

import "fmt"

func StringChallenge(str string) int {

	lenStr := len(str)
	hashmap := make(map[string]int)

	for _, value := range str {
		hashmap[string(value)]++
	}

	if hashmap["a"] == lenStr || hashmap["b"] == lenStr || hashmap["c"] == lenStr {
		return lenStr
	}

	if (hashmap["a"]%2) == (hashmap["a"]%2) && (hashmap["b"]%2) == (hashmap["c"]%2) {
		return 2
	}

	return 1

}

func main() {
	fmt.Println(StringChallenge("cbc"))    // 1
	fmt.Println(StringChallenge("cccc"))   // 4
	fmt.Println(StringChallenge("abcabc")) // ccabc -> cbbc -> bbc -> bb -> 2
}
