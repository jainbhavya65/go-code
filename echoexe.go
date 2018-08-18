package main
import(
    "os"
    "fmt"
)

func main(){
	var s string
	 for i := 1 ; i<len(os.Args[0:]) ; i++ { 
		s = os.Args[i]
		fmt.Println(i,s)
	 }
}