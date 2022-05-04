package palindrome

func isPalindrome(x int) bool {
    
    revNum :=0
    tempNum :=x

    if x<0{
        return false
    }
    for x >0 {
        revNum=(revNum*10)+(x%10)
        x= x/10
    }

    return revNum==tempNum

	}

