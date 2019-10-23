package main

import (
	"fmt"
	"gin/model"
)

func main()  {
	var use model.User
	fmt.Println(use)
	fmt.Println(&use)

}
