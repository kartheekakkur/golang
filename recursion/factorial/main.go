package main

import "fmt"

func main(){
n :=factorial(0)

fmt.Println(n)
}


func factorial(n int) (fact int) {
	switch n {
	case 0:
		return 1
    case 1:
	    return 1
	default:
		return n*factorial(n-1)
	}	
}