//https://www.urionlinejudge.com.br/judge/pt/problems/view/1240
package main

import "fmt"

func main() {

    var n, start int
    var num1, num2 string   

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&num1, &num2)
        if len(num2) > len(num1) {
            start = 0 
        } else {
            start = len(num1) - len(num2)
        }
        if (num1[start:] == num2) {
            fmt.Println("encaixa")
        } else {
            fmt.Println("nao encaixa")
        }
    }
}