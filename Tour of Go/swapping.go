package main

import "fmt"

func swap(x, y string) (string,string){
	return y, x
}

func main(){
	a , b := swap("hello", "word")
	fmt.Println(a,b)
}