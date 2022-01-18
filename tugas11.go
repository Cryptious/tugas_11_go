package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str1 string
	fmt.Scanln(&str1)
	var str2 string
	fmt.Scanln(&str2)
	var num1, err = strconv.Atoi(str1)  // Str to Int
	var num2, err2 = strconv.Atoi(str2) // Str to Int

	if err == nil {
		if err2 == nil {
			// fmt.Println(num)
			fmt.Println("hasilnya adalah : ", num1+num2) // test String or Int
		} else {
			fmt.Println("Data yang anda masukan salah!")
		}
	}
}
