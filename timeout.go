package main

import ("fmt"
        "time")

func main() {

    fmt.Println(1e10)
    
    t := make(chan bool,1)
    go func() {
        time.Sleep(2 * time.Second)
        t <- true
    }()

    <- t
    fmt.Println("unblocked!")

    t2 := time.After(4 * time.Second)
    <- t2
    fmt.Println("unblocked!")
}




