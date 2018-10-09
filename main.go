package main

import (
	"fmt"

)

func printString(content string){
	fmt.Printf("%c[%d;%dm%s%c[0m\n",0x1B, 0, 34,content, 0x1B)
}

func main(){
}
