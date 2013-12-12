package main

import ("fmt"
        "time"
        "math/rand")

//generator: function starts a go routine & returns a channel to read from
func boring(name string) <-chan string{

    c := make(chan string)

    go func(){
        for i:=1;;i++{
            time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
            c <- fmt.Sprintf("%s: boring %d!",name,i)
        }
    }()

    return c
}

//multiplexing: join 2 channels into 1
func fanIn(inputs... <-chan string) <-chan string {

    c := make(chan string)

    for _,input := range inputs {
        go func(_input <-chan string){
            for {
                c <- <- _input
            }
        }(input)
    }

    return c
}

func main() {
    
    c := fanIn(boring("alice"),boring("tom"),boring("charlie"))

    for i:=0;i<10;i++ {
        fmt.Println(<-c)
    }    
}




