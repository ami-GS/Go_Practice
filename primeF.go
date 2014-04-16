package main
import ("fmt"
	"math")

func main(){
	N := 100000
	v := make([]int, N)
	v[0], v[1] = 1, 1

	for i := 4; i < N; i+=i{
		v[i] = 1
	}

	for i := 3; float64(i) < math.Sqrt(float64(N)); i++{
		for j :=i*2; j < N; j+=i{
			v[j] = 1
		}
	}

	for i := 2; i < N; i++{
		if v[i] == 0{
			fmt.Print(i, " ")
		}
	}
}
