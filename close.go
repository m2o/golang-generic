package main

import ("fmt")

func main() {
    
     closing := make(chan chan error)

     go func(){
        for{
            select {
                case c := <- closing:
                    //do some cleanup
                    fmt.Println("cleaning up")
                    c <- nil
                    break
            }
        }
    }()

    _c := make(chan error)
    closing <- _c
    <- _c
    fmt.Println("main done")
}




