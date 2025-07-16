package main

import (
    "fmt"
    "regexp"
    "strings"
)

// rune -> for unicode character
func FrequenyCount(s string) map[string] int{
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w\s]`)
	s = re.ReplaceAllString(s, "")


	words := strings.Fields(s)

	count := make(map[string] int)

	for _ , word:= range(words){
		count[word] ++
	}

	return count
}

func main(){
	input := "Hello my name is GO dev hello again"
	result := FrequenyCount(input)

	fmt.Println("Frequency Table")
	for key, val := range result{
		fmt.Printf("%s: %d\n" , key, val)
	}
}