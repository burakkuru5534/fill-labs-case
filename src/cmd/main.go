package main

import (
	"fmt"
	"github.com/burakkuru5534/src/funcs"
)

func main() {

	//case 1: sorts a bunch of words by the number of character “a”s within the
	sorted := funcs.SortBunchOfWordsByTheNumberOfCharacterA([]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"})
	fmt.Printf("%v\n", sorted)

	//case2: the algorithm is x!+x
	funcs.MyBrilliantFunction(9)

	//case3: find the most repeated word in a bunch of words
	mostRepeated := funcs.FindMostRepeatedDataInArray([]string{"apple", "pie", "apple", "red", "red", "red"})
	fmt.Printf("%v\n", mostRepeated)

}
