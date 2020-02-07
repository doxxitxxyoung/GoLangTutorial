package accounts

import (
    "fmt"
    "errors"
)

//  create struct 

//  Account Struct
/*
type Account struct {
    //Owner string    //  should also be upper case, to be used when exported
    owner string
    balance int
}
*/

type Account struct {
    owner string
    balance int
}

//  Function creating struct

//  NewAccount Creates Account
func NewAccount(owner string) *Account{
    account := Account{owner: owner, balance: 0}
    return &account //  return address, new object.
}

//  Methods

//  Deposit x amount on your account
func (a *Account) Deposit(amount int) { //  * means that not copying, but using the obj.
    //  *Account now called "Pointer Receiver"
    fmt.Println("Gonna deposit", amount)
    a.balance += amount //  a here is the copy.of the object. the underlying obj is not modified.

}
//  a is called receiver, and is of type Account.

//  Balance of your account
func (a Account) Balance() int {
    return a.balance
}

//  We need to Error Handle as Well.

var errNoMoney = errors.New("Can't withdraw")
//  Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
    if a.balance < amount {
//        return error.Error()
//        return errors.New("Can't withdraw your money")
        return errNoMoney
    }
    a. balance -= amount
    return nil 
}

//  ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string){
    a.owner = newOwner
}
//  Owner of the account
func (a Account) Owner()  string {
    return a.owner
}

func (a Account) String() string {
//    return "whatever you want"
    return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
