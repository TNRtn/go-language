package main
import ("fmt")
func main(){
	a:=[]int{1,2,3,4,5}
	fmt.Println(a)
	a=append(a,6)
	fmt.Println(a)
	a=append(a,7,8)
	fmt.Println(a)
	b:=[]int{1,2}
	fmt.Println(b)
	c:=append(a,b...)
	fmt.Println(c)
}