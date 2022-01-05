package data

import (
	"fmt"
	"regexp"
)

//RemoveSpecialCharacters remove all special characters in a string that should only contain numbers
func RemoveSpecialCharacters(data []string)  {
	reg,err := regexp.Compile("[^0-9]+")
	if err !=nil{
		fmt.Println(err)
	}
	for index, d := range data {
		data[index] = reg.ReplaceAllString(d, "")
	}
	
}