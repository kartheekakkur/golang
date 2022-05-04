package palindrome

import (
	"testing"
)


func TestIsPalindrome(t *testing.T) {
 testCases :=[] int {121,-121,0,222,818}
 expectedResults :=[] bool {true,false,true,true,false}

 for index,data := range testCases{
	 if res :=isPalindrome(data); res != expectedResults[index]{
		 t.Errorf("expected %t, got %t ",expectedResults[index],res)
	 }
 }
}