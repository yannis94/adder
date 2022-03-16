package adder

import (
    "time"
    "log"
)

func Add(x ...int) (sum int) { // https://play.golang.org/p/i0nLnowJRgI
    for _, v := range x { sum += v }
    d := 2 * time.Second - time.Duration(sum) * 100 * time.Millisecond
    log.Printf("%v wait: %T %#v", x, d, d)
    time.Sleep(d)
    return
}

func Factoriel(n int) int {
    if (n < 0) {
        log.Printf("Error: Negativ number")
        return 0
    }
    result := 1

    for i:=1; i<n+1; i++ {
        result = result * i
    }

    return result
}

