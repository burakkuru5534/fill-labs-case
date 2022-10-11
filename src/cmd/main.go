package main

import (
	"fmt"
	"github.com/burakkuru5534/src/funcs"
)

func main() {
	println("Hello, world!")

	//sorts a bunch of words by the number of character “a”s within the

	sorted := funcs.SortBunchOfWordsByTheNumberOfCharacterA([]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"})

	fmt.Printf("%v\n", sorted)

	mostRepeated := funcs.FindMostRepeatedDataInArray([]string{"apple", "pie", "apple", "red", "red", "red"})

	fmt.Printf("%v\n", mostRepeated)

	output := funcs.C(9)
	fmt.Printf("%v\n", output)
}
