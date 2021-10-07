package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)
func main(){
	if len(os.Args) < 2 {
		fmt.Println("Enter any value")
	}else{
   	var cel,err = strconv.ParseFloat(os.Args[1],64)
   	var Fer float64 = (cel*9/5) + 32
   	if err != nil{
	   log.Fatal(err)
   	}
	fmt.Println(Fer,"°F1111")
}
}

