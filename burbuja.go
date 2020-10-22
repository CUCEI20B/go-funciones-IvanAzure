package main

import "fmt"

func swap(s []int64, i int, j int){
	placeHolder := s[i]
	s[i] = s[j]
	s[j] = placeHolder
}

func Burbuja(s []int64){

	var flag bool = false

	for i := 0; i < (len(s)-1); i++{
		if s[i] > s[i+1]{
			swap(s, i, i+1)
			flag = true
		}

		if (flag){
			i = -1
			flag = false
		}
	}
}

func main()  {
	var s = []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11}
	
	Burbuja(s)
	fmt.Print(s)
}