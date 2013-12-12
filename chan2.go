package main

import ("fmt"
        "time"
        "math/rand")

func main() {
    
    rand.Seed(time.Now().Unix())
    results := make(chan int,1)

    for i:=1;i<100;i++ {
        go func(index int){
            time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
            select {
            case results <- index:  //non-blocking send
            default:
            }
        }(i)
    }

    fmt.Println("result",<-results)
}




