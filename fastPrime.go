package main

import (
	   "fmt"
	   "math"
//	   "regexp"
//	   "index/suffixarray"
		)

func main(){
	 N := 1000000
	 prime := make([]bool, N)
	 prime[0], prime[1] = true, true
//	 fmt.Println(int(math.Pow(float64(N), 0.5)))
	 for i := 2; i <= int(math.Pow(float64(N), 0.5)); i++{
	 	 if !prime[i]{
		 	for j := i*2; j < N; j += i{
				  prime[j] = true
				}
		 }
	 } 
//	 index := suffixarray.New(prime)
//	 fmt.Println(index.Lookup(false, -1))

	 for i := 2; i < N; i++{
	 	 if !prime[i]{
	 	 	fmt.Printf("%d ", i)
			}
		 }
	 fmt.Println()
}