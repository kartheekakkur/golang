package main

import (
	"fmt"
	"net/http"
	"time"
)


func main()  {

	links :=[] string{
		"https://google.com",
		"https://amazon.com",
		"https://solarwinds.com",
		"https://microsoft.com",
		"https://linedin.com",
	}

	c := make(chan string)

	for _,link := range links{

		go checklinks(link,c)
	}

	for l :=range c{
	   go func (link string){
		   time.Sleep(5 * time.Second)
		   checklinks(link,c)
	   }(l)
	}
}

func checklinks(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {

		fmt.Println("The Site is Offline",link)
		c <- link
	}

 fmt.Println("Is Up",link)

  c <- link

}