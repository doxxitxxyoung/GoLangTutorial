package main

import (
//    "fmt"
    "github.com/doxxitxxyoung/GoLangTutorial/scrapper"
//    "net/http"
    "github.com/labstack/echo"
    "strings"
    "os"
)

func handleHome(c echo.Context) error {
//    return c.String(http.StatusOK, "Hello, World")
    return c.File("home.html")
}

const fileName string = "jobs.csv"

func handleScrape(c echo.Context) error {
    defer os.Remove(fileName)  //  when user download file, I want to delete file
    term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
    scrapper.Scrape(term)
    return c.Attachment(fileName, fileName)
}

func main(){
//    scrapper.Scrape("term")

    e := echo.New()
    e.GET("/",handleHome) 
    e.POST("/scrape", handleScrape)
    e.Logger.Fatal(e.Start(":1323"))
}
