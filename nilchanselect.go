package main

import ("fmt")

func main() {
    
    _ca := make(chan string)
    _cb := make(chan string)

    go func(){
        for{
            _ca <- "a";
        }
    }()
    
    go func(){
        for{
            _cb <- "b";
        }
    }()

    
    ca := _ca
    cb := _cb
    
    for{
        select {
            case v := <- ca:
                fmt.Println(v)
                ca = nil  //disables first case temporarily
                cb = _cb
            case v := <- cb:
                fmt.Println(v)
                cb = nil  //disables first case temporarily
                ca = _ca
        }
    }

    panic("show me the stacks!")
}




