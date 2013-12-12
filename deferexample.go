package main

import "fmt"

func closeme(i int) {
    fmt.Println("closing ",i)
}

func work() (code int) {
    closeme(10)
    closeme(20)

    for i:=100;i<1000;i+=100 {
        defer closeme(i)
    }

    closeme(10000)
    closeme(20000)

    return 0
}

func main() {
    fmt.Println(work())
}




