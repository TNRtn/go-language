package main
import ("fmt")
func factorial(n int) int{
	if n==0|| n==1{
		return 1
	} else{
		return n*factorial(n-1)
	}
}
func main(){
	a:=5
	fmt.Println(factorial(a))
}