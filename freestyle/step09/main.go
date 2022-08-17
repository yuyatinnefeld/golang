package main

import (
    "fmt"
    "os/exec"
)

func main() {
    echo := "echo"
    arg0 := "-e"
    arg1 := "Hello world"
    cmd1 := exec.Command(echo, arg0, arg1)
    stdout1, err := cmd1.Output()

    pwd := "pwd"
    cmd2 := exec.Command(pwd)
    stdout2, err := cmd2.Output()

	ls := "ls"
    cmd3 := exec.Command(ls)
    stdout3, err := cmd3.Output()


    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(stdout1))
    fmt.Println(string(stdout2))
    fmt.Println(string(stdout3))

}