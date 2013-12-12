package main

import ("fmt"
        "time"
        "math/rand")

type Message struct {
	str string
	wait chan bool
}

//generator: function starts a go routine & returns a channel to read from
func boring(name string) <-chan Message{

    c := make(chan Message)
    wait := make(chan bool)

    go func(){
        for i:=1;;i++{
            time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
            c <- Message{fmt.Sprintf("%s: boring %d!",name,i),wait}
            <- wait
        }
    }()

    return c
}

//multiplexing: join 2 channels into 1
func fanIn(inputs... <-chan Message) <-chan Message {

    c := make(chan Message)

    for _,input := range inputs {
        go func(_input <-chan Message){
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
        m1 := <-c
        m2 := <-c
        m3 := <-c

        fmt.Println(m1.str)
        fmt.Println(m2.str)
        fmt.Println(m3.str)

        m1.wait <- true
        m2.wait <- true
        m3.wait <- true
    }    
}




