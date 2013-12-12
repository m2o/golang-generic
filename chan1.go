package main

import ("fmt"
        "time")

func worker(ch1 chan int,ch2 chan int,ch3 chan int) {

    var t int
    ticker := time.NewTicker(5 * time.Second)
    lasttask := time.Now().Unix()
    
    for {
        select {
        case t = <-ch1:
            fmt.Println("task 2 *",t)
            ch3 <- t*2
            lasttask = time.Now().Unix()
        case t = <-ch2:
            fmt.Println("task 3 *",t)
            ch3 <- t*3
            lasttask = time.Now().Unix()
        case <- ticker.C:
            fmt.Println("tick")
            if time.Now().Unix() - lasttask > 3 {
                close(ch3)
                return
            }
        }
    }
}

func worker2(ch1 chan int,ch2 chan int,ch3 chan int) {

    var t int
    
    for {
        select {
        case t = <-ch1:
            fmt.Println("task 2 *",t)
            ch3 <- t*2
        case t = <-ch2:
            fmt.Println("task 3 *",t)
            ch3 <- t*3
        case <-time.After(3*time.Second):
            close(ch3)
            return
        }
    }
}

func main() {
    
    t1 := make(chan int,20)
    t2 := make(chan int,20)
    t3 := make(chan int,50)

    //go worker(t1,t2,t3)
    go worker2(t1,t2,t3)

    for i:=1;i<5;i++ {
        t1 <- i
        t2 <- i
        time.Sleep(1 * time.Second)
    }

    fmt.Println("sent all")

    for r := range t3 {
        fmt.Println("result ",r)
    }
}




