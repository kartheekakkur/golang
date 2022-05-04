package main

import (
	"fmt"
)

func main()  {

	fmt.Println(reverseString("abcd"))

	st := []byte {100,99}

	fmt.Println(st)
	
}

func reverseString(s string) string {

	var st []byte

	for i := len(s)-1; i>=0;i--{

		fmt.Println(s[i])
		st =append(st,s[i])
	}

	return string(st)
}