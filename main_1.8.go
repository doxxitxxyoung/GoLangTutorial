package main

import (
    "fmt"
)


func main(){
    /*
    a:= 2
    b := a
    a = 10
    fmt.Println(a, b)
    */

    //  How to look at memories of things
    a := 2
    b := &a //  b copies the memory address of the variables
    *b = 20 //  b is connected to a with memory
    fmt.Println(&a, &b) //  memory addresses of variables
    fmt.Println(*b) //  value of the address of memory
    fmt.Println(a)
}
