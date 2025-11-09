/*
 * Author: Bianca Bertinato
 * Version: 1.0.0
 * Date: 2025-11-08
 * This program displays the multiplication table of 9 from 0 to 12.
 */

package main

import "fmt"

func main() {

    // CONSTANT
    const NUMBER int = 9

    // PROCESS & OUTPUT
    fmt.Println("MULTIPLICATION TABLE OF", NUMBER, ":")

    fmt.Println(NUMBER, "x 0 =", NUMBER*0)
    fmt.Println(NUMBER, "x 1 =", NUMBER*1)
    fmt.Println(NUMBER, "x 2 =", NUMBER*2)
    fmt.Println(NUMBER, "x 3 =", NUMBER*3)
    fmt.Println(NUMBER, "x 4 =", NUMBER*4)
    fmt.Println(NUMBER, "x 5 =", NUMBER*5)
    fmt.Println(NUMBER, "x 6 =", NUMBER*6)
    fmt.Println(NUMBER, "x 7 =", NUMBER*7)
    fmt.Println(NUMBER, "x 8 =", NUMBER*8)
    fmt.Println(NUMBER, "x 9 =", NUMBER*9)
    fmt.Println(NUMBER, "x 10 =", NUMBER*10)
    fmt.Println(NUMBER, "x 11 =", NUMBER*11)
    fmt.Println(NUMBER, "x 12 =", NUMBER*12)

    fmt.Println("\nDone.")
}