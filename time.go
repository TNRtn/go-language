package main
import ("fmt"
        "time")
func main(){
	fmt.Println("welcome to study plan")
	ptime:=time.Now()
	fmt.Println("current time",ptime)
	fmt.Println("format time:",ptime.Format("01-02-2006"))
	fmt.Println("format time:",ptime.Format("01-02-2006 15:04:05 monday"))
	create:=time.Date(2004,time.April,5,1,1,1,1,time.Local)
	//time.Date(year int,time.monthname,day int,hour int,min int,sec int,nsec int,time.Location)
	fmt.Println(create)
}