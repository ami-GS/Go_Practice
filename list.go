package main

import "fmt"

func main(){
	 list := []bool{true, true, false}
	 list2 := make([]bool, 10)//default false
	 fmt.Println(list)
	 fmt.Println(list2[3:])
}