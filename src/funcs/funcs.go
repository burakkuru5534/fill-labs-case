package funcs

import "fmt"

//case 1

func SortBunchOfWordsByTheNumberOfCharacterA(words []string) []string {

	//sorts a bunch of words by the number of character “a”s within the
	//word (decreasing order). If some words contain the same amount of character “a”s then we
	//need to sort those words by their lengths.

	//Example:
	/*
		Input :
		["aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"]
		Output :
		["aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"]
	*/

	for i := range words {

		for j := range words {

			if findNumberOfAInString(words[i]) > findNumberOfAInString(words[j]) {

				words[i], words[j] = words[j], words[i]
			} else if findNumberOfAInString(words[i]) == findNumberOfAInString(words[j]) {

				if len(words[i]) > len(words[j]) {

					words[i], words[j] = words[j], words[i]
				} else if len(words[i]) == len(words[j]) {
					if words[i] < words[j] {
						words[i], words[j] = words[j], words[i]
					}
				}
			}
		}

	}

	return words

}

//case 1 helper func o find number of a in our words
func findNumberOfAInString(word string) int {

	var count int

	for i := range word {

		if word[i] == 'a' {

			count++
		}
	}

	return count
}

//case 2

//Write a recursive function which takes one integer parameter. Please bear in mind that
//finding the algorithm needed to generate the output below is the main point of the question.
//example:
//Input: 9
//Output: 2 4 9

// the algorithm: f(x) = x! + x

func MyBrilliantFunction(n int) {

	for i := 1; i < n; i++ {
		resp := factorial(i)
		if resp > n {
			return
		}
		fmt.Println(resp + i)
	}
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

//case 3

//Write a function which takes one parameter as an array/list. Find most repeated data
//within a given array.
//Test with different datasets.

func FindMostRepeatedDataInArray(array []string) string {

	var mostRepeated string

	for i := range array {

		for j := range array {

			if array[i] == array[j] {

				mostRepeated = array[i]
			}
		}
	}

	return mostRepeated
}
