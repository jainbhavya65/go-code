package main

import (
	"fmt"
	"log"
	"os"
)

var bigdigit = [][]string{
	{"000","0 0","0 0","0 0","0 0","0 0","000"},
	{" 1 ","11 "," 1 "," 1 "," 1 "," 1 ","111"},
	{"222","  2","  2"," 2 ","2  ","2  ","222"},
	{"333","  3","  3","333","  3","  3","333"},
	{"4 4","4 4","4 4","444","  4","  4","  4"},
	{"555","5  ","5  ","555","  5","  5","555"},
	{"666","6  ","6  ","666","6 6","6 6","666"},
	{"777","7 7","  7","  7","  7","  7","  7"},
	{"888","8 8","8 8","888","8 8","8 8","888"},
	{"999","9 9","9 9","  9","  9","  9","999"},
}

func main(){
  stringOfdigit := os.Args[1]
  for row := range bigdigit[0]{
	  line := ""
	  for colume := range stringOfdigit{
		  digit := stringOfdigit[colume] - '0'
		  if 0 <= digit && digit <= 9 {
			  line += bigdigit[digit][row] + " "
		  } else{
			  log.Fatal("invalid Whole Number")
		  }
		  }
		  fmt.Println(line)
	  }
  }