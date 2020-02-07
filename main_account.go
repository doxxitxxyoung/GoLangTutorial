package main
//  banking

import (
    "fmt"
//    "github.com/doxxitxxyoung/tutorial/banking"
    "github.com/doxxitxxyoung/tutorial/accounts"
//    "log"
)

func main(){
//    account := banking.Account{ Owner: "nicolas", Balance: 1000}
    
    /*
    account := banking.Account{ Owner: "nicolas"}
    account.Owner = "pepe"
    */

    //  we will make struct making function.

    account := accounts.NewAccount("nico")

    //  Methods
    account.Deposit(10)

    err := account.Withdraw(20)

    if err != nil{
//        log.Fatalln(err)    //  Print and Kills the program
        fmt.Println(err)
    }

    fmt.Println(account.Balance(), account.Owner())

    //  you can see methods of class as "__str__" in Python
    fmt.Println(account)
}
