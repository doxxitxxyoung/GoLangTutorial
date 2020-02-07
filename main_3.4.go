package main
//  Codes explaining GoRoutine

import (
    "fmt"
    "time"
)

/*
func main() {
   go sexyCount("nico") 
   go sexyCount("flynn") 
   //    Go routine will only last while the main function is running.
   //   The main function does not wait the go routine.
   go sexyCount("nico") 
   go sexyCount("flynn") 
   time.Sleep(time.Second * 5 )

   //   Channel is used to communicate with main function from go routine.
}
*/

func main() {
    //c := make(chan bool)
    c := make(chan string)
    people := [2]string{"nico", "flynn"}
    for _, person := range people {
        //go isSexy(person) 
        //result := go isSexy(person)   //  This is not Possible.
        go isSexy(person, c)
    }
//    result := <-c   //  tha main function waited even without time.Sleep.
//    fmt.Println(result)
    /*
    resultOne := <-c
    resultTwo := <-c

    fmt.Println("Waiting for messeges") //  Waiting for messege is Blocking Operation
    fmt.Println("Received 1st messege", resultOne)    //  we are getting a messege from a channel. Once we are given a messege, we will move to next line.
    fmt.Println("Received 2nd messege", resultTwo)    //  we are also waiting for A messege.
    */


    //  after the second messege received, main function will be finished
        

    //fmt.Println(<-c)    //  Deadlock: You are waiting for messages, but all the go routines are finished.

    for i:=0; i<len(people); i++{
        fmt.Println("Waiting for", i)
        fmt.Println(<-c)
    }
}


func sexyCount(person string){
    for i:=0; i<10; i++{
        fmt.Println(person, "is sexy", i)
        time.Sleep(time.Second)
    }
}
/*
func isSexy(person string) bool {
    time.Sleep(time.Second * 5)
    return true
}
*/
/*
func isSexy(person string, c chan bool) {
    time.Sleep(time.Second * 5)
    fmt.Println(person)
    c <- true
}
*/
func isSexy(person string, c chan string) {
    time.Sleep(time.Second * 5)
//    fmt.Println(person)
    c <- person + " is Sexy"    //  Messege is sent to channel which is given as arg
}


/*
Channels in Go Routines:
1. If your main func is finished : your Go Routines are gonna die.
2. You have to specify the type of data you are going to send & receive to the channel
3. You can send the messege by saying "channel <- messege"
4. <-c receiving a messege is a blocking operation.
