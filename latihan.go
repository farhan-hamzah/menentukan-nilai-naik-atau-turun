package main
import "fmt"

const NMAX int = 100

func main() {
    var A [NMAX]int
    var n, i int
    var isNaik bool = true 
    fmt.Scan(&n)
    for i = 0; i < n; i++ {
        fmt.Scan(&A[i])
    }

    for i = 0; i < n-1; i++ {
        if A[i] >= A[i+1] {
            isNaik = false
            break
        }
    }

    if isNaik {
        fmt.Print("Pola naik")
    } else {
        fmt.Print("Pola naik turun")
    }
}