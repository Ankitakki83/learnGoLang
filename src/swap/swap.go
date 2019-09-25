package main

import "fmt"

func swap(x,y string)(string,string){
	return y,x
}




func main(){
	a,b:=swap("hello","go")
	fmt.Println(a,b)
}