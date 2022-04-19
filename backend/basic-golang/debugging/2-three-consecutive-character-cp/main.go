package main

import "fmt"

func main() {
	/*
		Return the start index start from 1 (1 based) of any three consecutive character,
		if not exist return -1

		Example input/output
		www.ruangguru.com -> 1
		helloworld -> -1
	*/
	result1 := ThreeConsecutiveCharacterPosition("helloworld")
	fmt.Println(result1)

	result2 := ThreeConsecutiveCharacterPositionCorrect("wwwada")
	fmt.Println(result2)
}

func ThreeConsecutiveCharacterPosition(word string) int {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] && word[i] == word[i+2] {
			return i
		}
	}
	return -1
}

func ThreeConsecutiveCharacterPositionCorrect(word string) int {
	
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+1] && word[i] == word[i+2] {
			return i + 1
		}
	}
	return -1

}
