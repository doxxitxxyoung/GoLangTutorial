package main

import (
    "fmt"
)

//  First describe shape of struct.

type person struct {
    name string
    age int
    favFood []string
}

func main(){
    //  array
    /*
    names := [5]string{"nico", "lynn", "dal"}
    names[3] = "alala"
    names[4] = "alala"
    */
//    names[5] = "alala"  // out of bound for array

    //  unlimited bound : slice
    /*
    names := []string{"nico", "lynn", "dal"}
    names = append(names, "flynn")  //  does not modify name. append return modified slice.
    */

    //  Map
    /*
    nico := map[string]string{"name":"nico", "age":"12"}

    for key, value := range nico {
        fmt.Println(key, value)
    }
    */

    //  Struct
    favFood:= []string{"kimchi", "ramen"}
//    nico := person{"nico", 18, favFood}
    nico := person{name:"nico", age:18, favFood:favFood}
    fmt.Println(nico.name)

}
