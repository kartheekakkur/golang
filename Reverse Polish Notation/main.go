package main

import (
	"fmt"
	"strconv"
)

func main()  {

	str :=[]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}

	fmt.Println(evalRPN(str))
	
}

func evalRPN(tokens []string) int {
	stack :=make([]int,len(tokens))

	top :=0

	for i,v := range tokens {
		switch ch := tokens[i];ch {
		case "+":
			stack[top-2] +=stack[top-1]
			top--
		case "-":
			stack[top-2] -=stack[top-1]
			top--
		case "/":
			stack[top-2] /=stack[top-1]
			top--
		case "*":
			stack[top-2] *=stack[top-1]
			top--
		default:
			stack[top],_ = strconv.Atoi(v)
			top++
		}
	}
return stack[0]

}