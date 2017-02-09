package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

func main() {
  var url string
  if len(os.Args) != 2 {
      url = "https://github.com"
    } else {
      url = os.Args[1]
    }

    resp, err := http.Get(url)
    if err != nil {
      fmt.Println("Error occured while connecting. Exiting...")
      return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("Oops, I couldn't read the page. Sorry, bye...")
      return
    }
    fmt.Println(len(body))
  }
