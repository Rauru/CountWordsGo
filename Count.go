package main

import (
    "fmt"
    "strings"
)

func main() {
	goku := 0
	vegueta := 0
	yamcha :=0
    	Names :="Goku,vegueta,GOKU,yamcha,Vegueta,Goku"
    	result := strings.Split(Names, ",")
	for i := range result {
		if  strings.EqualFold(result[i], "goku"){
			goku ++
		}
		if  strings.EqualFold(result[i], "vegueta"){
			vegueta ++
		} 
		if  strings.EqualFold(result[i], "yamcha"){
			yamcha ++
		}  
		fmt.Println(result[i])
        	
        }
	
		fmt.Println(goku)
		fmt.Println(vegueta)
		fmt.Println(yamcha)
}