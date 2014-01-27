package main
import ("fmt"
       "math"
       )

func main(){
     N := 100000
     var v = []int{2, 3}

     for i := 5; i < N; i += 2{
     	 pFlag := true
	 for j := 0; j < len(v); j++{
	     if math.Sqrt(float64(v[j])) > float64(i){
	     	break
	     }
	     if i % v[j] == 0{
	     	pFlag = false
		break
	     }
	 }
	 if pFlag {
	 	 v = append(v,i)	 
	 }
     }
     fmt.Print(v)
}