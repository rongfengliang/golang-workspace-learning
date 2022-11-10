package main

import (
	"fmt"

	"github.com/rongfengliang/userlogin"
)

func main() {
	conf := userlogin.UserLoginConf{
		Name: "dalong",
		Age:  333,
	}
	fmt.Printf("%v", conf)

}
