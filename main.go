package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main(){
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s ! This is monkey programming language \n", usr.Username)
	fmt.Printf("Type commands below \n")
	repl.Start(os.Stdin, os.Stdout)
}