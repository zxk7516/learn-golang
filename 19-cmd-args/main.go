package main

import (
	"flag"
	"fmt"
	"os"
)

type Login struct {
	User string
	Pwd  string
	Host string
	Port int
}

func main() {

	for i, v := range os.Args {
		fmt.Printf("#%v args is : %v\n", i, v)
	}
	var login Login

	flag.StringVar(&login.User, "u", "", "用户名，默认空")
	flag.StringVar(&login.Pwd, "pwd", "", "用户名，默认空")
	flag.StringVar(&login.Host, "h", "localhost", "用户名，默认空")
	flag.IntVar(&login.Port, "port", 3306, "用户名，默认空")

	//  go run main.go -u root -port 123 -pwd zxkzxk
	flag.Parse()
	fmt.Println(login)
}
