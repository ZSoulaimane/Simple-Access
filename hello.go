package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage : [username] [password]"
	errUser  = "Access denied for"
	errPWD   = "invalid PWD"
	accessOK = "Access Granted to"

	user, user2 = "jack", "mike"
	pass, pass2 = "1888", "1889"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != user && u != user2 {
		fmt.Println(errUser, u)
	} else if u == user && p == pass || u == user2 && p == pass2 {
		fmt.Println(accessOK, u)
	} else if u == user2 && p == pass2 {
		fmt.Println(accessOK, u)
	} else {
		fmt.Println(errPWD, u)
	}
}
