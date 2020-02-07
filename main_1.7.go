package main

import (
    "fmt"
)

func canIDrink(age int) bool {
    /*
    if age < 18 {
        return false
    } else {
        return true
    }
    */

    //  Variable Expression is something very Unique in Go
    /*
    if koreanAge := age+2 ;koreanAge < 18 {     //  variable created here is only used for this if block.
        return false
    }
    return true 
    */

    /*
    switch age {
    case 10:
        return false
    case 18:
        return true
    }
    */

    /*
    switch koreanAge := age + 2; koreanAge {
    case koreanAge < 18:
        return false
    case koreanAge == 18:
        return true
    case koreanAge >  50:
        return false
    }
    return false

    */

}



func main() {
    fmt.Println(canIDrink(18))
}

