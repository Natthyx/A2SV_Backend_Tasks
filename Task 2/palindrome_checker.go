package main

import ("fmt"
		"strings"
		"regexp"
	)


func IsPalindrome(s string) bool{
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w]`)
	s_clean := re.ReplaceAllString(s, "")

	n := len(s_clean)
	for i:=0 ; i < n/2; i++{
		if s_clean[i] != s_clean[n-1-i]{
			return false
		}
	} 
	return true

}

func main(){
	input := "121!!"
	if IsPalindrome(input){
		fmt.Printf("%s is palindrome", input)
	} else{
		fmt.Printf("%s is not palindrome", input)
	}

}