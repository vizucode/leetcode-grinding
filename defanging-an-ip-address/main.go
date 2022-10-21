package main

import "fmt"

/*
	Given a valid (IPv4) IP address, return a defanged version of that IP address.
	A defanged IP address replaces every period "." with "[.]".
*/

/*
	Input: address = "1.1.1.1"
	Output: "1[.]1[.]1[.]1"
*/

func defangIPaddr(address string) string {
	newstring := ""
	for _, element := range address {
		if string(element) != "." {
			newstring += string(element)
			continue
		}
		newstring += "[.]"
	}
	return newstring
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
