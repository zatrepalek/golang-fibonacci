package main

import (
    "fmt"
    "time"
)

func main() {
    
    start := time.Now()
    var elements = 30
    var a int64 = 0
    var b int64 = 1
    var tmp int64

    fmt.Printf("%d: %d\n", 0, a)

    for i := 1; i < elements + 1; i++ {
        tmp = b;
        b = a + b;
        a = tmp;
        fmt.Printf("%d: %d\n", i, a)
    }

    elapsed := time.Since(start)
    fmt.Printf("Time elapsed: %s\n", elapsed)
}
