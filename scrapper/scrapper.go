package scrapper

import (
    "fmt"
    "net/http"
//    "errors"
    "log"
    "github.com/PuerkitoBio/goquery"    //  Import goquery
    "strconv"
    "strings"
    "encoding/csv"
    "os"
)

type extractedJob struct {
    id string
    title string
    location string
    salary string
    summary string
}

//  Scrape Indeed by a term
func Scrape(term string){
    //var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
    baseURL := "https://kr.indeed.com/jobs?q="+term+"&limit=50"
    var jobs []extractedJob
    c := make(chan []extractedJob)

    totalPages := getPages(baseURL)
    //fmt.Println("Total pages:", totalPages)

    for i := 0; i < totalPages; i ++ {
        go getPage(i, baseURL, c)
//        jobs = append(jobs, extractedJobs...)   //  append contents of the extractedJobs:
    }
//    fmt.Println(jobs)

    //  Wait go routines to be finished
    for i:=0; i<totalPages; i++{
        extractedJobs := <-c
        jobs = append(jobs, extractedJobs...)
    }

    writeJobs(jobs)
    fmt.Println("Done, extracted", len(jobs))
}


func writeJobs(jobs []extractedJob) {

    //  Code Challenge : use Go Routine on writing as well.

    file, err := os.Create("jobs.csv")
    checkErr(err)
    w := csv.NewWriter(file)

    defer w.Flush()   //  Flush writes to the file

    headers := []string{"ID", "Title", "Location", "Salary", "Summary"}

    wErr := w.Write(headers)
    checkErr(wErr)

    for _, job := range jobs {
        jobSlice := []string{job.id, job.title, job.location, job.salary, job.summary}
        jwErr := w.Write(jobSlice)
        checkErr(jwErr)
    }
}

func getPage(page int, url string, mainC chan<- []extractedJob) {

    var jobs []extractedJob

    c := make(chan extractedJob)

    pageURL := url + "&start=" + strconv.Itoa(page*50)
    fmt.Println("Requesting", pageURL)
    res, err := http.Get(pageURL)
    checkErr(err)
    checkCode(res)

    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)
    checkErr(err)
    //  we want to get all the cards in the page.

    searchCards := doc.Find(".jobsearch-SerpJobCard")

    searchCards.Each(func(i int, card *goquery.Selection){
//        job := extractJob(card)
        go extractJob(card, c)  //  Now receives channel
//        jobs = append(jobs, job)
    })
    //  Number of messeges are same as the number of cards!
    for i:=0; i<searchCards.Length(); i++ {
        //  receive the messege.
        job := <-c  //  give me the messege from the channel
        jobs = append(jobs, job)
    }
    //  We need to wait for get pages to finish


//    return jobs
    mainC <- jobs
}

func getPages(url string) int {

    pages := 0

    //  return how many pages.
    res, err:= http.Get(url)

    checkErr(err)
    checkCode(res)

    defer res.Body.Close()  //  to prevent memory leaks.

    //  Go query
    doc, err := goquery.NewDocumentFromReader(res.Body)
    checkErr(err)

    doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
        //  you get a function with an integer & a selection
//        fmt.Println(s.Html())
        pages = s.Find("a").Length()   //  How many links are there in class pagination

    })
    return pages
}

func checkErr(err error){
    if err != nil {
        log.Fatalln(err)
    }
}
func checkCode(res *http.Response){
    if res.StatusCode != 200 {
        log.Fatalln("Request failed with Status:", res.StatusCode)
    }
}

func CleanString(str string) string {
    return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")   //  Fields returns an array
}

//func extractJob(card *goquery.Selection) extractedJob{
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
    //  Now this func sends extractedJob to the channel


    //  I want to extract for each card 's', which is div.
    //  div here has attribute "data-jk"
    id, _ := card.Attr("data-jk")
    //  Lets not make a struct
    title := CleanString(card.Find(".title>a").Text())
    location := CleanString(card.Find(".sjcl").Text())
    salary := CleanString(card.Find(".salaryText").Text())
    summary := CleanString(card.Find(".summary").Text())
//  fmt.Println(id, title, location, salary, summary)
//    return extractedJob{
    c <- extractedJob{
        id: id,
        title: title,
        location: location,
        salary: salary,
        summary: summary}


}
