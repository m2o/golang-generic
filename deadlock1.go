package main

import ("fmt")

func main() {
    
    c := make(chan int)

    go func(){
        i := <- c
        fmt.Println(i)
    }()
    
    //i := <- c
    //afmt.Println(i)

    panic("show me the stacks!")
}




