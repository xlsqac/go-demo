package main

import (
 "fmt"
 "net/http"
 "sync"
 "time"
)

func main() {
     url := "https://httpbin.org/status/200"
     faild := 0

     var wg sync.WaitGroup
     wg.Add(3000)
     for range 3000 {
         go func() {
             defer wg.Done()
             _, err := http.Get(url)
             if err != nil {
                 faild++
                fmt.Println(err)
             }
          }()
     }
     now := time.Now()
     wg.Wait()
     fmt.Println(time.Since(now), faild)
}
