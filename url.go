package main
import (
	"fmt"
	"net/url"
)
const urls string="https://tnrdeveloping-hub.web.app/"
func main(){
	fmt.Println("url creation")
	result,err:=url.Parse(urls)
	if err!=nil{
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Fragment)
	fmt.Println(result.Port())
}