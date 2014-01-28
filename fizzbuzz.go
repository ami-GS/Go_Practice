package main

import(
	"strconv"
	"fmt"	
)

func main(){
     var fb string

     for i := 1; i <= 100; i++{
     	 switch{
		case i % 15 == 0:
		     fb = "FizzBuzz"
		case i % 5 == 0:
		     fb = "Buzz"
		case i % 3 == 0:
		     fb = "Fizz"
		default:
			fb = strconv.Itoa(i) 
	 }
	 fmt.Println(fb)
     }
}