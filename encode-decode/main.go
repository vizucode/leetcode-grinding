package main

import (
	"fmt"
	"strings"
)

/*
	Rules : lakukan 2 shift ke kanan untuk encode dan 2 unshift ke kiri untuk decode
*/

func encodeDecode(word string) string {
	commandArr := strings.Split(word, ">")
	command := strings.ReplaceAll(commandArr[0], "<", "")
	newstring := ""

	if command == "encode" {
		for _, element := range commandArr[1] {
			if element+2 > 122 {
				element = element + 2 - 26
			} else {
				element = element + 2
			}
			newstring += string(element)
		}
	} else {
		for _, element := range commandArr[1] {
			if element-2 < 97 {
				element = element - 2 + 26
			} else {
				element = element - 2
			}
			newstring += string(element)
		}
	}

	return newstring
}

func main() {
	fmt.Println(encodeDecode("<encode>abcd")) // cdef
	fmt.Println(encodeDecode("<encode>xyz"))  // zab

	fmt.Println(encodeDecode("<decode>cdef")) //abcd
	fmt.Println(encodeDecode("<decode>zab"))  // xyz
}
