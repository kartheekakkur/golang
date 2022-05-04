package main

import "fmt"




type MinStack struct{
	data []int
}

func main()  {
  
	ms :=Constructor()

	ms.data=[]int {1,2,3}

ms.pop()

fmt.Println(ms.top())

	
}

func Constructor() MinStack{

	return MinStack{}
}

func (s MinStack) GetSize() int {
	return len(s.data)
}

func (s MinStack) IsEmpty()bool{
	return s.GetSize() == 0
}

func (s *MinStack) push(val int){
     s.data= append(s.data,val)
}

func (s *MinStack) pop(){
	if !s.IsEmpty(){
		s.data=s.data[:s.GetSize()-1]
	}
}

func (s *MinStack)top() int {

	return s.data[s.GetSize()-1]
	
}

func (s *MinStack) GetMin() int{

	stacksize := s.GetSize()

	ret := s.data[stacksize -1]
	for i :=stacksize-2; i>=0;i--{
		if s.data[i] < ret {
			ret =s.data[i]
		}
	}

	return ret
}