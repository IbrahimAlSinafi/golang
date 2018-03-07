package main 
import "fmt"

func mult(a,b string) (string, string){
	return a, b
}

func sub(a,b int) int{
	return a-b
}

func main (){
	x, y := mult("hello", "world") // string func
	fmt.Println(x,y)

	fmt.Println(sub(10, 3)) // int func
}