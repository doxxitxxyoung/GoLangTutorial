package main

import (
    "fmt"
    "net/http"
    "errors"
)
var errRequestFailed = errors.New("Request Failed")

func main(){
//    var results map[string]string    //  Empty map, but you cannot write to uninitialized map.
//    var results map[string]string{}
    var results = make(map[string]string)
    urls := []string{
        "https://www.naver.com",
        "https://www.google.com",
        "https://www.airbnb.com",
        "https://www.facebook.com", //  Q: why need comma on the last 
    }

    //  Loop thru each of urls
    for _, url := range urls {
        //fmt.Println(url)
        result := "OK"
        err := hitURL(url)
        if err != nil {
            result = "FAILED"
        }
        results[url] = result
    }
    for url, result := range results {
        fmt.Println(url, result)
    }
}


func hitURL(url string) error {
    fmt.Println("Checking URL")
    resp, err := http.Get(url)
    if err != nil || resp.StatusCode >= 400{
        fmt.Println(err, resp.StatusCode)
        return errRequestFailed
    }
    return nil
}
