package main

import(
		"strconv"
	   "fmt"
	   "os"
	   )

func fib(n int) int{
	 if n <= 2 {
	 	return 1
	 }
		return fib(n-1) + fib(n-2)
}

func main(){
	 i, _ := strconv.Atoi(os.Args[1])
	 fmt.Println(i)
	 fmt.Println(fib(i))
}