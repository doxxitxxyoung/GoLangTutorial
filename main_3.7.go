package main

import (
    "fmt"
    "net/http"
    "errors"
)

type requestResult struct {
    url string
    status string
}

var errRequestFailed = errors.New("Request Failed")

func main(){
    results := make(map[string]string)
    c := make(chan requestResult) //  sends and receive a struct 
    urls := []string{
        "https://www.naver.com",
        "https://www.google.com",
        "https://www.airbnb.com",
        "https://www.facebook.com", //  Q: why need comma on the last 
    }

    //  Loop thru each of urls

    //  We will send channel to hitURL. Then we will send messege thru the channel in hitURL.
    for _, url := range urls {
        go hitURL(url, c)
    }
    //  Wait for the messege

    for i:=0; i<len(urls); i++ {
        //fmt.Println(<-c)    //  print messeges we are receiving
        result :=   <-c
        results[result.url] = result.status
    }

    for url, status := range results {
        fmt.Println(url, status)
    }

}


func hitURL(url string, c chan<- requestResult) {  //  chan<- means this channel is send only. not receive.k
    //c <- result{} //  I can send to the channel
    //fmt.Println(<-c)    //  I can receive from the channel

    resp, err := http.Get(url)
    status := "OK"
    if err != nil || resp.StatusCode >= 400 {
        c <-  requestResult{url: url, status: "FAILED"}  //  to channel we will send result struct 
    }
    c <- requestResult{url: url, status: status}
}
