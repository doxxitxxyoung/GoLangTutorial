package main

import (
    "fmt"
//    "github.com/doxxitxxyoung/tutorial/something"
//    "strings"
)

/*
func main(){
    fmt.Println("Hello World")
    something.SayHello()    //  Exported function, since uppercased.
    //  something.sayBye()  //  Private function

    const name string = "const"
    //  name = "lynn" //    cannot assign to name

    var name string = "variable"

    //  inside of a function,
    name := "variable"  //  go will guess the type for you.
}
*/


/*
func multiply(a int, b int) int {   //  3rd int shows that function returns int
    return a*b
}
//  Go has multiple return values!
func lenAndUpper(name string) (int, string) {
    return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
    fmt.Println(words)
}

func main(){
    fmt.Println(multiply(2,2))
    //totalLength := lenAndUpper("paul")    //  You cannot init 1 var with 2 values
    //totalLength, _ = lenAndUpper("paul")    //  use _ to ignore value from return!
    totalLength, upperName := lenAndUpper("paul")

    fmt.Println(totalLength, upperName)

    repeatMe("A","b","C")
}
*/

/*
func lenAndUpper(name string) (length int, uppercase string) {
    defer fmt.Println("I am done")  //  this is executed after function is executed.
    length = len(name)   //  not :=, we are not creating, length is already created
    uppercase = strings.ToUpper(name)
    return 
}

//  You can do something after the function is finished!!


func main(){
    totalLength, upperName:= lenAndUpper("paul")

    fmt.Println(totalLength, upperName)
}
*/


/*
func superAdd(numbers ...int) int {
    fmt.Println(numbers)    //  numbers are in array

    /*
    for i:=0; i<len(numbers); i++{
        fmt.Println(numbers[i])
    }

    for index, number := range numbers{
        fmt.Println(index, number)
    }
    return 1
    */
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main(){
    result := superAdd(1,2,3,4,5)
    fmt.Println(result)
}
*/




