package main
import ("fmt")
func main(){
	s1:=[]int{}
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)
	s2:=[]int{1,2,3,4}
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)

	//slice from arry
	arr1:=[5]int{10,11,12,13,14}
	s3:=arr1[2:5]
	fmt.Println(s3)

	//using make function
	s4:=make([]int,5,10)//(datatype,size,capacity)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	fmt.Println(s4)
}