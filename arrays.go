package main
import ("fmt")
func main(){
	var a=[3]int{1,2,3}
	var b=[...]int{4,5,6}
	c:=[3]int{7,8,9}
	d:=[...]int{10,11,12}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}