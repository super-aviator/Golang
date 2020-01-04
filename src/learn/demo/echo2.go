package main
import (
	"fmt"
	"os"
)

func main(){
	s,sep:=""," " 
	for index,val:= range os.Args[1:]{
		s+= sep+val
		fmt.Println(index)
	}
	fmt.Println(s)
}