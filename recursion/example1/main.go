package main

import "fmt"

func main()  {
	countdown(10)
}

func countdown(n int){
	if n ==0{
		return
	}else{
	fmt.Println(n)
	countdown(n-1)
	}
}