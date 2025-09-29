package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)


func main() {
user, err := user.Current()
if err != nil {
	panic(err)
}

fmt.Printf("Hwllo %s! This isthe Monkey proggramming language!\n", user.Username)
fmt.Printf("Feel free to type in commands\n")
repl.Start(os.Stdin, os.Stdout)

}
