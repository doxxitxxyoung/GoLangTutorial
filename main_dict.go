package main

import (
    "fmt"
    "github.com/doxxitxxyoung/tutorial/mydict"
)

//  Create Dictionary
/*
func main(){
    dictionary := mydict.Dictionary{"first": "First words"}
//    dictionary["hello"] = "hello"
//    fmt.Println(dictionary["first"])

    //  Use methods to type
    definition, err := dictionary.Search("second")
    if err != nil {
        fmt.Println(err)
    } else{
        fmt.Println(definition)
    }
}
*/

/*
func main(){
    dictionary := mydict.Dictionary{"first": "First words"}
//    dictionary["hello"] = "hello"
//    fmt.Println(dictionary["first"])

    //  Add def
    word := "hello"
    definition := "Greeting"
    err := dictionary.Add(word, definition)
    if err != nil{
        fmt.Println(err)
    }
    hello, _ := dictionary.Search(word)
    fmt.Println("found ", word, " definition: ", hello)


    err2 := dictionary.Add(word, definition)
    if err2 != nil{
        fmt.Println(err2)
    }

    // Update
    dictionary.Update(word, "Second")
    word, _ = dictionary.Search(word)
    fmt.Println(word)
}
*/

/*
func main(){
    dictionary := mydict.Dictionary{}
    baseWord := "hello"
    dictionary.Add(baseWord, "First")
//    err := dictionary.Update(baseword, "Second")
    err := dictionary.Update("asdf", "Second")
    if err != nil{
        fmt.Println(err)
    }
    word, _ := dictionary.Search(baseWord)
    fmt.Println(word)
}
*/

func main(){
    dictionary := mydict.Dictionary{}
    baseWord := "hello"
    dictionary.Add(baseWord, "First")
    dictionary.Search(baseWord)
    dictionary.Delete(baseWord)
    word, err := dictionary.Search(baseWord)

    if err != nil{
        fmt.Println(err)
    } else {
        fmt.Println(word)
    }
}
