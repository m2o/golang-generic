package main

import ("fmt"
        "time"
        "math/rand")

//generator: function starts a go routine & returns a channel to read from
func boring(name string, quit <-chan bool) <-chan string{

    c := make(chan string)

    go func(){
        for i:=1;;i++{
            time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
            select {
                case c <- fmt.Sprintf("%s: boring %d!",name,i):
                case <- quit:
                    fmt.Printf("%s shutdown\n",name)
                    return
            }
        }
    }()

    return c
}

func main() {
    
    quit := make(chan bool)  //quit channel
    c := boring("alice",quit)

    for i:=0;i<4;i++ {
        fmt.Println(<-c)
    }

    quit <- true
    time.Sleep(1*time.Second)
}




