package main

import (
	"fmt"
	"encoding/json"
)

type watch struct {
	P_ID   int    `json:"p_id"`
	P_Name string `json:"p_name"`
	P_Cost int    `json:"p_cost"`
	P_Cat  string
	P_Desc string `json:"p_desc"`
	P_Img  string `json:"p_img"`
}

func main() {
	fmt.Println("JSON Data:")
	encode()
	decode()
}

func encode() {
	watches := []watch{
		{1, "Provogue Basic Watch", 599, "Men's Watch", "Provogue Basic Watch with day and date display. Stylish brown strap and white dial.", "https://rukminim2.flixcart.com/image/612/612/xif0q/watch/e/v/e/1-sk-pg-4078-wyt-brwn-basic-with-day-and-date-display-provogue-original-imahffrywrx3x8zb.jpeg?q=70"},
		{2, "Elegant Black Watch", 899, "Unisex Watch", "Classic black watch with minimalist design, suitable for both men and women.", "https://rukminim2.flixcart.com/image/612/612/xif0q/watch/6/n/v/-original-imagpzzk4h4eqwmf.jpeg?q=70"},
	}
	finaljson, err := json.Marshal(watches)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",finaljson)
}
func decode(){
	jsonfromweb:=`{
		"P_Id": 2,
		"P_Name": "Elegant Black Watch",
		"P_Cost": 899,
		"P_Cat": "Unisex Watch",
		"P_Desc": "Classic black watch with minimalist design, suitable for both men and women.",
		"P_Img": "https://rukminim2.flixcart.com/image/612/612/xif0q/watch/6/n/v/-original-imagpzzk4h4eqwmf.jpeg?q=70"
	}`
	var watches watch
	check:=json.Valid([]byte(jsonfromweb))
	if check{
		fmt.Println("valid json")
		json.Unmarshal([]byte(jsonfromweb),&watches)
		fmt.Printf("%#v\n",watches)
	}else{
		fmt.Println("not valid")
	}
}