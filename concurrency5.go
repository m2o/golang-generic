package main

import ("fmt")

func main() {
    
    const n = 100000

    left := make(chan int)
    right := make(chan int)
    leftmost := left
    rightmost := right
    
    for i:=0;i<n;i++ {
        go func(in <-chan int,out chan<- int){
           out <- 1+<-in
        }(left,right)
        
        rightmost = right
        left = right
        right = make(chan int)
    }
    

    go func(){
        leftmost <- 1
    }()

    fmt.Println(<-rightmost)
}




