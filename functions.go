package main
import ("fmt")
func mesg(){
	fmt.Println("hi,i am tnr")
}
func sum(a int,b int) int {
	return a+b
}
func pro(a int,b int){
	fmt.Println(a*b)
}
func main(){
	mesg()
	fmt.Println(sum(3,2))
	pro(3,2)
}