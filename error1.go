package main

import ("fmt")

type myerror string

func Error(text string) string {
    return "error:"+text
}

func main() {
    
    a := myerror("burek")
    fmt.Println(a)

    b := fmt.Errorf("pita")
    fmt.Println(b)
}




